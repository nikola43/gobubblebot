package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/fatih/color"
	"github.com/mymmrac/telego"
	Pulsedoge "github.com/nikola43/gobubblebot/Pulsedoge"
)

func HandleAction(chatID int64, action string, bot *telego.Bot) error {
	_, err := SendMessage(chatID, INPUT_CAPTIONS[action], nil, bot, false)
	if err != nil {
		return err
	}
	state[chatID]["InputMode"] = action
	return nil
}

func ActionStart(chatID int64, bot *telego.Bot) error {
	fmt.Println("start")

	/*
		_, err := SendMessage(chatID, "welcome", nil, bot)
		if err != nil {
			return err
		}
	*/

	HandleActionWithKeyboard(chatID, ShowMenu, "Configure your bot", bot)

	return nil
}

func HandleActionWithKeyboard(chatID int64, action string, message string, bot *telego.Bot) error {
	keyboard := KEYBOARDS[action]
	_, err := SendMessage(chatID, message, keyboard, bot, false)
	if err != nil {
		return err
	}
	//state[chatID]["InputMode"] = action
	return nil
}

/*
func ActionBack(chatID int64, bot *telego.Bot) error {
	fmt.Println("start")

	_, err := SendMessage(chatID, "welcome", nil, bot)
	if err != nil {
		return err
	}

	txt := "Welcome to the wallet bot. Please choose an option below."

	_, err = SendMessage(chatID, txt, WALLET_KEYBOARD, bot)
	if err != nil {
		return err
	}
	return nil
}

*/

func ActionStartBot(chatID int64, bot *telego.Bot) error {
	if state[chatID]["gID"] != nil {
		SendMessage(chatID, "Bot already started", nil, bot, false)
		return nil
	}

	tokenConfig := state[chatID]["TokenConfig"].(TokenConfig)

	/*
		if tokenConfig.Address == "" {
			SendMessage(chatID, "Token Address is required!", nil, bot)
			return nil
		}

		if tokenConfig.Group == "" {
			SendMessage(chatID, "Token Address is required!", nil, bot)
			return nil
		}
	*/

	fmt.Println(tokenConfig)

	done := make(chan bool)
	goroutineId++
	goroutines[goroutineId] = done
	wg.Add(1)
	SendMessage(chatID, "Bot started", nil, bot, false)
	state[chatID]["gID"] = goroutineId

	contractAbi, _ := abi.JSON(strings.NewReader(string(Pulsedoge.PulsedogeABI)))
	logs := make(chan types.Log)
	tokenAddress := tokenConfig.Address
	pair := tokenConfig.Pair
	sub := BuildContractEventSubscription(bscWeb3, tokenAddress, logs)
	fmt.Println(color.YellowString("  ----------------- Blockchain Events -----------------"))
	fmt.Println(color.CyanString("\tListen token address: "), color.GreenString(tokenAddress))

	go func() {
		defer wg.Done()
		defer delete(goroutines, goroutineId)

		isBuy := false
		ethAmount := big.NewInt(0)
		tokenAmount := big.NewInt(0)

		for {
			select {
			case <-done:
				fmt.Printf("done Goroutine %d exiting\n", goroutineId)
				return
			case <-stopChan:
				fmt.Printf("stopChan Goroutine %d exiting\n", goroutineId)
				return // Exit the goroutine when stopChan is closed
			case err := <-sub.Err():
				fmt.Println("error socket", err)
				//bscWeb3 = web3helper.NewWeb3GolangHelper(BSC_RPC_URL, BSC_WS_URL)
				//sub = BuildContractEventSubscription(bscWeb3, BSC_TOKEN_ADDRESS, logs)

			case vLog := <-logs:
				event, err := contractAbi.EventByID(vLog.Topics[0])
				if err != nil {
					panic(err)
				}
				fmt.Printf("Event Name: %s\n", event.Name)

				tx, pending, err := bscWeb3.HttpClient().TransactionByHash(context.Background(), vLog.TxHash)
				_ = pending
				if err != nil {
					log.Fatal(err)
				}
				userAddress, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
				if err != nil {
					log.Fatal(err)
				}

				//recipient, err := bscWeb3.HttpClient().TransactionReceipt(context.Background(), vLog.TxHash)

				//t := tx.To().Hex()
				//fmt.Println("t")
				//fmt.Println(t)

				// check if event is Transfer
				if event.Name == "Transfer" {

					fmt.Println("userAddress", userAddress)

					// parse event data
					from, to, value, err := ExtractEventLogData(vLog, contractAbi, event.Name)
					if err != nil {
						panic(err)
					}

					fmt.Println("vLog.TxHash: " + vLog.TxHash.Hex())
					fmt.Println("From: " + from.Hex())
					fmt.Println("To: " + to.Hex())
					fmt.Printf("Value: %v\n", value)
					fmt.Println()

					// Buy
					//ethAmount = tx.Value()
					//if to.Hex() == common.HexToAddress(pair).Hex() {
					//if to.Hex() == common.HexToAddress(pair).Hex() {
					//	ethAmount = tx.Value()
					//	fmt.Println("ethAmount")
					//	fmt.Println(ethAmount)
					//}

					// Buy token
					//if from.Hex() == common.HexToAddress(pair).Hex() &&  token_address == user_config['token_address'] and receiver_address == from_address {

					//fmt.Println("from.Hex()", from.Hex())
					//fmt.Println("common.HexToAddress(pair).Hex()", common.HexToAddress(pair).Hex())
					//fmt.Println("to.Hex()", to.Hex())
					//fmt.Println("userAddress.Hex()", userAddress.Hex())

					if from.Hex() == common.HexToAddress(pair).Hex() &&
						to.Hex() == userAddress.Hex() { //&& tokenAddress.Hex() == tx.To().Hex() {
						tokenAmount = value
						ethAmount = tx.Value()
					}

					isBuy = from.Hex() == common.HexToAddress(pair).Hex()
				}

				// fmt.Println("BUY", isBuy)
				// fmt.Println("ethAmount")
				// fmt.Println(ethAmount)

				//fmt.Println("tokenAmount")
				//fmt.Println(tokenAmount)

				if isBuy && ethAmount.Cmp(big.NewInt(0)) == 1 && tokenAmount.Cmp(big.NewInt(0)) == 1 {
					fmt.Println("BUY")
					fmt.Println("ethAmount")
					fmt.Println(ethAmount)

					fmt.Println("tokenAmount")
					fmt.Println(tokenAmount)

					//msg := "ethAmount " + ToDecimal(ethAmount, 18).String() + "\n"
					//msg += "tokenAmount " + ToDecimal(tokenAmount, 18).String() + "\n"

					tokenInfo, _ := GetTokenInfoDex(tokenConfig.Address)
					println("*tokenInfo.Fdv", *tokenInfo.Fdv)
					mc := fmt.Sprintf("%f", *tokenInfo.Fdv)
					volume24 := fmt.Sprintf("%f", *tokenInfo.Volume.H24)
					tokenHoldersTaxes, _ := GetTokenHoldersAndTaxes(tokenConfig.Pair)

					msg := BuildBubbleMessage(tokenConfig, userAddress.Hex(), tx.Hash().Hex(), ethAmount, tokenAmount, mc, volume24, *tokenHoldersTaxes.Holders, tokenHoldersTaxes.Buy, tokenHoldersTaxes.Sell)
					msg = escapeMarkdown(msg)
					fmt.Println(msg)

					_, err := SendMessage(chatID, msg, nil, bot, true)
					fmt.Println(err)
					fmt.Println(err)

					isBuy = false
					ethAmount = big.NewInt(0)
					tokenAmount = big.NewInt(0)
				}

			default:

				// Your goroutine's work here
				//fmt.Printf("Goroutine %d is running\n", goroutineId)
				//time.Sleep(time.Second)

			}
		}
	}()

	return nil
}

func escapeMarkdown(message string) string {
	message = strings.ReplaceAll(message, "_", "\\_")
	message = strings.ReplaceAll(message, "`", "\\`")
	message = strings.ReplaceAll(message, ".", "\\.")
	message = strings.ReplaceAll(message, "%", "\\%")
	message = strings.ReplaceAll(message, "-", "\\-")
	message = strings.ReplaceAll(message, "+", "\\+")
	message = strings.ReplaceAll(message, "#", "\\#")
	message = strings.ReplaceAll(message, "=", "\\=")
	message = strings.ReplaceAll(message, "|", "\\|")
	message = strings.ReplaceAll(message, "!", "\\!")
	return message
}

func ActionStopBot(chatID int64, bot *telego.Bot) error {

	if state[chatID]["gID"] == nil {
		SendMessage(chatID, "Bot not started", nil, bot, false)
		return nil
	}

	id := state[chatID]["gID"].(int)
	if done, exists := goroutines[id]; exists {
		close(done)
	}
	wg.Wait()
	state[chatID]["gID"] = nil
	goroutineId--
	SendMessage(chatID, "Bot stopped", nil, bot, false)
	return nil
}
