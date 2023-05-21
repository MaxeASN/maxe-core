package relayer

import (
	"sync"
	"time"

	"github.com/MaxeASN/maxe-core/relayer/event"
	"github.com/MaxeASN/maxe-core/service/txmgr"
	"github.com/ethereum/go-ethereum/log"
)

type Handle func(event *event.Work, receiptCh chan *event.TxReceipt, receiptMgr *txmgr.DefaultTxManager) error

// WorkerPool is a pool of workers. It is used to be thread-safe.
type WorkerPool struct {
	WorkerFunc  Handle
	MaxWorkers  int
	idleTimeout time.Duration

	activeWorkers int
	canStop       bool

	ready        []*workerCh
	stop         chan struct{}
	workerChPool sync.Pool
	lock         sync.Mutex

	receiptMgr *txmgr.DefaultTxManager
}

type workerCh struct {
	lastUseTime time.Time
	ch          chan *event.Work
	receiptCh   chan *event.TxReceipt
}

// Start starts a worker pool
func (wp *WorkerPool) Start() {
	// check if the pool is already started
	// if wp.stop != nil {
	// 	log.Warn("Worker pool already started", "worker pool len", len(wp.ready))
	// 	return
	// }
	// init the stop channel
	wp.stop = make(chan struct{})
	stop := wp.stop
	// init the worker channel pool
	wp.workerChPool.New = func() any {
		return &workerCh{
			lastUseTime: time.Now(),
			ch:          make(chan *event.Work, wp.MaxWorkers),
			receiptCh:   make(chan *event.TxReceipt),
		}
	}

	// main loop
	go func() {
		for {
			select {
			case <-stop:
				log.Info(" ===== Worker pool stopped ===== ")
				return
			default:
				time.Sleep(wp.idleTimeout)
			}
		}
	}()

	log.Info("Worker pool started", "max-workers", wp.MaxWorkers)
}

// Stop stops the worker pool
func (wp *WorkerPool) Stop() {
	if wp.stop == nil {
		return
	}
	// close the stop channel
	close(wp.stop)
	wp.stop = nil
	// wait for all workers to finish
	wp.lock.Lock()
	defer wp.lock.Unlock()

	ready := wp.ready
	for i := range ready {
		ready[i].ch <- &event.Work{}
		ready[i] = nil
	}
	wp.ready = ready
	wp.canStop = true
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(h Handle, worker int, timeout time.Duration, cfg *txmgr.Config) *WorkerPool {
	return &WorkerPool{
		WorkerFunc:    h,
		MaxWorkers:    worker,
		receiptMgr:    txmgr.NewDefaultTxManager(cfg),
		idleTimeout:   timeout,
		activeWorkers: 0,
		canStop:       false,
		ready:         make([]*workerCh, 0),
		stop:          make(chan struct{}),
	}
}

// Serve is the main function of the worker pool. Add a **new work** to the WorkerCh.
func (wp *WorkerPool) Serve(work *event.Work, receipt chan *event.TxReceipt) bool {
	wc := wp.getWorkCh()
	if wc == nil {
		return false
	}
	wc.ch <- work
	wc.receiptCh = receipt
	return true
}

// getWorkCh returns a channel of workerCh.
// when we need to process new work, we need to get a new worker channel.
func (wp *WorkerPool) getWorkCh() *workerCh {
	var wc *workerCh
	shouldCreateWorker := false
	// lock
	wp.lock.Lock()
	//
	ready := wp.ready
	l := len(ready)
	if l == 0 {
		if wp.activeWorkers < wp.MaxWorkers {
			shouldCreateWorker = true
			wp.activeWorkers++
		}
	} else {
		wc = ready[l-1]
		ready[l-1] = nil
		wp.ready = ready[:l-1]
	}
	wp.lock.Unlock()
	//
	if wc == nil {
		if !shouldCreateWorker {
			return nil
		}
		poolWorker := wp.workerChPool.Get()
		wc = poolWorker.(*workerCh)
		go func() {
			wp.workerFunc(wc)
			wp.workerChPool.Put(poolWorker)
		}()
	}
	return wc
}

// release releases a workerCh to the worker pool(ready).
func (wp *WorkerPool) release(wc *workerCh) bool {
	// record time of the last use
	wc.lastUseTime = time.Now()
	// lock
	wp.lock.Lock()
	defer wp.lock.Unlock()
	// release the worker channel
	if wp.canStop {
		wp.lock.Unlock()
		return false
	}
	wp.ready = append(wp.ready, wc)
	return true
}

func (wp *WorkerPool) workerFunc(wc *workerCh) {
	// for
	for work := range wc.ch {
		if work == nil {
			break
		}
		log.Info("Workerpool working on", "l2txhash", work.L2txHash)

		if err := wp.WorkerFunc(work, wc.receiptCh, wp.receiptMgr); err != nil {
			log.Error("Error handling work", "err", err)
		}

		work = nil

		if !wp.release(wc) {
			break
		}
	}
	wp.lock.Lock()
	wp.activeWorkers--
	wp.lock.Unlock()
}
