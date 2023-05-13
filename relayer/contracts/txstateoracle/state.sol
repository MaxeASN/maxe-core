// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
import "../interfaces/IEntryPoint.sol";
import "../interfaces/IAccount.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
contract TxState is Ownable {
    address public DEPOSITOR_ACCOUNT;
    IEntryPoint private immutable _entryPoint;
    
    enum State {GENERATED, SENT, PENDING, SUCCESSFUL, FAILED}
    struct TransactionInfo{
        uint64 chainId;  //L1链ID
        address from;    //在L1交易发起地址
        uint64 seqNum;   //from账户下交易序号
        address receiver; //L1交易接收地址
        uint256 amount;   //交易的金额大小
        State   state;   //交易的状态
        bytes   data;    //交易携带的合约调用数据 
    }
    //账户在L2上交易数
    //mapping (address=>uint64) public SequenceNumber;
    //存储交易哈希
    mapping(address=>mapping(uint64=>bytes)) public StoreL1Txhash;

   
    constructor(IEntryPoint anEntryPoint ){
        _entryPoint = anEntryPoint;
    }

    //mapping(bytes32 => State) public L1txHashToState;
    event L1transferEvent(address indexed account, address indexed from,uint64 indexed seqNum,TransactionInfo txInfo);
    event updateTxStateSuccess(bytes indexed L1Txhash, uint state);
     //获取entryPoint地址
    function entryPoint() public view   returns (IEntryPoint) {
        return _entryPoint;
    }
   /**
     * ensure the request comes from the known entrypoint.
     */
    function _requireFromEntryPoint() internal virtual view {
        require(msg.sender == address(entryPoint()), "account: not from EntryPoint");
    }
    
    
    //更新L1交易状态
    //后端通过from和seqNum获取到对应的账户合约地址，然后调用账户合约的updateTxState方法，更新交易状态
    //question:后端发起的TxState call Gas费怎么收
    function setL1TxState(bytes calldata _txHash,address l2Account, address _from,uint64 _seqNum,uint _state) public payable onlyOwner{
        
        StoreL1Txhash[_from][_seqNum]=_txHash;
        try IAccount(l2Account).updateTxState{gas: gasleft()}(_from,_seqNum,_state){
            emit updateTxStateSuccess(_txHash,_state);   
        } catch {
            revert("updateTxState failed");
        }
    }
    //发出交易事件
    function proposeTxToL1(uint64 _chainId,
        address _from,
        uint64 _seqNum,
        address _receiver,
        uint256 _value,
        bytes memory _data)  external {
      
      TransactionInfo memory txInfo=(TransactionInfo(_chainId,_from,_seqNum ,_receiver,_value,State.GENERATED,_data));
      emit L1transferEvent(msg.sender,_from,_seqNum,txInfo);
      
      
    }
    //获取L1交易hash
    function getL1Txhash(address _from,uint64 _seqNum) public view returns (bytes memory ){
        return StoreL1Txhash[_from][_seqNum];
    }
}
