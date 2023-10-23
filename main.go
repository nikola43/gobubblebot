package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
)

func main() {
	//botToken := os.Getenv("TOKEN")
	// init bot
	botToken := "6230144909:AAGDs-5kxbDMR82jdfj7A3cYcIDafGpK-WY"
	//bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	bot, err := telego.NewBot(botToken, telego.WithWarnings())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Receiving updates
	err = HandleUpdates(bot)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
