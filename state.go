package main

import (
	"sync"

	"github.com/nikola43/web3golanghelper/web3helper"
)

var wg sync.WaitGroup
var state = make(map[int64]map[string]interface{})
var goroutines = make(map[int]chan bool)
var goroutineId = 0


var bscWeb3 = web3helper.NewWeb3GolangHelper("https://mainnet.infura.io/v3/fe0b9cf93b1047bda0a6e7915f041380", "wss://mainnet.infura.io/ws/v3/fe0b9cf93b1047bda0a6e7915f041380")
//var bscWeb3 = web3helper.NewWeb3GolangHelper("https://goerli.infura.io/v3/fe0b9cf93b1047bda0a6e7915f041380", "wss://goerli.infura.io/ws/v3/fe0b9cf93b1047bda0a6e7915f041380")
