pragma solidity >=0.5.0;

// ----------------------------------------------------------------------------
// Entity:          Dragon Den
// Token Name:      Matter
// symbol:          MTR
// MaxSupply:       433494437
// Mintable Chunk:  1000
// Authored By:     Baxter Finch, Lance Rogers & Eric Tesenair
// Description:     Custom minting function with
//                  time based and ETH address verification
// ----------------------------------------------------------------------------

import "./dependencies/token/ERC20/ERC20.sol";
import "./dependencies/token/ERC20/ERC20Detailed.sol";
import "./dependencies/ownership/Ownable.sol";
import "./dependencies/math/SafeMath.sol";

contract Matter is ERC20, ERC20Detailed, Ownable {
    /*
    * @title Matter Contract
    * @author Baxter Finch, Lance Rogers & Eric Tesenair
    * @notice ERC20 contract with custom mint function
    */

    using SafeMath for uint256;

    uint8 public mintHour;
    address public bank;
    uint256 public maxSupply;
    uint256 private _lastMintedBlockTimestamp;
    uint256 private _mintMTRChunk;
    uint256 private _minMintInterval;
    uint256 constant private SECONDS_PER_DAY = 86400;
    uint256 constant private SECONDS_PER_HOUR = 3600;

    constructor() public ERC20Detailed("Matter", "MTR", 18) {
        bank = msg.sender;
        mintHour = 21;
        maxSupply = 433494437000000000000000000;
        _mintMTRChunk = 1000000000000000000000;
        _lastMintedBlockTimestamp = 1;
        _minMintInterval = 36000; // Seconds
    }

    /*
     * @notice Get Minimum Mint Interval
     */
    function minMintInterval() public view returns (uint256) {
        return _minMintInterval;
    }

    /*
     * @notice Get last minted block timestamp
     */
    function lastMinted() public view returns (uint256) {
        return _lastMintedBlockTimestamp;
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
     * @notice set mintHour
     * @notice must be contract owner
     * @param newMintHour An hour from 1-23 UTC where minting will be alloud to occur
     */
    function setMintHour(uint8 newMintHour) public onlyOwner returns (bool success) {
        mintHour = newMintHour;
        return true;
    }

    /*
     * @notice Returns the hour from a given epoch
     * @param timestamp An epoch time value
     */
    function getHour(uint256 timestamp) internal pure returns (uint hour) {
        uint256 secs = timestamp.mod(SECONDS_PER_DAY);
        hour = secs.div(SECONDS_PER_HOUR);
    }

    /*
     * @notice Will mint 1000 MTR
     * @notice only if the _totalSupply + _mintMTRChunk doesnt exceed the _MaxSupply
     * @param _timestamp The current timestamp where minting was first triggered
     * @param _to The address that will receive the coin.
     */

    function mint(uint256 _timestamp) private returns (bool) {
        require(totalSupply().add(_mintMTRChunk) <= maxSupply, "MTR: Cannot mint more than max supply"); // overflow checks
        _mint(bank, _mintMTRChunk);
        _lastMintedBlockTimestamp = _timestamp;
        return true;
    }

    /*
     * @notice Time-based function that only allows callers to mint if a certain amount of time has passed
     * @notice and only if the transaction was created in the valid mintHour
     */
    function mint() public returns (bool) {
        uint256 thisTimestamp = block.timestamp;
        require(thisTimestamp.sub(_lastMintedBlockTimestamp) > _minMintInterval, "MTR: The required time period has not passed");
        uint256 currentHour = getHour(thisTimestamp);
        require(currentHour == mintHour, "MTR: Can only mint during the minting hour");
        mint(thisTimestamp);
        return true;
    }

    /*
     * @notice Owner can transfer out any accidentally sent ERC20 tokens
     */
    function transferAnyERC20Token(address tokenAddress, address toAddress, uint256 tokens) public onlyOwner returns (bool success) {
        return ERC20(tokenAddress).transfer(toAddress, tokens);
    }
}
