pragma solidity >=0.5.0;

import "./ERC20Interface.sol";
import "./TokenRecipient.sol";
import "../../ownership/Ownable.sol";

/**
Token implement
 */
contract Token is ERC20Interface, Ownable {

  mapping (address => uint256) public balances;
  mapping (address => mapping (address => uint256)) public allowed;

  // This notifies clients about the amount burnt
  event Burn(address indexed from, uint256 value);

  function balanceOf(address _owner) public view returns (uint256 balance) {
    return balances[_owner];
  }

  function transfer(address _to, uint256 _value) public returns (bool success) {
    _transfer(msg.sender, _to, _value);
    return true;
  }

  function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {
    require(_value <= allowed[_from][msg.sender]);
    allowed[_from][msg.sender] -= _value;
    _transfer(_from, _to, _value);
    return true;
  }

  function approve(address _spender, uint256 _value) public returns (bool success) {
    allowed[msg.sender][_spender] = _value;
    emit Approval(msg.sender, _spender, _value);
    return true;
  }

  function allowance(address _owner, address _spender) public view returns (uint256 remaining) {
    return allowed[_owner][_spender];
  }

  /**
  Owner can transfer out any accidentally sent ERC20 tokens
   */
  function transferAnyERC20Token(address tokenAddress, uint tokens) public onlyOwner returns (bool success) {
    return ERC20Interface(tokenAddress).transfer(owner, tokens);
  }

  /**
  Approves and then calls the receiving contract
   */
  function approveAndCall(address _spender, uint256 _value, bytes memory _extraData) public returns (bool success) {
    TokenRecipient spender = TokenRecipient(_spender);
    approve(_spender, _value);
    spender.receiveApproval(msg.sender, _value, this, _extraData);
    return true;
  }

  /**
  Destroy tokens.
  Remove `_value` tokens from the system irreversibly
    */
  function burn(uint256 _value) public returns (bool success) {
    require(balances[msg.sender] >= _value);
    balances[msg.sender] -= _value;
    totalSupply -= _value;
    emit Burn(msg.sender, _value);
    return true;
  }

  /**
  Destroy tokens from other account.
  Remove `_value` tokens from the system irreversibly on behalf of `_from`.
    */
  function burnFrom(address _from, uint256 _value) public returns (bool success) {
    require(balances[_from] >= _value);
    require(_value <= allowed[_from][msg.sender]);
    balances[_from] -= _value;
    allowed[_from][msg.sender] -= _value;
    totalSupply -= _value;
    emit Burn(_from, _value);
    return true;
  }

  /**
  Internal transfer, only can be called by this contract
    */
  function _transfer(address _from, address _to, uint _value) internal {
    // Prevent transfer to 0x0 address. Use burn() instead
    require(_to != 0x0);
    // Check if the sender has enough
    require(balances[_from] >= _value);
    // Check for overflows
    require(balances[_to] + _value > balances[_to]);
    // Save this for an assertion in the future
    uint previousBalances = balances[_from] + balances[_to];
    // Subtract from the sender
    balances[_from] -= _value;
    // Add the same to the recipient
    balances[_to] += _value;
    emit Transfer(_from, _to, _value);
    // Asserts are used to use static analysis to find bugs in your code. They should never fail
    assert(balances[_from] + balances[_to] == previousBalances);
  }

}
