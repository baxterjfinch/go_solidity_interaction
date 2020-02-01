"# go_solidity_interaction"


// Compile contract to ABI
solc --abi contract.sol -o build


// convert abi to Go
abigen --abi=./build/contract.abi --pkg=contract --out=Contract.go


// convert contract to bytecode for deployment
solc --bin Contract.sol -o build

// compile the go contract
abigen --bin=./test/ERC1155Mintable.bin --abi=./test/ERC1155Mintable.abi --pkg=store --out=TestContract.go
