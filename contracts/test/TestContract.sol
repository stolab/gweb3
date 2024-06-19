// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract HelloWorld {

    string private Helloworld;

    constructor() {
        Helloworld = "Hello World !";
    }

    function getMessage() public view returns(string memory){
        return Helloworld;
    }

    function setMessage(string calldata message) public {
        Helloworld = message;
    }
}
