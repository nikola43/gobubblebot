package main

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

// actions
const (
	Back       = "back"
	Proceed    = "proceed"
	Start      = "start"
	ShowMenu   = "show_menu"
	Disconnect = "disconnect"

	ETH = "eth"
	BSC = "bsc"

	SetToken    = "set_token"
	SetGroup    = "set_group"
	SetTelegram = "set_telegram"
	SetWebsite  = "set_website"
	SetTwitter  = "set_twitter"
	StartBot    = "startbot"
	StopBot     = "stopbot"
)

var INPUT_CAPTIONS = map[string]string{
	"Welcome":   "Welcome to app",
	SetToken:    "Set token address",
	SetGroup:    "Set group link",
	SetWebsite:  "Set website link",
	SetTelegram: "Set telegram link",
	SetTwitter:  "Set twitter link",
}

var MAIN_MENU_KEYBOARD = tu.InlineKeyboard(
	tu.InlineKeyboardRow( // Row 2
		tu.InlineKeyboardButton("Buy").WithCallbackData("callback_1"),
		tu.InlineKeyboardButton("Continue").WithCallbackData("callback_1"),
		tu.InlineKeyboardButton("Continue").WithCallbackData("callback_1"),
	),
	tu.InlineKeyboardRow( // Row 2
		tu.InlineKeyboardButton("❌ Disconnect").WithCallbackData(Disconnect),
	),
)

/*
	tu.InlineKeyboardRow( // Row 2
		tu.InlineKeyboardButton("🟢 ETH").WithCallbackData(SetToken),
		tu.InlineKeyboardButton("⚪ BSC").WithCallbackData(SetGroup),
	),
*/

var CONFIG_KEYBOARD = tu.InlineKeyboard(
	tu.InlineKeyboardRow( // Row 2
		tu.InlineKeyboardButton("📄 Token").WithCallbackData(SetToken),
		tu.InlineKeyboardButton("🚻 Group").WithCallbackData(SetGroup),
	),
	tu.InlineKeyboardRow( // Row 2
		tu.InlineKeyboardButton("➤ Telegram").WithCallbackData(SetTelegram),
		tu.InlineKeyboardButton("🌐 Website").WithCallbackData(SetWebsite),
		tu.InlineKeyboardButton("𝕏 Twitter").WithCallbackData(SetTwitter),
	),
	tu.InlineKeyboardRow( // Row 2
		tu.InlineKeyboardButton("✅ Start").WithCallbackData(StartBot),
		tu.InlineKeyboardButton("❌ Stop").WithCallbackData(StopBot),
	),
)

var KEYBOARDS = map[string]*telego.InlineKeyboardMarkup{
	ShowMenu: CONFIG_KEYBOARD,
}
