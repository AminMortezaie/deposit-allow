// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract DepositAllowance {

    uint256 public constant ttl = 1 hours;

    enum Status {
        Pending,
        Deposited,
        Submitted,
        Revoked,
        Executed,
        Withdrawn
    }

    struct Transaction {
        address userA;
        address userB;
        address destination;
        address tokenA;
        address tokenB;
        uint256 amountA;
        uint256 amountB;
        Status statusA;
        Status statusB;
        uint256 timeLock;
    }

    event Deposit(address, address, uint256);
    event Submit(address, address, uint256);
    event Execute(address, address, uint256, address, address, uint256);
    event Withdraw(address, address, uint256);
    event Revoke(address, address, uint256);


    mapping(bytes32 => Transaction[]) txSet;


    function hashUserAddresses(address _userA, address _userB) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(_userA, _userB));
    }


    function initiateTransaction(
        address _userA,
        address _userB,
        address _destination,
        address _tokenA,
        address _tokenB,
        uint256 _amountA,
        uint256 _amountB
    ) external {
        bytes32 txKey = hashUserAddresses(_userA, _userB);

        Transaction memory newTransaction = Transaction({
            userA: _userA,
            userB: _userB,
            destination: _destination,
            tokenA: _tokenA,
            tokenB: _tokenB,
            amountA: _amountA,
            amountB: _amountB,
            statusA: Status.Pending,
            statusB: Status.Pending,
            timeLock: 0
        });

        txSet[txKey].push(newTransaction);
    }


    function deposit(bytes32 _txKey, uint _index) external {
        Transaction storage transaction = txSet[_txKey][_index];
        require(msg.sender == transaction.userA || msg.sender == transaction.userB, "You're not allowed.");
        bool isUserA = (msg.sender == transaction.userA);
        if (isUserA) {
            require(IERC20(transaction.tokenA).transferFrom(msg.sender, address(this), transaction.amountA), "Insufficient Allowance.");
            require(transaction.statusA == Status.Pending, "Transaction status not in the expected state.");
            transaction.statusA = Status.Deposited;
        } else {
            require(IERC20(transaction.tokenB).transferFrom(msg.sender, address(this), transaction.amountB), "Insufficient Allowance.");
            require(transaction.statusB == Status.Pending, "Transaction status not in the expected state.");
            transaction.statusB = Status.Deposited;
        }

        if (transaction.statusA == Status.Deposited && transaction.statusB == Status.Deposited) {
            transaction.timeLock = block.timestamp + ttl;
        }

        emit Deposit(msg.sender, isUserA ? transaction.tokenA : transaction.tokenB, isUserA ? transaction.amountA : transaction.amountB);
    }


    function submit(bytes32 _txKey, uint _index) external {
        Transaction storage transaction = txSet[_txKey][_index];
        require(msg.sender == transaction.userA || msg.sender == transaction.userB, "You're not allowed.");
        require(transaction.timeLock > block.timestamp, "Transaction is past the lock time.");
        if (msg.sender == transaction.userA) {
            require(transaction.statusA == Status.Pending, "Transaction status not in the expected state.");
            transaction.statusA = Status.Submitted;
            emit Submit(msg.sender, transaction.tokenA, transaction.amountA);
        } else {
            require(transaction.statusB == Status.Pending, "Transaction status not in the expected state.");
            transaction.statusB = Status.Submitted;
            emit Submit(msg.sender, transaction.tokenB, transaction.amountB);
        }

        if (transaction.statusA == Status.Submitted && transaction.statusB == Status.Submitted) {
            IERC20(transaction.tokenA).transfer(transaction.destination, transaction.amountA);
            IERC20(transaction.tokenB).transfer(transaction.destination, transaction.amountB);
            emit Execute(
                transaction.userA, transaction.tokenA, transaction.amountA,
                transaction.userB, transaction.tokenB, transaction.amountB
            );
        }
    }


    function revoke(bytes32 _txKey, uint256 _index) external {
        Transaction storage transaction = txSet[_txKey][_index];
        require(msg.sender == transaction.userA || msg.sender == transaction.userB, "You're not allowed.");
        if (msg.sender == transaction.userA) {
            require(transaction.statusA == Status.Submitted, "Transaction status not in the expected state.");
            transaction.statusA = Status.Revoked;
            emit Revoke(msg.sender, transaction.tokenA, transaction.amountA);
        } else {
            require(transaction.statusB == Status.Submitted, "Transaction status not in the expected state.");
            transaction.statusB = Status.Revoked;
            emit Revoke(msg.sender, transaction.tokenB, transaction.amountB);
        }
    }

    function withdraw(bytes32 _txKey, uint256 _index) external {
        Transaction storage transaction = txSet[_txKey][_index];
        require(msg.sender == transaction.userA || msg.sender == transaction.userB, "You're not allowed.");
        require(block.timestamp >= transaction.timeLock, "Timelock not yet expired.");
        require(transaction.statusA != Status.Executed && transaction.statusB != Status.Executed, "Transaction already executed.");

        if (msg.sender == transaction.userA) {
            require(transaction.statusA == Status.Submitted, "Transaction status not in the expected state.");
            transaction.statusA = Status.Withdrawn;
            IERC20(transaction.tokenA).transfer(msg.sender, transaction.amountA);
            emit Withdraw(msg.sender, transaction.tokenA, transaction.amountA);
        } else {
            require(transaction.statusB == Status.Submitted, "Transaction status not in the expected state.");
            transaction.statusB = Status.Withdrawn;
            IERC20(transaction.tokenB).transfer(msg.sender, transaction.amountB);
            emit Withdraw(msg.sender, transaction.tokenB, transaction.amountB);
        }
    }
}
