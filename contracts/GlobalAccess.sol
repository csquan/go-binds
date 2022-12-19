// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./GlobalAccessStorage.sol";
import "./GlobalAccessProxy.sol";


contract GlobalAccess is GlobalAccessV1Storage {
    //events
    event SetUserInfo(address indexed, uint256, uint256);
    event SetUserStatus(address indexed, Status indexed, uint256, uint256);
    event SetQuota(uint256 indexed, uint256 indexed, uint256);

    modifier onlyOperator(address _operator){
        require(operators[_operator] != 0, "only operator allowed");
        _;
    }

    function become(GlobalAccessProxy _globalAccessProxy) public {
        require(_globalAccessProxy.administrators(msg.sender) != 0, "only administrator allowed");
        require(_globalAccessProxy.acceptImplementation() == 0, "fail to accept implementation");
    }


    function setQuota(uint256 _type, uint256 _level, uint256 _quota) onlyOperator(msg.sender) external{
        quotas[_type][_level] = _quota;
        emit SetQuota(_type, _level,_quota);
    }

    function setUserInfo(address _account, uint256 _type, uint256 _level) onlyOperator(msg.sender) external {
        UserInfo memory userInfo;
        if (users[_account].idx == 0){
            usersList.push(_account);
            userInfo.idx = usersList.length;
        }
        userInfo.typ = _type;
        userInfo.level = _level;

        users[_account] = userInfo;
        emit SetUserInfo(_account, _type, _level);
    }


    function setUserStatus(address _account, Status _status, uint256 _pausedFrom, uint256 _pausedTo) onlyOperator(msg.sender) external {
        require(users[_account].idx != 0, "user not exist");
        require(_pausedFrom < _pausedTo, "paused start time > end time");

        UserStatus memory sta;
        sta.status = _status;
        sta.pauseFrom = _pausedFrom;
        sta.pausedTo =  _pausedTo;

        userStatus[_account] = sta;
        emit SetUserStatus(_account, _status, _pausedFrom, _pausedTo);
    }

    function getUserNumber() external view returns(uint256){
        return usersList.length;
    }

    function getUserInfo(address _account) public view returns(UserInfo memory){
        return users[_account];
    }

    function getUserStatus(address _account) public view returns(UserStatus memory){
        return userStatus[_account];
    }

}
