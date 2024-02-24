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
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "没有定义")

			msg.ReplyMarkup = numericKeyboard1

			fmt.Println(update.Message.Text, "我在这里打印刚刚输入的内容")
			if update.Message.Photo != nil && len(update.Message.Photo) > 0 {
				text := "​按以下格式发送链接：\n[按钮文字+链接]\n\n例子：\n[翻译+https://t.me/TransioBot]\n\n要在一行中添加多个按钮，请在前面的按钮旁边写下链接。\n格式：\n[第一个文本 + 第一个链接] [第二个文本 + 第二个链接]\n\n要将多个按钮添加到一行，请从新行写入新链接。\n格式：\n[第一个文字+第一个链接]\n[第二条文字+第二条链接]"
				ser.SendMessageText(update.Message.Chat.ID, text, "")
				return
			}
			switch update.Message.Text {
			case "open":
				msg.Text = "1233"
				msg.ReplyMarkup = numericKeyboard
				break
			case "6":

				file := tgbotapi.FileID("AgACAgIAAxkBAAIB72XYuugAAX-HkbphYKJk-OT22J5IewACW9QxG6klyEo43dfrCTb_vQEAAwIAA3MAAzQE")
				var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonURL("链接按钮", "https://t.me/afeng12a"),
					),
				)
				msg.ReplyMarkup = numericKeyboard
				_, err := ser.SendMessagePhoto(update.Message.Chat.ID, file, "")
				if err != nil {
					fmt.Println(err, "在这里查看错误")
					return
				}

				break
			case "7":
				_, _ = ser.CopyMessage(5903362571, update.Message.Chat.ID, 543)
				//if err != nil {
				//	return
				//}
				//msg.Text = message.Text
				//msg.ReplyMarkup = message.ReplyMarkup
				//if _, err = bot.Send(msg); err != nil {
				//	panic(err)
				//}
				break

			case "📃Post yaratish":
				fmt.Println("入我这里输5了！！！！！！！！！！！！！！！！！")
				msg.Text = "我这里输了：📃Post yaratish"
				msg.ReplyMarkup = numericKeyboard2
				break
			case "Photo":
				msg.Text = "请发图片"
				break
			case "复读":
				if _, err = bot.Send(msg); err != nil {
					panic(err)
				}
				break
			default:
				msg.Text = "没有定义"
				break
			}

			// Send the message.
			//if _, err = bot.Send(msg); err != nil {
			//	panic(err)
			//}
		} else if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}
			fmt.Println("---->" + update.CallbackQuery.Data)
			// And finally, send a message containing the data received.
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}

	}
}

//file_id:AgACAgIAAxkBAAIB72XYuugAAX-HkbphYKJk-OT22J5IewACW9QxG6klyEo43dfrCTb_vQEAAwIAA3MAAzQE
