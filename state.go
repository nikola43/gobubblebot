package main

import (
	"sync"

	"github.com/nikola43/web3golanghelper/web3helper"
)

var wg sync.WaitGroup
var state = make(map[int64]map[string]interface{})
var goroutines = make(map[int]chan bool)
var goroutineId = 0


var bscWeb3 = web3helper.NewWeb3GolangHelper("https://mainnet.infura.io/v3/14c8be7030d142fb9704771c6acfddb5", "wss://mainnet.infura.io/ws/v3/14c8be7030d142fb9704771c6acfddb5")
//var bscWeb3 = web3helper.NewWeb3GolangHelper("https://goerli.infura.io/v3/14c8be7030d142fb9704771c6acfddb5", "wss://goerli.infura.io/ws/v3/14c8be7030d142fb9704771c6acfddb5")
