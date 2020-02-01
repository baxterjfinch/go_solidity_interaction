pragma solidity >=0.5.0;

// ----------------------------------------------------------------------------
// Entity:          Dragon Den
// Token Name:      Lair Ownership Rights
// symbol:          LOR
// MaxSupply:       2971215073
// Authored By:     Baxter Finch, Lance Rogers & Eric Tesenair
// Description:
//
// ----------------------------------------------------------------------------

import "./dependencies/token/ERC20/ERC20.sol";
import "./dependencies/token/ERC20/ERC20Detailed.sol";
import "./dependencies/ownership/Ownable.sol";
import "./dependencies/math/SafeMath.sol";

contract LairOwnershipRights is ERC20, ERC20Detailed, Ownable {
    /*
    * @title Matter Contract
    * @author Baxter Finch, Lance Rogers & Eric Tesenair
    * @notice ERC20 contract with custom mint function
    */

    using SafeMath for uint256;

    address public bank;
    uint256 public lockedSupply;

    constructor() public ERC20Detailed("Lair Ownership Rights", "LOR", 2) {
        bank = msg.sender;
        _mint(msg.sender, 297121507300);
        lockedSupply = 0;
    }

    /*
     * @notice set bank address
     * @notice must be contract owner
     * @param newAddress The new bank address where freshly minted MTR will be sent
     */
    function setBank(address newAddress) public onlyOwner returns (bool success) {
        bank = newAddress;
        return true;
    }

    /*
     * @notice set locked supply
     * @notice must be contract owner
     */
    function setLockedSupply(uint256 amount) public onlyOwner returns (bool success) {
        require(lockedSupply == 0);
        lockedSupply = amount;
        return true;
    }

    function getLockedSupply() public view returns (uint256) {
        return lockedSupply;
    }

    /*
     * @notice Owner can transfer out any accidentally sent ERC20 tokens
     */
    function transferAnyERC20Token(address tokenAddress, address toAddress, uint256 tokens) public onlyOwner returns (bool success) {
        return ERC20(tokenAddress).transfer(toAddress, tokens);
    }
}
