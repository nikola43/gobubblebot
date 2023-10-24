package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/fatih/color"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	Pulsedoge "github.com/nikola43/gobubblebot/Pulsedoge"
	"github.com/shopspring/decimal"
)

var stopChan = make(chan struct{})

func HandleInput(inputMode string, update telego.Update, bot *telego.Bot) error {
	chatID := update.Message.Chat.ID
	msgText := update.Message.Text

	fmt.Println(inputMode)

	switch inputMode {

	// --------------------- SetToken ---------------------
	case SetToken:
		if len(msgText) != 42 {
			txt := "Invalid token address, please try again."
			SendMessage(chatID, txt, nil, bot)
			return nil
		}
		tokenConfig := TokenConfig{
			Address: msgText,
		}
		state[chatID]["TokenConfig"] = tokenConfig

	// --------------------- SetGroup ---------------------
	case SetGroup:
		/*
			if len(msgText) != 42 {
				txt := "Invalid token address, please try again."
				SendMessage(chatID, txt, nil, bot)
				return nil
			}
		*/
		tokenConfig := state[chatID]["TokenConfig"].(TokenConfig)
		tokenConfig.Group = msgText
		state[chatID]["TokenConfig"] = tokenConfig
	// --------------------- SetTelegram ---------------------
	case SetTelegram:
		/*
			if len(msgText) != 42 {
				txt := "Invalid token address, please try again."
				SendMessage(chatID, txt, nil, bot)
				return nil
			}
		*/
		tokenConfig := state[chatID]["TokenConfig"].(TokenConfig)
		tokenConfig.Telegram = msgText
		state[chatID]["TokenConfig"] = tokenConfig
	// --------------------- SetWebsite ---------------------
	case SetWebsite:
		/*
			if len(msgText) != 42 {
				txt := "Invalid token address, please try again."
				SendMessage(chatID, txt, nil, bot)
				return nil
			}
		*/
		tokenConfig := state[chatID]["TokenConfig"].(TokenConfig)
		tokenConfig.Website = msgText
		state[chatID]["TokenConfig"] = tokenConfig
	// --------------------- SetTwitter ---------------------
	case SetTwitter:
		/*
			if len(msgText) != 42 {
				txt := "Invalid token address, please try again."
				SendMessage(chatID, txt, nil, bot)
				return nil
			}
		*/
		tokenConfig := state[chatID]["TokenConfig"].(TokenConfig)
		tokenConfig.Twitter = msgText
		state[chatID]["TokenConfig"] = tokenConfig
	}

	err := DeleteMessage(chatID, update.Message.MessageID, bot)
	if err != nil {
		return err
	}

	msg := BuildConfigMessage(state[chatID]["TokenConfig"])
	HandleActionWithKeyboard(chatID, ShowMenu, msg, bot)
	return nil
}

func BuildConfigMessage(config interface{}) string {
	tokenConfig := reflect.ValueOf(config)
	msg := "⚙️ Bot config ⚙️" + "\n\n"
	msg += "📄 Address: " + tokenConfig.FieldByName("Address").String() + "\n\n"
	msg += "🚻 Group: " + tokenConfig.FieldByName("Group").String() + "\n\n"
	msg += " ➤ Telegram: " + tokenConfig.FieldByName("Telegram").String() + "\n\n"
	msg += "🌐 Website: " + tokenConfig.FieldByName("Website").String() + "\n\n"
	msg += " 𝕏 Twitter: " + tokenConfig.FieldByName("Twitter").String() + "\n\n"

	return msg
}

func HandleButtonCallback(callback *telego.CallbackQuery, bot *telego.Bot) error {
	fmt.Println("Received callback with data:", callback.Data)
	fmt.Println("Received callback with message:", callback.Message.Text)
	fmt.Println("Received callback with chat id:", callback.Message.Chat.ID)

	// get chat id
	chatID := callback.Message.Chat.ID

	/*
		if state[chatID]["BotMessageID"] != nil {
			err := DeleteMessage(chatID, state[chatID]["BotMessageID"].(int), bot)
			if err != nil {
				return err
			}
		}
	*/

	// check fi callback data contais "back" word
	isBack := strings.Contains(callback.Data, Back)
	if isBack {
		// check if callback data is back/..., if so, handle back
		back := strings.Split(callback.Data, "/")[0]
		path := strings.Split(callback.Data, "/")[1]
		fmt.Println("back:", back)
		fmt.Println("path:", path)
		HandleBack(chatID, path, bot)
	}

	switch callback.Data {
	case ShowMenu:
		msg := BuildConfigMessage(state[chatID]["TokenConfig"])
		HandleActionWithKeyboard(chatID, ShowMenu, msg, bot)
	case Disconnect:
		HandleActionWithKeyboard(chatID, Start, "message", bot)
	case SetToken:
		HandleAction(chatID, SetToken, bot)
	case SetGroup:
		HandleAction(chatID, SetGroup, bot)
	case SetTelegram:
		HandleAction(chatID, SetTelegram, bot)
	case SetWebsite:
		HandleAction(chatID, SetWebsite, bot)
	case SetTwitter:
		HandleAction(chatID, SetTwitter, bot)

	case StartBot:
		if state[chatID]["gID"] != nil {
			SendMessage(chatID, "Bot already started", nil, bot)
			return nil
		}

		done := make(chan bool)
		goroutineId++
		goroutines[goroutineId] = done
		wg.Add(1)
		SendMessage(chatID, "Buy init", nil, bot)
		state[chatID]["gID"] = goroutineId

		contractAbi, _ := abi.JSON(strings.NewReader(string(Pulsedoge.PulsedogeABI)))
		logs := make(chan types.Log)
		tokenAddress := common.HexToAddress("0x0080428794a79a40ae03cf6e6c1d56bd5467a4a2")
		pair := "0x7cE52DC08B49e0Bc11252cD267d2614dfC616D9f"
		sub := BuildContractEventSubscription(bscWeb3, tokenAddress.Hex(), logs)
		fmt.Println(color.YellowString("  ----------------- Blockchain Events -----------------"))
		fmt.Println(color.CyanString("\tListen token address: "), color.GreenString(tokenAddress.Hex()))

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
					fmt.Println("userAddress", userAddress)
					//recipient, err := bscWeb3.HttpClient().TransactionReceipt(context.Background(), vLog.TxHash)

					//t := tx.To().Hex()
					//fmt.Println("t")
					//fmt.Println(t)

					// check if event is Transfer
					if event.Name == "Transfer" {

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
						ethAmount = tx.Value()
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

						msg := "ethAmount " + ToDecimal(ethAmount, 18).String() + "\n"
						msg += "tokenAmount " + ToDecimal(tokenAmount, 18).String() + "\n"
						SendMessage(chatID, msg, nil, bot)

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

	case StopBot:
		if state[chatID]["gID"] == nil {
			SendMessage(chatID, "Bot not started", nil, bot)
			return nil
		}

		id := state[chatID]["gID"].(int)
		if done, exists := goroutines[id]; exists {
			close(done)
		}
		wg.Wait()
		state[chatID]["gID"] = nil
		goroutineId--
		SendMessage(chatID, "Bot stopped", nil, bot)
	}

	return nil
}

func BuildKeyboard(btnLabels []string, callbacks []string) telego.ReplyMarkup {
	btns := make([][]telego.InlineKeyboardButton, 0)
	for i := 0; i < len(btnLabels); i++ {
		btns = append(btns, []telego.InlineKeyboardButton{
			tu.InlineKeyboardButton(btnLabels[i]).WithCallbackData(callbacks[i]),
		})
	}

	keyboard := tu.InlineKeyboard(
		btns...,
	)
	return keyboard
}

func HandleMessage(update telego.Update, bot *telego.Bot) error {

	chatID := update.Message.Chat.ID
	msgText := update.Message.Text

	if state[chatID] == nil {
		state[chatID] = make(map[string]interface{})
		state[chatID]["BotMessageID"] = 0
		state[chatID]["InputMode"] = ""
		state[chatID]["Account"] = nil
	}
	fmt.Println(state[chatID])

	botMessageID := state[chatID]["BotMessageID"].(int)
	inputMode := state[chatID]["InputMode"].(string)

	switch msgText {
	case "/start":
		_, err := SendMessage(chatID, "welcome", nil, bot)
		if err != nil {
			return err
		}
		ActionStart(chatID, bot)
	}

	if inputMode != "" {
		HandleInput(inputMode, update, bot)
		DeleteMessage(chatID, botMessageID, bot)
	}

	return nil
}

func showConfig() {

}

func showPage(page string) {

}

func SendMessage(chatID int64, msg string, replyMarkup telego.ReplyMarkup, bot *telego.Bot) (*telego.Message, error) {
	var message *telego.SendMessageParams

	if replyMarkup != nil {
		message = tu.Message(
			tu.ID(chatID),
			msg,
		).WithReplyMarkup(replyMarkup)
	} else {
		message = tu.Message(
			tu.ID(chatID),
			msg,
		)
	}
	// Sending message
	res, err := bot.SendMessage(message)
	if err != nil {
		return nil, err
	}

	state[chatID]["BotMessageID"] = res.MessageID
	return res, nil
}

func DeleteMessage(chatID int64, msgId int, bot *telego.Bot) error {
	params := &telego.DeleteMessageParams{
		ChatID:    telego.ChatID{ID: chatID},
		MessageID: msgId,
	}
	err := bot.DeleteMessage(params)
	if err != nil {
		return err
	}
	return nil
}

func HandleUpdates(bot *telego.Bot) error {
	updates, err := bot.UpdatesViaLongPolling(nil)
	if err != nil {
		return err
	}
	defer bot.StopLongPolling()

	for update := range updates {
		// handle messages
		if update.Message != nil {
			err := HandleMessage(update, bot)
			if err != nil {
				return err
			}
		}

		// handle button callback
		if update.CallbackQuery != nil {
			err := HandleButtonCallback(update.CallbackQuery, bot)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func HandleBack(chatID int64, backRoute string, bot *telego.Bot) {
	switch backRoute {
	case Start:
		ActionStart(chatID, bot)
	}
}

func HandleConfirm(backRoute, proceedRoute string, chatID int64, msg string, bot *telego.Bot) {
	inlineKeyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow( // Row 1
			tu.InlineKeyboardButton("🔙 Back"). // Column 1
								WithCallbackData(Back+"/"+backRoute),
			tu.InlineKeyboardButton("✅ Proceed"). // Column 2
								WithCallbackData(proceedRoute),
		),
	)

	message := tu.Message(
		tu.ID(chatID),
		msg,
	).WithReplyMarkup(inlineKeyboard)

	// Sending message
	res, err := bot.SendMessage(message)
	if err != nil {
		fmt.Println(err)
	}

	state[chatID]["BotMessageID"] = res.MessageID
	_ = res

}

func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}
