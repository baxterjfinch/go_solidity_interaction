// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lor

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// LorABI is the input ABI used to generate the binding from.
const LorABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLockedSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockedSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"setBank\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setLockedSupply\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferAnyERC20Token\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// LorBin is the compiled bytecode used for deploying new contracts.
var LorBin = "0x60806040523480156200001157600080fd5b506040518060400160405280601581526020017f4c616972204f776e6572736869702052696768747300000000000000000000008152506040518060400160405280600381526020017f4c4f520000000000000000000000000000000000000000000000000000000000815250600282600390805190602001906200009892919062000442565b508160049080519060200190620000b192919062000442565b5080600560006101000a81548160ff021916908360ff1602179055505050506000620000e2620001e760201b60201c565b905080600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35033600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620001d93364452dd267e4620001ef60201b60201c565b6000600781905550620004f1565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141562000293576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b620002af81600254620003b960201b620017401790919060201c565b6002819055506200030d816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054620003b960201b620017401790919060201c565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b60008082840190508381101562000438576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200048557805160ff1916838001178555620004b6565b82800160010185558215620004b6579182015b82811115620004b557825182559160200191906001019062000498565b5b509050620004c59190620004c9565b5090565b620004ee91905b80821115620004ea576000816000905550600101620004d0565b5090565b90565b611a6b80620005016000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c806376cdb03b116100b8578063a457c2d71161007c578063a457c2d7146105ae578063a9059cbb14610614578063ca5c7b911461067a578063d493b9ac14610698578063dd62ed3e1461071e578063f2fde38b1461079657610137565b806376cdb03b1461042f5780638da5cb5b146104795780638f32d59b146104c357806395d89b41146104e5578063a27851991461056857610137565b8063313ce567116100ff578063313ce567146103255780633950935114610349578063547199f2146103af57806370a08231146103cd578063715018a61461042557610137565b806306fdde031461013c578063090d23b9146101bf578063095ea7b31461021b57806318160ddd1461028157806323b872dd1461029f575b600080fd5b6101446107da565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610184578082015181840152602081019050610169565b50505050905090810190601f1680156101b15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610201600480360360208110156101d557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061087c565b604051808215151515815260200191505060405180910390f35b6102676004803603604081101561023157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610942565b604051808215151515815260200191505060405180910390f35b610289610960565b6040518082815260200191505060405180910390f35b61030b600480360360608110156102b557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061096a565b604051808215151515815260200191505060405180910390f35b61032d610a43565b604051808260ff1660ff16815260200191505060405180910390f35b6103956004803603604081101561035f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a5a565b604051808215151515815260200191505060405180910390f35b6103b7610b0d565b6040518082815260200191505060405180910390f35b61040f600480360360208110156103e357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b17565b6040518082815260200191505060405180910390f35b61042d610b5f565b005b610437610c9a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610481610cc0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6104cb610cea565b604051808215151515815260200191505060405180910390f35b6104ed610d49565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561052d578082015181840152602081019050610512565b50505050905090810190601f16801561055a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6105946004803603602081101561057e57600080fd5b8101908080359060200190929190505050610deb565b604051808215151515815260200191505060405180910390f35b6105fa600480360360408110156105c457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610e86565b604051808215151515815260200191505060405180910390f35b6106606004803603604081101561062a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610f53565b604051808215151515815260200191505060405180910390f35b610682610f71565b6040518082815260200191505060405180910390f35b610704600480360360608110156106ae57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610f77565b604051808215151515815260200191505060405180910390f35b6107806004803603604081101561073457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506110be565b6040518082815260200191505060405180910390f35b6107d8600480360360208110156107ac57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611145565b005b606060038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108725780601f1061084757610100808354040283529160200191610872565b820191906000526020600020905b81548152906001019060200180831161085557829003601f168201915b5050505050905090565b6000610886610cea565b6108f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b81600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060019050919050565b600061095661094f6111cb565b84846111d3565b6001905092915050565b6000600254905090565b60006109778484846113ca565b610a38846109836111cb565b610a33856040518060600160405280602881526020016119a060289139600160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006109e96111cb565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546116809092919063ffffffff16565b6111d3565b600190509392505050565b6000600560009054906101000a900460ff16905090565b6000610b03610a676111cb565b84610afe8560016000610a786111cb565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461174090919063ffffffff16565b6111d3565b6001905092915050565b6000600754905090565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b610b67610cea565b610bd9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610d2d6111cb565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b606060048054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610de15780601f10610db657610100808354040283529160200191610de1565b820191906000526020600020905b815481529060010190602001808311610dc457829003601f168201915b5050505050905090565b6000610df5610cea565b610e67576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600060075414610e7657600080fd5b8160078190555060019050919050565b6000610f49610e936111cb565b84610f4485604051806060016040528060258152602001611a116025913960016000610ebd6111cb565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546116809092919063ffffffff16565b6111d3565b6001905092915050565b6000610f67610f606111cb565b84846113ca565b6001905092915050565b60075481565b6000610f81610cea565b610ff3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561107a57600080fd5b505af115801561108e573d6000803e3d6000fd5b505050506040513d60208110156110a457600080fd5b810190808051906020019092919050505090509392505050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b61114d610cea565b6111bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6111c8816117c8565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611259576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806119ed6024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156112df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806119586022913960400191505060405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611450576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806119c86025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156114d6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602381526020018061190f6023913960400191505060405180910390fd5b6115418160405180606001604052806026815260200161197a602691396000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546116809092919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506115d4816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461174090919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b600083831115829061172d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156116f25780820151818401526020810190506116d7565b50505050905090810190601f16801561171f5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b6000808284019050838110156117be576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561184e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806119326026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505056fe45524332303a207472616e7366657220746f20746865207a65726f20616464726573734f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220addfdb786e05a60dba3875c6887ebed9647d993219d867412987c34495e57e1f64736f6c63430006020033"

// DeployLor deploys a new Ethereum contract, binding an instance of Lor to it.
func DeployLor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Lor, error) {
	parsed, err := abi.JSON(strings.NewReader(LorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lor{LorCaller: LorCaller{contract: contract}, LorTransactor: LorTransactor{contract: contract}, LorFilterer: LorFilterer{contract: contract}}, nil
}

// Lor is an auto generated Go binding around an Ethereum contract.
type Lor struct {
	LorCaller     // Read-only binding to the contract
	LorTransactor // Write-only binding to the contract
	LorFilterer   // Log filterer for contract events
}

// LorCaller is an auto generated read-only Go binding around an Ethereum contract.
type LorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LorSession struct {
	Contract     *Lor              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LorCallerSession struct {
	Contract *LorCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LorTransactorSession struct {
	Contract     *LorTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LorRaw is an auto generated low-level Go binding around an Ethereum contract.
type LorRaw struct {
	Contract *Lor // Generic contract binding to access the raw methods on
}

// LorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LorCallerRaw struct {
	Contract *LorCaller // Generic read-only contract binding to access the raw methods on
}

// LorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LorTransactorRaw struct {
	Contract *LorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLor creates a new instance of Lor, bound to a specific deployed contract.
func NewLor(address common.Address, backend bind.ContractBackend) (*Lor, error) {
	contract, err := bindLor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lor{LorCaller: LorCaller{contract: contract}, LorTransactor: LorTransactor{contract: contract}, LorFilterer: LorFilterer{contract: contract}}, nil
}

// NewLorCaller creates a new read-only instance of Lor, bound to a specific deployed contract.
func NewLorCaller(address common.Address, caller bind.ContractCaller) (*LorCaller, error) {
	contract, err := bindLor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LorCaller{contract: contract}, nil
}

// NewLorTransactor creates a new write-only instance of Lor, bound to a specific deployed contract.
func NewLorTransactor(address common.Address, transactor bind.ContractTransactor) (*LorTransactor, error) {
	contract, err := bindLor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LorTransactor{contract: contract}, nil
}

// NewLorFilterer creates a new log filterer instance of Lor, bound to a specific deployed contract.
func NewLorFilterer(address common.Address, filterer bind.ContractFilterer) (*LorFilterer, error) {
	contract, err := bindLor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LorFilterer{contract: contract}, nil
}

// bindLor binds a generic wrapper to an already deployed contract.
func bindLor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lor *LorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lor.Contract.LorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lor *LorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lor.Contract.LorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lor *LorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lor.Contract.LorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lor *LorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lor *LorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lor *LorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lor.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_Lor *LorCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Lor.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_Lor *LorSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Lor.Contract.Allowance(&_Lor.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_Lor *LorCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Lor.Contract.Allowance(&_Lor.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Lor *LorCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Lor.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Lor *LorSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Lor.Contract.BalanceOf(&_Lor.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Lor *LorCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Lor.Contract.BalanceOf(&_Lor.CallOpts, account)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() constant returns(address)
func (_Lor *LorCaller) Bank(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lor.contract.Call(opts, out, "bank")
	return *ret0, err
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() constant returns(address)
func (_Lor *LorSession) Bank() (common.Address, error) {
	return _Lor.Contract.Bank(&_Lor.CallOpts)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() constant returns(address)
func (_Lor *LorCallerSession) Bank() (common.Address, error) {
	return _Lor.Contract.Bank(&_Lor.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Lor *LorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Lor.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Lor *LorSession) Decimals() (uint8, error) {
	return _Lor.Contract.Decimals(&_Lor.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Lor *LorCallerSession) Decimals() (uint8, error) {
	return _Lor.Contract.Decimals(&_Lor.CallOpts)
}

// GetLockedSupply is a free data retrieval call binding the contract method 0x547199f2.
//
// Solidity: function getLockedSupply() constant returns(uint256)
func (_Lor *LorCaller) GetLockedSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Lor.contract.Call(opts, out, "getLockedSupply")
	return *ret0, err
}

// GetLockedSupply is a free data retrieval call binding the contract method 0x547199f2.
//
// Solidity: function getLockedSupply() constant returns(uint256)
func (_Lor *LorSession) GetLockedSupply() (*big.Int, error) {
	return _Lor.Contract.GetLockedSupply(&_Lor.CallOpts)
}

// GetLockedSupply is a free data retrieval call binding the contract method 0x547199f2.
//
// Solidity: function getLockedSupply() constant returns(uint256)
func (_Lor *LorCallerSession) GetLockedSupply() (*big.Int, error) {
	return _Lor.Contract.GetLockedSupply(&_Lor.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Lor *LorCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Lor.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Lor *LorSession) IsOwner() (bool, error) {
	return _Lor.Contract.IsOwner(&_Lor.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Lor *LorCallerSession) IsOwner() (bool, error) {
	return _Lor.Contract.IsOwner(&_Lor.CallOpts)
}

// LockedSupply is a free data retrieval call binding the contract method 0xca5c7b91.
//
// Solidity: function lockedSupply() constant returns(uint256)
func (_Lor *LorCaller) LockedSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Lor.contract.Call(opts, out, "lockedSupply")
	return *ret0, err
}

// LockedSupply is a free data retrieval call binding the contract method 0xca5c7b91.
//
// Solidity: function lockedSupply() constant returns(uint256)
func (_Lor *LorSession) LockedSupply() (*big.Int, error) {
	return _Lor.Contract.LockedSupply(&_Lor.CallOpts)
}

// LockedSupply is a free data retrieval call binding the contract method 0xca5c7b91.
//
// Solidity: function lockedSupply() constant returns(uint256)
func (_Lor *LorCallerSession) LockedSupply() (*big.Int, error) {
	return _Lor.Contract.LockedSupply(&_Lor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Lor *LorCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Lor.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Lor *LorSession) Name() (string, error) {
	return _Lor.Contract.Name(&_Lor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Lor *LorCallerSession) Name() (string, error) {
	return _Lor.Contract.Name(&_Lor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Lor *LorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lor.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Lor *LorSession) Owner() (common.Address, error) {
	return _Lor.Contract.Owner(&_Lor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Lor *LorCallerSession) Owner() (common.Address, error) {
	return _Lor.Contract.Owner(&_Lor.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Lor *LorCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Lor.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Lor *LorSession) Symbol() (string, error) {
	return _Lor.Contract.Symbol(&_Lor.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Lor *LorCallerSession) Symbol() (string, error) {
	return _Lor.Contract.Symbol(&_Lor.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Lor *LorCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Lor.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Lor *LorSession) TotalSupply() (*big.Int, error) {
	return _Lor.Contract.TotalSupply(&_Lor.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Lor *LorCallerSession) TotalSupply() (*big.Int, error) {
	return _Lor.Contract.TotalSupply(&_Lor.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Lor *LorTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lor.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Lor *LorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lor.Contract.Approve(&_Lor.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Lor *LorTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lor.Contract.Approve(&_Lor.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Lor *LorTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Lor.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Lor *LorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Lor.Contract.DecreaseAllowance(&_Lor.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Lor *LorTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Lor.Contract.DecreaseAllowance(&_Lor.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Lor *LorTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Lor.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Lor *LorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Lor.Contract.IncreaseAllowance(&_Lor.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Lor *LorTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Lor.Contract.IncreaseAllowance(&_Lor.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lor *LorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lor.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lor *LorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lor.Contract.RenounceOwnership(&_Lor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lor *LorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lor.Contract.RenounceOwnership(&_Lor.TransactOpts)
}

// SetBank is a paid mutator transaction binding the contract method 0x090d23b9.
//
// Solidity: function setBank(address newAddress) returns(bool success)
func (_Lor *LorTransactor) SetBank(opts *bind.TransactOpts, newAddress common.Address) (*types.Transaction, error) {
	return _Lor.contract.Transact(opts, "setBank", newAddress)
}

// SetBank is a paid mutator transaction binding the contract method 0x090d23b9.
//
// Solidity: function setBank(address newAddress) returns(bool success)
func (_Lor *LorSession) SetBank(newAddress common.Address) (*types.Transaction, error) {
	return _Lor.Contract.SetBank(&_Lor.TransactOpts, newAddress)
}

// SetBank is a paid mutator transaction binding the contract method 0x090d23b9.
//
// Solidity: function setBank(address newAddress) returns(bool success)
func (_Lor *LorTransactorSession) SetBank(newAddress common.Address) (*types.Transaction, error) {
	return _Lor.Contract.SetBank(&_Lor.TransactOpts, newAddress)
}

// SetLockedSupply is a paid mutator transaction binding the contract method 0xa2785199.
//
// Solidity: function setLockedSupply(uint256 amount) returns(bool success)
func (_Lor *LorTransactor) SetLockedSupply(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Lor.contract.Transact(opts, "setLockedSupply", amount)
}

// SetLockedSupply is a paid mutator transaction binding the contract method 0xa2785199.
//
// Solidity: function setLockedSupply(uint256 amount) returns(bool success)
func (_Lor *LorSession) SetLockedSupply(amount *big.Int) (*types.Transaction, error) {
	return _Lor.Contract.SetLockedSupply(&_Lor.TransactOpts, amount)
}

// SetLockedSupply is a paid mutator transaction binding the contract method 0xa2785199.
//
// Solidity: function setLockedSupply(uint256 amount) returns(bool success)
func (_Lor *LorTransactorSession) SetLockedSupply(amount *big.Int) (*types.Transaction, error) {
	return _Lor.Contract.SetLockedSupply(&_Lor.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Lor *LorTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lor.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Lor *LorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lor.Contract.Transfer(&_Lor.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Lor *LorTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lor.Contract.Transfer(&_Lor.TransactOpts, recipient, amount)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xd493b9ac.
//
// Solidity: function transferAnyERC20Token(address tokenAddress, address toAddress, uint256 tokens) returns(bool success)
func (_Lor *LorTransactor) TransferAnyERC20Token(opts *bind.TransactOpts, tokenAddress common.Address, toAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Lor.contract.Transact(opts, "transferAnyERC20Token", tokenAddress, toAddress, tokens)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xd493b9ac.
//
// Solidity: function transferAnyERC20Token(address tokenAddress, address toAddress, uint256 tokens) returns(bool success)
func (_Lor *LorSession) TransferAnyERC20Token(tokenAddress common.Address, toAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Lor.Contract.TransferAnyERC20Token(&_Lor.TransactOpts, tokenAddress, toAddress, tokens)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xd493b9ac.
//
// Solidity: function transferAnyERC20Token(address tokenAddress, address toAddress, uint256 tokens) returns(bool success)
func (_Lor *LorTransactorSession) TransferAnyERC20Token(tokenAddress common.Address, toAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Lor.Contract.TransferAnyERC20Token(&_Lor.TransactOpts, tokenAddress, toAddress, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Lor *LorTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lor.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Lor *LorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lor.Contract.TransferFrom(&_Lor.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Lor *LorTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lor.Contract.TransferFrom(&_Lor.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lor *LorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Lor.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lor *LorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lor.Contract.TransferOwnership(&_Lor.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lor *LorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lor.Contract.TransferOwnership(&_Lor.TransactOpts, newOwner)
}

// LorApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Lor contract.
type LorApprovalIterator struct {
	Event *LorApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LorApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LorApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LorApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LorApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LorApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LorApproval represents a Approval event raised by the Lor contract.
type LorApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Lor *LorFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LorApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Lor.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LorApprovalIterator{contract: _Lor.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Lor *LorFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LorApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Lor.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LorApproval)
				if err := _Lor.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Lor *LorFilterer) ParseApproval(log types.Log) (*LorApproval, error) {
	event := new(LorApproval)
	if err := _Lor.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Lor contract.
type LorOwnershipTransferredIterator struct {
	Event *LorOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LorOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LorOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LorOwnershipTransferred represents a OwnershipTransferred event raised by the Lor contract.
type LorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lor *LorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lor.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LorOwnershipTransferredIterator{contract: _Lor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lor *LorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lor.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LorOwnershipTransferred)
				if err := _Lor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lor *LorFilterer) ParseOwnershipTransferred(log types.Log) (*LorOwnershipTransferred, error) {
	event := new(LorOwnershipTransferred)
	if err := _Lor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LorTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Lor contract.
type LorTransferIterator struct {
	Event *LorTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LorTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LorTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LorTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LorTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LorTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LorTransfer represents a Transfer event raised by the Lor contract.
type LorTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Lor *LorFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LorTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lor.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LorTransferIterator{contract: _Lor.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Lor *LorFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LorTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lor.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LorTransfer)
				if err := _Lor.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Lor *LorFilterer) ParseTransfer(log types.Log) (*LorTransfer, error) {
	event := new(LorTransfer)
	if err := _Lor.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
