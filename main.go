package main

import (
	"TG_bot/global"
	"TG_bot/service"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("6719549841:AAGzObsgkeOQPHUbM3x_TeeFhUWqxk4ys5U")
	if err != nil {
		log.Panic(err)
		return
	}
	global.TelegramBot = bot

	if err != nil {
		return
	}

	ser := service.NewTelegramService()

	me, err := ser.GetMe()
	//file := tgbotapi.FilePath("./oceans.mp4")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", me)
}

// chatid  -1001941525772
//userid 6719549841
