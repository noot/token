pragma solidity ^0.4.24;

contract Token {
        event Transfer(address _to, address _from, uint _value);

        address public owner;

        constructor() {
                owner = msg.sender;
        }

        function transfer(address _to, uint _value) public {
                emit Transfer(_to, msg.sender, _value);
        }
}

