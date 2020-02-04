This project takes Solidity Smart Contracts and compiles them into GoLang. The main.go script provides some example calls that can be made on the Smart Contracts after they are compiled. The steps below list out how I compiled the matter.sol contract into a go file. From root:

1.) Compile Soidity contract into ABI and Bytecode and save them into the build/<whatever> folder:

$ solc --abi contracts/matter.sol -o build/matter
$ solc --bin contracts/matter.sol -o build/matter/

2.) Compile the ABI and Bytecode into a Go file:
$ abigen --bin=./build/matter/Matter.bin --abi=./build/matter/Matter.abi --pkg=store --out=./compiled/matter/MatterContract.go
