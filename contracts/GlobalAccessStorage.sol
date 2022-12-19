// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract GlobalAccessAdminStorage{
    address public globalAccessImplementation;
    address public pendingGlobalAccessImplementation;

    mapping(address => uint256) public administrators;
    address[] administratorList;

    mapping(address => uint256) public operators;
    address[] operatorList;

}

contract GlobalAccessV1Storage is GlobalAccessAdminStorage{
    enum Status{
        SEND_AND_RECEIVE,
        RECEIVE_ONLY,
        SEND_ONLY,
        NO_SEND_AND_RECEIVE
    }

    struct UserInfo{
        uint256 idx;
        uint256 typ;   //0: personal; 1: company; 2:financial company; ...
        uint256 level;  //0: leve1; 1:level2; ....
    }

    struct UserStatus {
        Status status;
        uint256 pauseFrom; //timestamp, till which specified operation is blocked
        uint256 pausedTo;  //timestamp, till which specified operation is blocked
    }

    mapping(uint256 => mapping(uint256 =>uint256)) public quotas;  //type=>level=>quota, record kyc's quota
    mapping(address => UserInfo) public users;                      //record user's kyc info,
    address[] public usersList;                //userList by type
    mapping(address => UserStatus) public userStatus;
}

