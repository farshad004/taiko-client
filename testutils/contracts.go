package testutils

import "github.com/ethereum/go-ethereum/common"

var (
	TaikoL2Address      = common.HexToAddress("0x1000777700000000000000000000000000000001")
	OracleProverAddress = common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	TreasuryAddress     = common.HexToAddress("0xdf09A0afD09a63fb04ab3573922437e1e637dE8b")
)

// variables need to be initialized
var (
	TaikoL1Address       common.Address
	TaikoL1TokenAddress  common.Address
	TaikoL1SignalService common.Address
)