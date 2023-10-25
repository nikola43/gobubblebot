package main

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nikola43/web3golanghelper/web3helper"
)

func BuildContractEventSubscription(web3GolangHelper *web3helper.Web3GolangHelper, contractAddress string, logs chan types.Log) (ethereum.Subscription, error) {
	sub, err := web3GolangHelper.WebSocketClient().SubscribeFilterLogs(
		context.Background(),
		ethereum.FilterQuery{
			Addresses: []common.Address{common.HexToAddress(contractAddress)},
		}, logs)
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func ExtractEventLogData(vLog types.Log, contractAbi abi.ABI, eventName string) (*TransferEventData, error) {
	outputDataMap := make(map[string]interface{})
	err := contractAbi.UnpackIntoMap(outputDataMap, eventName, vLog.Data)
	if err != nil {
		return nil, err
	}

	return &TransferEventData{
		From:   common.HexToAddress(vLog.Topics[1].Hex()),
		To:     common.HexToAddress(vLog.Topics[2].Hex()),
		Amount: outputDataMap["value"].(*big.Int),
	}, nil
}
