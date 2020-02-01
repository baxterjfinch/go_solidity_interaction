// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package matter

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

// MatterABI is the input ABI used to generate the binding from.
const MatterABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minMintInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintHour\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"setBank\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newMintHour\",\"type\":\"uint8\"}],\"name\":\"setMintHour\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferAnyERC20Token\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MatterBin is the compiled bytecode used for deploying new contracts.
var MatterBin = "0x60806040523480156200001157600080fd5b506040518060400160405280600681526020017f4d617474657200000000000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f4d54520000000000000000000000000000000000000000000000000000000000815250601282600390805190602001906200009892919062000220565b508160049080519060200190620000b192919062000220565b5080600560006101000a81548160ff021916908360ff1602179055505050506000620000e26200021860201b60201c565b905080600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35033600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506015600560156101000a81548160ff021916908360ff1602179055506b01669404848dd26da7740000600781905550683635c9adc5dea000006009819055506001600881905550618ca0600a81905550620002cf565b600033905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200026357805160ff191683800117855562000294565b8280016001018555821562000294579182015b828111156200029357825182559160200191906001019062000276565b5b509050620002a39190620002a7565b5090565b620002cc91905b80821115620002c8576000816000905550600101620002ae565b5090565b90565b61219c80620002df6000396000f3fe608060405234801561001057600080fd5b50600436106101585760003560e01c80638da5cb5b116100c3578063b0601abe1161007c578063b0601abe1461069b578063bbc19ab0146106e4578063d493b9ac14610702578063d5abeb0114610788578063dd62ed3e146107a6578063f2fde38b1461081e57610158565b80638da5cb5b146104bc5780638f32d59b14610506578063942cd61b1461052857806395d89b411461054c578063a457c2d7146105cf578063a9059cbb1461063557610158565b8063313ce56711610115578063313ce56714610368578063395093511461038c5780635620583c146103f257806370a0823114610410578063715018a61461046857806376cdb03b1461047257610158565b806306fdde031461015d578063090d23b9146101e0578063095ea7b31461023c5780631249c58b146102a257806318160ddd146102c457806323b872dd146102e2575b600080fd5b610165610862565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101a557808201518184015260208101905061018a565b50505050905090810190601f1680156101d25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610222600480360360208110156101f657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610904565b604051808215151515815260200191505060405180910390f35b6102886004803603604081101561025257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506109ca565b604051808215151515815260200191505060405180910390f35b6102aa6109e8565b604051808215151515815260200191505060405180910390f35b6102cc610ae6565b6040518082815260200191505060405180910390f35b61034e600480360360608110156102f857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610af0565b604051808215151515815260200191505060405180910390f35b610370610bc9565b604051808260ff1660ff16815260200191505060405180910390f35b6103d8600480360360408110156103a257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610be0565b604051808215151515815260200191505060405180910390f35b6103fa610c93565b6040518082815260200191505060405180910390f35b6104526004803603602081101561042657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c9d565b6040518082815260200191505060405180910390f35b610470610ce5565b005b61047a610e20565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6104c4610e46565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61050e610e70565b604051808215151515815260200191505060405180910390f35b610530610ecf565b604051808260ff1660ff16815260200191505060405180910390f35b610554610ee2565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610594578082015181840152602081019050610579565b50505050905090810190601f1680156105c15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61061b600480360360408110156105e557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610f84565b604051808215151515815260200191505060405180910390f35b6106816004803603604081101561064b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611051565b604051808215151515815260200191505060405180910390f35b6106ca600480360360208110156106b157600080fd5b81019080803560ff16906020019092919050505061106f565b604051808215151515815260200191505060405180910390f35b6106ec61110f565b6040518082815260200191505060405180910390f35b61076e6004803603606081101561071857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611119565b604051808215151515815260200191505060405180910390f35b610790611260565b6040518082815260200191505060405180910390f35b610808600480360360408110156107bc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611266565b6040518082815260200191505060405180910390f35b6108606004803603602081101561083457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506112ed565b005b606060038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108fa5780601f106108cf576101008083540402835291602001916108fa565b820191906000526020600020905b8154815290600101906020018083116108dd57829003601f168201915b5050505050905090565b600061090e610e70565b610980576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b81600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060019050919050565b60006109de6109d7611373565b848461137b565b6001905092915050565b600080429050600a54610a066008548361157290919063ffffffff16565b11610a5c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c81526020018061207e602c913960400191505060405180910390fd5b6000610a67826115bc565b9050600560159054906101000a900460ff1660ff168114610ad3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180612118602a913960400191505060405180910390fd5b610adc826115f4565b5060019250505090565b6000600254905090565b6000610afd8484846116aa565b610bbe84610b09611373565b610bb98560405180606001604052806028815260200161205660289139600160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000610b6f611373565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546119609092919063ffffffff16565b61137b565b600190509392505050565b6000600560009054906101000a900460ff16905090565b6000610c89610bed611373565b84610c848560016000610bfe611373565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611a2090919063ffffffff16565b61137b565b6001905092915050565b6000600a54905090565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b610ced610e70565b610d5f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610eb3611373565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b600560159054906101000a900460ff1681565b606060048054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f7a5780601f10610f4f57610100808354040283529160200191610f7a565b820191906000526020600020905b815481529060010190602001808311610f5d57829003601f168201915b5050505050905090565b6000611047610f91611373565b84611042856040518060600160405280602581526020016121426025913960016000610fbb611373565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546119609092919063ffffffff16565b61137b565b6001905092915050565b600061106561105e611373565b84846116aa565b6001905092915050565b6000611079610e70565b6110eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b81600560156101000a81548160ff021916908360ff16021790555060019050919050565b6000600854905090565b6000611123610e70565b611195576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561121c57600080fd5b505af1158015611230573d6000803e3d6000fd5b505050506040513d602081101561124657600080fd5b810190808051906020019092919050505090509392505050565b60075481565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6112f5610e70565b611367576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b61137081611aa8565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611401576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806120f46024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611487576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602281526020018061200e6022913960400191505060405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b60006115b483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611960565b905092915050565b6000806115d56201518084611bee90919063ffffffff16565b90506115ec610e1082611c3890919063ffffffff16565b915050919050565b6000600754611615600954611607610ae6565b611a2090919063ffffffff16565b111561166c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806120cf6025913960400191505060405180910390fd5b61169a600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600954611c82565b8160088190555060019050919050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611730576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806120aa6025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156117b6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180611fc56023913960400191505060405180910390fd5b61182181604051806060016040528060268152602001612030602691396000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546119609092919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506118b4816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611a2090919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b6000838311158290611a0d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156119d25780820151818401526020810190506119b7565b50505050905090810190601f1680156119ff5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b600080828401905083811015611a9e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611b2e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180611fe86026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000611c3083836040518060400160405280601881526020017f536166654d6174683a206d6f64756c6f206279207a65726f0000000000000000815250611e3d565b905092915050565b6000611c7a83836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250611efe565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611d25576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b611d3a81600254611a2090919063ffffffff16565b600281905550611d91816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611a2090919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b6000808314158290611eea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611eaf578082015181840152602081019050611e94565b50505050905090810190601f168015611edc5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50828481611ef457fe5b0690509392505050565b60008083118290611faa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611f6f578082015181840152602081019050611f54565b50505050905090810190601f168015611f9c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506000838581611fb657fe5b04905080915050939250505056fe45524332303a207472616e7366657220746f20746865207a65726f20616464726573734f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e63654d54523a205468652072657175697265642074696d6520706572696f6420686173206e6f742070617373656445524332303a207472616e736665722066726f6d20746865207a65726f20616464726573734d54523a2043616e6e6f74206d696e74206d6f7265207468616e206d617820737570706c7945524332303a20617070726f76652066726f6d20746865207a65726f20616464726573734d54523a2043616e206f6e6c79206d696e7420647572696e6720746865206d696e74696e6720686f757245524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa264697066735822122089bd463009f3978d6d41bee3d5afa13b2ff3d457090f2c9aaaded7545df793f764736f6c63430006020033"

// DeployMatter deploys a new Ethereum contract, binding an instance of Matter to it.
func DeployMatter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Matter, error) {
	parsed, err := abi.JSON(strings.NewReader(MatterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MatterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Matter{MatterCaller: MatterCaller{contract: contract}, MatterTransactor: MatterTransactor{contract: contract}, MatterFilterer: MatterFilterer{contract: contract}}, nil
}

// Matter is an auto generated Go binding around an Ethereum contract.
type Matter struct {
	MatterCaller     // Read-only binding to the contract
	MatterTransactor // Write-only binding to the contract
	MatterFilterer   // Log filterer for contract events
}

// MatterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MatterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MatterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MatterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MatterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MatterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MatterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MatterSession struct {
	Contract     *Matter           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MatterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MatterCallerSession struct {
	Contract *MatterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MatterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MatterTransactorSession struct {
	Contract     *MatterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MatterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MatterRaw struct {
	Contract *Matter // Generic contract binding to access the raw methods on
}

// MatterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MatterCallerRaw struct {
	Contract *MatterCaller // Generic read-only contract binding to access the raw methods on
}

// MatterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MatterTransactorRaw struct {
	Contract *MatterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMatter creates a new instance of Matter, bound to a specific deployed contract.
func NewMatter(address common.Address, backend bind.ContractBackend) (*Matter, error) {
	contract, err := bindMatter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Matter{MatterCaller: MatterCaller{contract: contract}, MatterTransactor: MatterTransactor{contract: contract}, MatterFilterer: MatterFilterer{contract: contract}}, nil
}

// NewMatterCaller creates a new read-only instance of Matter, bound to a specific deployed contract.
func NewMatterCaller(address common.Address, caller bind.ContractCaller) (*MatterCaller, error) {
	contract, err := bindMatter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MatterCaller{contract: contract}, nil
}

// NewMatterTransactor creates a new write-only instance of Matter, bound to a specific deployed contract.
func NewMatterTransactor(address common.Address, transactor bind.ContractTransactor) (*MatterTransactor, error) {
	contract, err := bindMatter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MatterTransactor{contract: contract}, nil
}

// NewMatterFilterer creates a new log filterer instance of Matter, bound to a specific deployed contract.
func NewMatterFilterer(address common.Address, filterer bind.ContractFilterer) (*MatterFilterer, error) {
	contract, err := bindMatter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MatterFilterer{contract: contract}, nil
}

// bindMatter binds a generic wrapper to an already deployed contract.
func bindMatter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MatterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Matter *MatterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Matter.Contract.MatterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Matter *MatterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Matter.Contract.MatterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Matter *MatterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Matter.Contract.MatterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Matter *MatterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Matter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Matter *MatterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Matter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Matter *MatterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Matter.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_Matter *MatterCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Matter.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_Matter *MatterSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Matter.Contract.Allowance(&_Matter.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_Matter *MatterCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Matter.Contract.Allowance(&_Matter.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Matter *MatterCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Matter.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Matter *MatterSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Matter.Contract.BalanceOf(&_Matter.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Matter *MatterCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Matter.Contract.BalanceOf(&_Matter.CallOpts, account)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() constant returns(address)
func (_Matter *MatterCaller) Bank(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Matter.contract.Call(opts, out, "bank")
	return *ret0, err
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() constant returns(address)
func (_Matter *MatterSession) Bank() (common.Address, error) {
	return _Matter.Contract.Bank(&_Matter.CallOpts)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() constant returns(address)
func (_Matter *MatterCallerSession) Bank() (common.Address, error) {
	return _Matter.Contract.Bank(&_Matter.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Matter *MatterCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Matter.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Matter *MatterSession) Decimals() (uint8, error) {
	return _Matter.Contract.Decimals(&_Matter.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Matter *MatterCallerSession) Decimals() (uint8, error) {
	return _Matter.Contract.Decimals(&_Matter.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Matter *MatterCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Matter.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Matter *MatterSession) IsOwner() (bool, error) {
	return _Matter.Contract.IsOwner(&_Matter.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Matter *MatterCallerSession) IsOwner() (bool, error) {
	return _Matter.Contract.IsOwner(&_Matter.CallOpts)
}

// LastMinted is a free data retrieval call binding the contract method 0xbbc19ab0.
//
// Solidity: function lastMinted() constant returns(uint256)
func (_Matter *MatterCaller) LastMinted(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Matter.contract.Call(opts, out, "lastMinted")
	return *ret0, err
}

// LastMinted is a free data retrieval call binding the contract method 0xbbc19ab0.
//
// Solidity: function lastMinted() constant returns(uint256)
func (_Matter *MatterSession) LastMinted() (*big.Int, error) {
	return _Matter.Contract.LastMinted(&_Matter.CallOpts)
}

// LastMinted is a free data retrieval call binding the contract method 0xbbc19ab0.
//
// Solidity: function lastMinted() constant returns(uint256)
func (_Matter *MatterCallerSession) LastMinted() (*big.Int, error) {
	return _Matter.Contract.LastMinted(&_Matter.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() constant returns(uint256)
func (_Matter *MatterCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Matter.contract.Call(opts, out, "maxSupply")
	return *ret0, err
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() constant returns(uint256)
func (_Matter *MatterSession) MaxSupply() (*big.Int, error) {
	return _Matter.Contract.MaxSupply(&_Matter.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() constant returns(uint256)
func (_Matter *MatterCallerSession) MaxSupply() (*big.Int, error) {
	return _Matter.Contract.MaxSupply(&_Matter.CallOpts)
}

// MinMintInterval is a free data retrieval call binding the contract method 0x5620583c.
//
// Solidity: function minMintInterval() constant returns(uint256)
func (_Matter *MatterCaller) MinMintInterval(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Matter.contract.Call(opts, out, "minMintInterval")
	return *ret0, err
}

// MinMintInterval is a free data retrieval call binding the contract method 0x5620583c.
//
// Solidity: function minMintInterval() constant returns(uint256)
func (_Matter *MatterSession) MinMintInterval() (*big.Int, error) {
	return _Matter.Contract.MinMintInterval(&_Matter.CallOpts)
}

// MinMintInterval is a free data retrieval call binding the contract method 0x5620583c.
//
// Solidity: function minMintInterval() constant returns(uint256)
func (_Matter *MatterCallerSession) MinMintInterval() (*big.Int, error) {
	return _Matter.Contract.MinMintInterval(&_Matter.CallOpts)
}

// MintHour is a free data retrieval call binding the contract method 0x942cd61b.
//
// Solidity: function mintHour() constant returns(uint8)
func (_Matter *MatterCaller) MintHour(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Matter.contract.Call(opts, out, "mintHour")
	return *ret0, err
}

// MintHour is a free data retrieval call binding the contract method 0x942cd61b.
//
// Solidity: function mintHour() constant returns(uint8)
func (_Matter *MatterSession) MintHour() (uint8, error) {
	return _Matter.Contract.MintHour(&_Matter.CallOpts)
}

// MintHour is a free data retrieval call binding the contract method 0x942cd61b.
//
// Solidity: function mintHour() constant returns(uint8)
func (_Matter *MatterCallerSession) MintHour() (uint8, error) {
	return _Matter.Contract.MintHour(&_Matter.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Matter *MatterCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Matter.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Matter *MatterSession) Name() (string, error) {
	return _Matter.Contract.Name(&_Matter.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Matter *MatterCallerSession) Name() (string, error) {
	return _Matter.Contract.Name(&_Matter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Matter *MatterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Matter.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Matter *MatterSession) Owner() (common.Address, error) {
	return _Matter.Contract.Owner(&_Matter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Matter *MatterCallerSession) Owner() (common.Address, error) {
	return _Matter.Contract.Owner(&_Matter.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Matter *MatterCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Matter.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Matter *MatterSession) Symbol() (string, error) {
	return _Matter.Contract.Symbol(&_Matter.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Matter *MatterCallerSession) Symbol() (string, error) {
	return _Matter.Contract.Symbol(&_Matter.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Matter *MatterCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Matter.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Matter *MatterSession) TotalSupply() (*big.Int, error) {
	return _Matter.Contract.TotalSupply(&_Matter.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Matter *MatterCallerSession) TotalSupply() (*big.Int, error) {
	return _Matter.Contract.TotalSupply(&_Matter.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Matter *MatterTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Matter.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Matter *MatterSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Matter.Contract.Approve(&_Matter.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Matter *MatterTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Matter.Contract.Approve(&_Matter.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Matter *MatterTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Matter.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Matter *MatterSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Matter.Contract.DecreaseAllowance(&_Matter.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Matter *MatterTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Matter.Contract.DecreaseAllowance(&_Matter.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Matter *MatterTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Matter.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Matter *MatterSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Matter.Contract.IncreaseAllowance(&_Matter.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Matter *MatterTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Matter.Contract.IncreaseAllowance(&_Matter.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(bool)
func (_Matter *MatterTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Matter.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(bool)
func (_Matter *MatterSession) Mint() (*types.Transaction, error) {
	return _Matter.Contract.Mint(&_Matter.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(bool)
func (_Matter *MatterTransactorSession) Mint() (*types.Transaction, error) {
	return _Matter.Contract.Mint(&_Matter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Matter *MatterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Matter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Matter *MatterSession) RenounceOwnership() (*types.Transaction, error) {
	return _Matter.Contract.RenounceOwnership(&_Matter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Matter *MatterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Matter.Contract.RenounceOwnership(&_Matter.TransactOpts)
}

// SetBank is a paid mutator transaction binding the contract method 0x090d23b9.
//
// Solidity: function setBank(address newAddress) returns(bool success)
func (_Matter *MatterTransactor) SetBank(opts *bind.TransactOpts, newAddress common.Address) (*types.Transaction, error) {
	return _Matter.contract.Transact(opts, "setBank", newAddress)
}

// SetBank is a paid mutator transaction binding the contract method 0x090d23b9.
//
// Solidity: function setBank(address newAddress) returns(bool success)
func (_Matter *MatterSession) SetBank(newAddress common.Address) (*types.Transaction, error) {
	return _Matter.Contract.SetBank(&_Matter.TransactOpts, newAddress)
}

// SetBank is a paid mutator transaction binding the contract method 0x090d23b9.
//
// Solidity: function setBank(address newAddress) returns(bool success)
func (_Matter *MatterTransactorSession) SetBank(newAddress common.Address) (*types.Transaction, error) {
	return _Matter.Contract.SetBank(&_Matter.TransactOpts, newAddress)
}

// SetMintHour is a paid mutator transaction binding the contract method 0xb0601abe.
//
// Solidity: function setMintHour(uint8 newMintHour) returns(bool success)
func (_Matter *MatterTransactor) SetMintHour(opts *bind.TransactOpts, newMintHour uint8) (*types.Transaction, error) {
	return _Matter.contract.Transact(opts, "setMintHour", newMintHour)
}

// SetMintHour is a paid mutator transaction binding the contract method 0xb0601abe.
//
// Solidity: function setMintHour(uint8 newMintHour) returns(bool success)
func (_Matter *MatterSession) SetMintHour(newMintHour uint8) (*types.Transaction, error) {
	return _Matter.Contract.SetMintHour(&_Matter.TransactOpts, newMintHour)
}

// SetMintHour is a paid mutator transaction binding the contract method 0xb0601abe.
//
// Solidity: function setMintHour(uint8 newMintHour) returns(bool success)
func (_Matter *MatterTransactorSession) SetMintHour(newMintHour uint8) (*types.Transaction, error) {
	return _Matter.Contract.SetMintHour(&_Matter.TransactOpts, newMintHour)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Matter *MatterTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Matter.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Matter *MatterSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Matter.Contract.Transfer(&_Matter.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Matter *MatterTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Matter.Contract.Transfer(&_Matter.TransactOpts, recipient, amount)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xd493b9ac.
//
// Solidity: function transferAnyERC20Token(address tokenAddress, address toAddress, uint256 tokens) returns(bool success)
func (_Matter *MatterTransactor) TransferAnyERC20Token(opts *bind.TransactOpts, tokenAddress common.Address, toAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Matter.contract.Transact(opts, "transferAnyERC20Token", tokenAddress, toAddress, tokens)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xd493b9ac.
//
// Solidity: function transferAnyERC20Token(address tokenAddress, address toAddress, uint256 tokens) returns(bool success)
func (_Matter *MatterSession) TransferAnyERC20Token(tokenAddress common.Address, toAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Matter.Contract.TransferAnyERC20Token(&_Matter.TransactOpts, tokenAddress, toAddress, tokens)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xd493b9ac.
//
// Solidity: function transferAnyERC20Token(address tokenAddress, address toAddress, uint256 tokens) returns(bool success)
func (_Matter *MatterTransactorSession) TransferAnyERC20Token(tokenAddress common.Address, toAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Matter.Contract.TransferAnyERC20Token(&_Matter.TransactOpts, tokenAddress, toAddress, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Matter *MatterTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Matter.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Matter *MatterSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Matter.Contract.TransferFrom(&_Matter.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Matter *MatterTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Matter.Contract.TransferFrom(&_Matter.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Matter *MatterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Matter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Matter *MatterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Matter.Contract.TransferOwnership(&_Matter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Matter *MatterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Matter.Contract.TransferOwnership(&_Matter.TransactOpts, newOwner)
}

// MatterApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Matter contract.
type MatterApprovalIterator struct {
	Event *MatterApproval // Event containing the contract specifics and raw log

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
func (it *MatterApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MatterApproval)
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
		it.Event = new(MatterApproval)
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
func (it *MatterApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MatterApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MatterApproval represents a Approval event raised by the Matter contract.
type MatterApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Matter *MatterFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MatterApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Matter.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MatterApprovalIterator{contract: _Matter.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Matter *MatterFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MatterApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Matter.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MatterApproval)
				if err := _Matter.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Matter *MatterFilterer) ParseApproval(log types.Log) (*MatterApproval, error) {
	event := new(MatterApproval)
	if err := _Matter.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MatterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Matter contract.
type MatterOwnershipTransferredIterator struct {
	Event *MatterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MatterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MatterOwnershipTransferred)
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
		it.Event = new(MatterOwnershipTransferred)
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
func (it *MatterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MatterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MatterOwnershipTransferred represents a OwnershipTransferred event raised by the Matter contract.
type MatterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Matter *MatterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MatterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Matter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MatterOwnershipTransferredIterator{contract: _Matter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Matter *MatterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MatterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Matter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MatterOwnershipTransferred)
				if err := _Matter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Matter *MatterFilterer) ParseOwnershipTransferred(log types.Log) (*MatterOwnershipTransferred, error) {
	event := new(MatterOwnershipTransferred)
	if err := _Matter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MatterTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Matter contract.
type MatterTransferIterator struct {
	Event *MatterTransfer // Event containing the contract specifics and raw log

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
func (it *MatterTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MatterTransfer)
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
		it.Event = new(MatterTransfer)
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
func (it *MatterTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MatterTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MatterTransfer represents a Transfer event raised by the Matter contract.
type MatterTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Matter *MatterFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MatterTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Matter.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MatterTransferIterator{contract: _Matter.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Matter *MatterFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MatterTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Matter.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MatterTransfer)
				if err := _Matter.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Matter *MatterFilterer) ParseTransfer(log types.Log) (*MatterTransfer, error) {
	event := new(MatterTransfer)
	if err := _Matter.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
