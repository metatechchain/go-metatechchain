package driver

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// GetContractBin is NodeDriver contract genesis implementation bin code
// Has to be compiled with flag bin-runtime
// Built from metatechchain-sfc c1d33c81f74abf82c0e22807f16e609578e10ad8, solc 0.5.17+commit.d19bba13.Emscripten.clang, optimize-runs 10000
func GetContractBin() []byte {
	return hexutil.MustDecode("0x608060405234801561001057600080fd5b5061239d806100206000396000f3fe6080604052600436106100f1576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806307690b2a146100f65780630aeeca001461016757806318f628d4146101a25780631e702f8314610244578063242a6e3f14610289578063267ab4461461031957806339e503ab14610354578063485cc955146103b95780634feb92f31461042a57806379bead381461050d578063a4066fbe14610568578063b9cc6b1c146105ad578063d6a0c7af14610633578063da7fc24f146106a4578063e08d7e66146106f5578063e30443bc1461077b578063ebdf104c146107d6575b600080fd5b34801561010257600080fd5b506101656004803603604081101561011957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061095b565b005b34801561017357600080fd5b506101a06004803603602081101561018a57600080fd5b8101908080359060200190929190505050610b2d565b005b3480156101ae57600080fd5b5061024260048036036101208110156101c657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190505050610c2c565b005b34801561025057600080fd5b506102876004803603604081101561026757600080fd5b810190808035906020019092919080359060200190929190505050610df0565b005b34801561029557600080fd5b50610317600480360360408110156102ac57600080fd5b8101908080359060200190929190803590602001906401000000008111156102d357600080fd5b8201836020820111156102e557600080fd5b8035906020019184600183028401116401000000008311171561030757600080fd5b9091929391929390505050610f49565b005b34801561032557600080fd5b506103526004803603602081101561033c57600080fd5b8101908080359060200190929190505050611077565b005b34801561036057600080fd5b506103b76004803603606081101561037757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190505050611176565b005b3480156103c557600080fd5b50610428600480360360408110156103dc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611325565b005b34801561043657600080fd5b5061050b600480360361010081101561044e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561049557600080fd5b8201836020820111156104a757600080fd5b803590602001918460018302840111640100000000831117156104c957600080fd5b90919293919293908035906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190505050611530565b005b34801561051957600080fd5b506105666004803603604081101561053057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611718565b005b34801561057457600080fd5b506105ab6004803603604081101561058b57600080fd5b8101908080359060200190929190803590602001909291905050506118be565b005b3480156105b957600080fd5b50610631600480360360208110156105d057600080fd5b81019080803590602001906401000000008111156105ed57600080fd5b8201836020820111156105ff57600080fd5b8035906020019184600183028401116401000000008311171561062157600080fd5b90919293919293905050506119bf565b005b34801561063f57600080fd5b506106a26004803603604081101561065657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611aeb565b005b3480156106b057600080fd5b506106f3600480360360208110156106c757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611cbd565b005b34801561070157600080fd5b506107796004803603602081101561071857600080fd5b810190808035906020019064010000000081111561073557600080fd5b82018360208201111561074757600080fd5b8035906020019184602083028401116401000000008311171561076957600080fd5b9091929391929390505050611e09565b005b34801561078757600080fd5b506107d46004803603604081101561079e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611f89565b005b3480156107e257600080fd5b50610959600480360360808110156107f957600080fd5b810190808035906020019064010000000081111561081657600080fd5b82018360208201111561082857600080fd5b8035906020019184602083028401116401000000008311171561084a57600080fd5b90919293919293908035906020019064010000000081111561086b57600080fd5b82018360208201111561087d57600080fd5b8035906020019184602083028401116401000000008311171561089f57600080fd5b9091929391929390803590602001906401000000008111156108c057600080fd5b8201836020820111156108d257600080fd5b803590602001918460208302840111640100000000831117156108f457600080fd5b90919293919293908035906020019064010000000081111561091557600080fd5b82018360208201111561092757600080fd5b8035906020019184602083028401116401000000008311171561094957600080fd5b909192939192939050505061212f565b005b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610a20576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166307690b2a83836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b158015610b1157600080fd5b505af1158015610b25573d6000803e3d6000fd5b505050505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610bf2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b7f0151256d62457b809bbc891b1f81c6dd0b9987552c70ce915b519750cd434dd1816040518082815260200191505060405180910390a150565b600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610cd0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f6e6f742063616c6c61626c65000000000000000000000000000000000000000081525060200191505060405180910390fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318f628d48a8a8a8a8a8a8a8a8a6040518a63ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018881526020018781526020018681526020018581526020018481526020018381526020018281526020019950505050505050505050600060405180830381600087803b158015610dcd57600080fd5b505af1158015610de1573d6000803e3d6000fd5b50505050505050505050505050565b600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610e94576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f6e6f742063616c6c61626c65000000000000000000000000000000000000000081525060200191505060405180910390fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631e702f8383836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083815260200182815260200192505050600060405180830381600087803b158015610f2d57600080fd5b505af1158015610f41573d6000803e3d6000fd5b505050505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561100e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b827f0f0ef1ab97439def0a9d2c6d9dc166207f1b13b99e62b442b2993d6153c63a6e838360405180806020018281038252848482818152602001925080828437600081840152601f19601f820116905080830192505050935050505060405180910390a2505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561113c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b7f2ccdfd47cf0c1f1069d949f1789bb79b2f12821f021634fc835af1de66ea2feb816040518082815260200191505060405180910390a150565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561123b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166339e503ab8484846040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020019350505050600060405180830381600087803b15801561130857600080fd5b505af115801561131c573d6000803e3d6000fd5b50505050505050565b600060019054906101000a900460ff1680611344575061134361235a565b5b8061135b57506000809054906101000a900460ff16155b15156113f5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e8152602001807f436f6e747261637420696e7374616e63652068617320616c726561647920626581526020017f656e20696e697469616c697a656400000000000000000000000000000000000081525060400191505060405180910390fd5b60008060019054906101000a900460ff161590508015611445576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b82603460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff167f64ee8f7bfc37fc205d7194ee3d64947ab7b57e663cd0d1abd3ef24503583069360405160405180910390a281603560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550801561152b5760008060016101000a81548160ff0219169083151502179055505b505050565b600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156115d4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f6e6f742063616c6c61626c65000000000000000000000000000000000000000081525060200191505060405180910390fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634feb92f38a8a8a8a8a8a8a8a8a6040518a63ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001898152602001806020018781526020018681526020018581526020018481526020018381526020018281038252898982818152602001925080828437600081840152601f19601f8201169050808301925050509a5050505050505050505050600060405180830381600087803b1580156116f557600080fd5b505af1158015611709573d6000803e3d6000fd5b50505050505050505050505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156117dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166379bead3883836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1580156118a257600080fd5b505af11580156118b6573d6000803e3d6000fd5b505050505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611983576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b817fb975807576e3b1461be7de07ebf7d20e4790ed802d7a0c4fdd0a1a13df72a935826040518082815260200191505060405180910390a25050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611a84576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b7f47d10eed096a44e3d0abc586c7e3a5d6cb5358cc90e7d437cd0627f7e765fb99828260405180806020018281038252848482818152602001925080828437600081840152601f19601f820116905080830192505050935050505060405180910390a15050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611bb0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d6a0c7af83836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b158015611ca157600080fd5b505af1158015611cb5573d6000803e3d6000fd5b505050505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611d82576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff167f64ee8f7bfc37fc205d7194ee3d64947ab7b57e663cd0d1abd3ef24503583069360405160405180910390a280603460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611ead576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f6e6f742063616c6c61626c65000000000000000000000000000000000000000081525060200191505060405180910390fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e08d7e6683836040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018281038252848482818152602001925060200280828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b158015611f6d57600080fd5b505af1158015611f81573d6000803e3d6000fd5b505050505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561204e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e30443bc83836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561211357600080fd5b505af1158015612127573d6000803e3d6000fd5b505050505050565b600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156121d3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f6e6f742063616c6c61626c65000000000000000000000000000000000000000081525060200191505060405180910390fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ebdf104c89898989898989896040518963ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808060200180602001806020018060200185810385528d8d82818152602001925060200280828437600081840152601f19601f82011690508083019250505085810384528b8b82818152602001925060200280828437600081840152601f19601f8201169050808301925050508581038352898982818152602001925060200280828437600081840152601f19601f8201169050808301925050508581038252878782818152602001925060200280828437600081840152601f19601f8201169050808301925050509c50505050505050505050505050600060405180830381600087803b15801561233857600080fd5b505af115801561234c573d6000803e3d6000fd5b505050505050505050505050565b6000803090506000813b905060008114925050509056fea165627a7a72305820e66389cb7a564eb545b7c1e7e4f3b44d0eecd6a3e174b047759b9c8115efb7d50029")
}

// ContractAddress is the NodeDriver contract address
var ContractAddress = common.HexToAddress("0xd100a01e00000000000000000000000000000000")
