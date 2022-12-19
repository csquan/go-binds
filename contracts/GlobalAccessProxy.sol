// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./GlobalAccessStorage.sol";

contract GlobalAccessProxy is GlobalAccessAdminStorage{
    event NewPendingImplementation(address newPendingImplementation);
    event NewImplementation(address oldImplementation, address newImplementation);

    event AddAdministrator(address indexed, address indexed);
    event RemoveAdministrator(address indexed, address indexed);
    event AddOperator(address indexed, address indexed);
    event RemoveOperator(address indexed, address indexed);

    modifier onlyAdministrator(address _admin){
        require(administrators[_admin] != 0, "only administrator allowed");
        _;
    }

    constructor(address[] memory _administrators){
        require(_administrators.length > 0, "no administrator");

        uint256 adminNumber = _administrators.length;
        for (uint256 i= 0; i< adminNumber - 1 ; i++){
            for (uint256 j = i+1; j< adminNumber; j++){
                require(_administrators[i] != _administrators[j], "duplicated administrator");
            }
        }

        for (uint256 i= 0; i< adminNumber; i++){
            administratorList.push(_administrators[i]);
            administrators[_administrators[i]] = administratorList.length; //index from 1
            emit AddAdministrator(_administrators[i], msg.sender);
        }
    }

    function addAdministrator(address _newAdmin) onlyAdministrator(msg.sender) external returns(uint256){
        require(administrators[_newAdmin] == 0, "administrator already exist");
        require(operators[_newAdmin] == 0, "administrator can not be operator");

        administratorList.push(_newAdmin);
        administrators[_newAdmin] = administratorList.length;
        emit AddAdministrator(_newAdmin, msg.sender);
        return  administrators[_newAdmin];
    }

    function removeAdministrator(address _newAdmin) onlyAdministrator(msg.sender) external{
        require(administrators[_newAdmin] != 0, "administrator not exist");
        require(administratorList.length > 1, "at least one administrator");

        uint256 index = administrators[_newAdmin];
        uint256 adminNumber = administratorList.length;

        if (index == adminNumber){
            administratorList.pop();
        }else{
            address lastAdmin = administratorList[adminNumber-1];
            administratorList[index-1] = lastAdmin;
            administrators[lastAdmin] = index;
            administratorList.pop();
        }
        delete administrators[_newAdmin];
        emit AddAdministrator(_newAdmin, msg.sender);
    }

    function addOperator(address _newOperator) onlyAdministrator(msg.sender) external returns(uint256) {
        require(operators[_newOperator] == 0, "operator already exist");
        require(administrators[_newOperator] == 0, "operator can not be administrator");

        operatorList.push(_newOperator);
        operators[_newOperator] = operatorList.length;
        emit AddOperator(_newOperator, msg.sender);
        return operators[_newOperator];
    }

    function removeOperator(address _newOperator) onlyAdministrator(msg.sender) external{
        require(operators[_newOperator] != 0, "operator not exist");

        uint256 index = operators[_newOperator];
        uint256 operatorNumber = operatorList.length;

        if (index == operatorNumber){
            operatorList.pop();
        }else{
            address lastOperator = operatorList[operatorNumber-1];
            operatorList[index-1] = lastOperator;
            operators[lastOperator] = index;
            operatorList.pop();
        }
        delete operators[_newOperator];
        emit AddOperator(_newOperator, msg.sender);
    }


    function setPendingImplementation(address _newPendingImplementation) external onlyAdministrator(msg.sender) {
        pendingGlobalAccessImplementation = _newPendingImplementation;
        emit NewPendingImplementation(_newPendingImplementation);
    }

    function acceptImplementation() external returns(uint256){
        // Check caller is pendingImplementation and pendingImplementation ≠ address(0)
        require(msg.sender == pendingGlobalAccessImplementation && pendingGlobalAccessImplementation != address(0),
            "incorrect pendingGlobalAccessImplementation");

        // Save current values for inclusion in log
        address oldImplementation = globalAccessImplementation;
        globalAccessImplementation = pendingGlobalAccessImplementation;
        pendingGlobalAccessImplementation = address(0);
        emit NewImplementation(oldImplementation, globalAccessImplementation);

        return 0;
    }

    function getAdministrators() external view returns(address[] memory){
        return administratorList;
    }

    function getOperators() external view returns (address[] memory){
        return operatorList;
    }
    
    fallback() external {
        // delegate all other functions to current implementation
        (bool success, ) = globalAccessImplementation.delegatecall(msg.data);

        assembly {
            let free_mem_ptr := mload(0x40)
            returndatacopy(free_mem_ptr, 0, returndatasize())

            switch success
            case 0 { revert(free_mem_ptr, returndatasize()) }
            default { return(free_mem_ptr, returndatasize()) }
        }
    }



}
