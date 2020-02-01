pragma solidity >=0.5.0;

/**
Function to receive approval and execute function in one call.
 */
contract TokenRecipient {
  function receiveApproval(address _from, uint256 _value, address _token, bytes memory _extraData) public;
}
