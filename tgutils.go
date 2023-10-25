package main

import (
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
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
			SendMessage(chatID, txt, nil, bot, false)
			return nil
		}

		pair, _ := GetTokenInfoDex(msgText)
		//pair := "0x7ce52dc08b49e0bc11252cd267d2614dfc616d9f"

		/*
			tokenInfo, err := GetTokenInfo(msgText)
			if err != nil {
				return err
			}
		*/

		//fmt.Println(tokenInfo)
		fmt.Println("pair", pair.PairAddress)
		tokenConfig := state[chatID]["TokenConfig"].(TokenConfig)
		tokenConfig.Address = msgText
		tokenConfig.Pair = *pair.PairAddress
		tokenConfig.Symbol = *pair.BaseToken.Symbol
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
		fmt.Println(msgText)
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
	fmt.Println("msg")
	fmt.Println(msg)
	HandleActionWithKeyboard(chatID, ShowMenu, msg, bot)
	return nil
}

/*
LONG Buy!
⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️⚪️

💠  0.4786 ETH $801.70764000
🧩  1 LONG
💵 $763.70000000
❌ Not New Holder!
📂 Address | TX

🔘 Market Cap $7,637,086
📈 Price Impact 0.00%
⭐️ 24h Volume $3,226,726
🧸 Holders 16538
🔪 Taxes B/S | 2.0000000000000058/2.0062588425094834

Chart ▫️ Buy
Website ▫️ Twitter ▫️ Telegram
*/

func BuildBubbleMessage(tokenConfig TokenConfig, userAddress, tx string, ethAmount, tokenAmount *big.Int, marketCap, volume24, holders, buyTax, sellTax string) string {
	ethValue := ToDecimal(ethAmount, 18)
	ethValueFloat, _ := ethValue.Float64()
	tokenValue := ToDecimal(tokenAmount, 18).String()

	emoji := "🟢"
	emojiText := "🟢"
	emojiValue := 0.01
	emojiCount := ethValueFloat / emojiValue
	for i := 0; i < int(emojiCount); i++ {
		emojiText += emoji
	}

	msg := "*" + tokenConfig.Symbol + " Buy!" + "*\n\n"
	msg += emojiText + "\n\n"
	msg += "*" + ethValue.String() + " ETH" + "*\n\n"
	msg += "*" + tokenValue + " " + tokenConfig.Symbol + "*\n\n"
	msg += "*[Address](https://etherscan.io/address/" + userAddress + ") | [Tx](https://etherscan.io/tx/" + tx + ")" + "*\n\n"
	msg += "\n"
	msg += "🔘 *Market Cap $" + marketCap + "*\n\n"
	msg += "🔘 *24h Volume $" + volume24 + "*\n\n"
	msg += "🧸 *Holders " + holders + "*\n\n"
	msg += "🔪 *Taxes B/S " + buyTax + "/" + sellTax + "*\n\n"
	msg += "\n"
	msg += "*[Chart](https://dexscreener.com/ethereum/" + tokenConfig.Address + ") | [Buy](https://app.uniswap.org/tokens/ethereum/" + tokenConfig.Address + ")" + "*\n\n"
	msg += "*[Website](" + tokenConfig.Website + ") | [Twitter](" + tokenConfig.Twitter + ") | [Telegram](" + tokenConfig.Telegram + ")" + "*\n\n"

	return msg
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
		ActionStartBot(chatID, bot)

	case StopBot:
		ActionStopBot(chatID, bot)
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
		_, err := SendMessage(chatID, "welcome", nil, bot, false)
		if err != nil {
			return err
		}

		if state[chatID]["TokenConfig"] == nil {
			tokenConfig := TokenConfig{}
			state[chatID]["TokenConfig"] = tokenConfig
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

func SendMessage(chatID int64, msg string, replyMarkup telego.ReplyMarkup, bot *telego.Bot, withMarkdown bool) (*telego.Message, error) {
	var message *telego.SendMessageParams

	if replyMarkup != nil {
		message = tu.Message(
			tu.ID(chatID),
			msg,
		).WithReplyMarkup(replyMarkup).WithDisableWebPagePreview()
		if withMarkdown {
			message.WithParseMode("MarkdownV2")
		}
	} else {
		message = tu.Message(
			tu.ID(chatID),
			msg,
		).WithDisableWebPagePreview()

		if withMarkdown {
			message.WithParseMode("MarkdownV2")
		}

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
