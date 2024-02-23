//package main
//
//import (
//	"TG_bot/global"
//	"TG_bot/service"
//	"fmt"
//	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
//	"log"
//)
//
//func main() {
//
//	bot, err := tgbotapi.NewBotAPI("6557688480:AAGsybjO5HMMjA0OX3Yv6MKPl6UUdPx8T24")
//	if err != nil {
//		log.Panic(err)
//		return
//	}
//	global.TelegramBot = bot
//
//	ser := service.NewTelegramService()
//
//	me, err := ser.GetMe()
//
//	//file := tgbotapi.FilePath("./oceans.mp4")
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Printf("%+v", me)
//}
//
//// chatid  -1001941525772  阿斌 -1001814790612  个人7026478261
////userid 6719549841
////阿斌 userid 6557688480

package main

import (
	"TG_bot/global"
	"TG_bot/service"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var asd = "📃Post yaratish"

// 内联键盘
var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)

// 实体键盘按钮
var numericKeyboard1 = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("👳Saralnganlar"),
		tgbotapi.NewKeyboardButton(asd),
	),
)

// 实体键盘按钮2
var numericKeyboard2 = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Text"),
		tgbotapi.NewKeyboardButton("Photo"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("GIF"),
		tgbotapi.NewKeyboardButton("video"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🉐Back"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6719549841:AAGzObsgkeOQPHUbM3x_TeeFhUWqxk4ys5U")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	global.TelegramBot = bot

	ser := service.NewTelegramService()

	// Loop through each update.
	for update := range updates {
		// Check if we've gotten a message update.
		if update.Message != nil {
			asd = "哈哈"
			msg := tgbotapi.MessageConfig{
				BaseChat: tgbotapi.BaseChat{
					ChatID:           update.Message.Chat.ID,
					ReplyToMessageID: 0,
				},
				Text:                  "11",
				DisableWebPagePreview: false,
			}

			msg.ReplyMarkup = numericKeyboard1

			fmt.Println(update.Message.Text, "我在这里打印刚刚输入的内容")

			switch update.Message.Text {
			case "open":
				msg.ReplyMarkup = numericKeyboard
			case "6":

				file := tgbotapi.FileURL("https://cdn.pixabay.com/photo/2017/03/12/11/30/alishan-2136879_1280.jpg")
				_, err := ser.SendMessagePhoto(update.Message.Chat.ID, file, "这里是图片")
				if err != nil {
					fmt.Println(err, "在这里查看错误")
					return
				}

			case "📃Post yaratish":
				fmt.Println("我这里输入5了！！！！！！！！！！！！！！！！！")
				msg.ReplyMarkup = numericKeyboard2
			case "photo":
				fmt.Println("123")
			case "复读":
				if _, err = bot.Send(msg); err != nil {
					panic(err)
				}
			}

			// Send the message.
			if _, err = bot.Send(msg); err != nil {
				panic(err)
			}
		} else if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			// And finally, send a message containing the data received.
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}

	}
}
