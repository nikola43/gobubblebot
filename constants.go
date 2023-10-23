package main

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

// actions
const (
	Back         = "back"
	Proceed      = "proceed"
	Start        = "start"
	ShowMenu     = "show_menu"
	Disconnect   = "disconnect"

	SetToken    = "set_token"
	SetGroup    = "set_group"
	SetTelegram = "set_telegra"
	SetWebsite  = "set_website"
	SetTwitter  = "set_twitter"
	StartBot    = "startbot"
	StopBot     = "stopbot"
)

var INPUT_CAPTIONS = map[string]string{
	"Welcome":   "Welcome to app",
	SetToken:    "Set token",
	SetGroup:    "Set group",
	SetWebsite:  "Set website",
	SetTelegram: "Set telegram",
	SetTwitter:  "Set twitter",
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

var CONFIG_KEYBOARD = tu.InlineKeyboard(
	tu.InlineKeyboardRow( // Row 2
		tu.InlineKeyboardButton("Set Token").WithCallbackData(SetToken),
		tu.InlineKeyboardButton("Set Group").WithCallbackData(SetGroup),
	),
	tu.InlineKeyboardRow( // Row 2
		tu.InlineKeyboardButton("Set Telegram").WithCallbackData(SetTelegram),
		tu.InlineKeyboardButton("Set Website").WithCallbackData(SetWebsite),
		tu.InlineKeyboardButton("Set Twitter").WithCallbackData(SetTwitter),
	),
	tu.InlineKeyboardRow( // Row 2
		tu.InlineKeyboardButton("Start").WithCallbackData(StartBot),
		tu.InlineKeyboardButton("Stop").WithCallbackData(StopBot),
	),
)

var KEYBOARDS = map[string]*telego.InlineKeyboardMarkup{
	ShowMenu: CONFIG_KEYBOARD,
}
