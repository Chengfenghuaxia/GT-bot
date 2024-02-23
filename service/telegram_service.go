package service

import (
	"TG_bot/global"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var telegramService *TelegramService

type TelegramService struct {
}

func NewTelegramService() *TelegramService {
	if telegramService == nil {
		telegramService = &TelegramService{}
	}
	return telegramService
}

// 获取自身
func (m *TelegramService) GetMe() (tgbotapi.User, error) {
	return global.TelegramBot.GetMe()
}

// 获取会话信息
func (m *TelegramService) GetChat(chatID int64, superGroupUsername string) (tgbotapi.Chat, error) {
	chatConfig := tgbotapi.ChatConfig{
		ChatID:             chatID,
		SuperGroupUsername: superGroupUsername,
	}
	chatInfoConfig := tgbotapi.ChatInfoConfig{ChatConfig: chatConfig}
	return global.TelegramBot.GetChat(chatInfoConfig)
}

// 发送文本消息
func (m *TelegramService) SendMessageText(chatID int64, text string, parseMode string) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = parseMode
	return global.TelegramBot.Send(msg)
}

// 发送图片消息
func (m *TelegramService) SendMessagePhoto(chatID int64, file tgbotapi.RequestFileData, caption string) (tgbotapi.Message, error) {
	msg := tgbotapi.NewPhoto(chatID, file)
	msg.Caption = caption
	return global.TelegramBot.Send(msg)
}

// 发送音频消息
func (m *TelegramService) SendMessageAudio(chatID int64, file tgbotapi.RequestFileData, caption string) (tgbotapi.Message, error) {
	msg := tgbotapi.NewAudio(chatID, file)
	msg.Caption = caption
	return global.TelegramBot.Send(msg)
}

// 发送视频消息
func (m *TelegramService) SendMessageVideo(chatID int64, file tgbotapi.RequestFileData, caption string) (tgbotapi.Message, error) {
	msg := tgbotapi.NewVideo(chatID, file)
	msg.Caption = caption
	return global.TelegramBot.Send(msg)
}

// 发送动图消息
func (m *TelegramService) SendMessageAnimation(chatID int64, file tgbotapi.RequestFileData, caption string) (tgbotapi.Message, error) {
	msg := tgbotapi.NewAnimation(chatID, file)
	msg.Caption = caption
	return global.TelegramBot.Send(msg)
}

// 发送文档
func (m *TelegramService) SendMessageDocument(chatID int64, file tgbotapi.RequestFileData, caption string) (tgbotapi.Message, error) {
	msg := tgbotapi.NewDocument(chatID, file)
	msg.Caption = caption
	return global.TelegramBot.Send(msg)
}

// 发送贴纸
func (m *TelegramService) SendMessageSticker(chatID int64, file tgbotapi.RequestFileData, caption string) (tgbotapi.Message, error) {
	msg := tgbotapi.NewSticker(chatID, file)
	return global.TelegramBot.Send(msg)
}

// 发送多媒体组
func (m *TelegramService) SendMediaGroup(chatID int64, files []interface{}) ([]tgbotapi.Message, error) {
	group := tgbotapi.NewMediaGroup(chatID, files)
	return global.TelegramBot.SendMediaGroup(group)
}

// 发送携带按钮组的文本消息
func (m *TelegramService) SendKeyboardButtonRow(chatID int64, text string, rows ...tgbotapi.KeyboardButton) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = tgbotapi.NewKeyboardButtonRow(rows...)
	return global.TelegramBot.Send(msg)
}

// 发送携带内连按钮组的文本消息
func SendInlineKeyboardMarkup(chatID int64, text string, rows ...[]tgbotapi.InlineKeyboardButton) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows...)
	return global.TelegramBot.Send(msg)
}

// 修改文本消息
func EditMessageText(chatID int64, messageID int, text string) (tgbotapi.Message, error) {
	msg := tgbotapi.NewEditMessageText(chatID, messageID, text)
	return global.TelegramBot.Send(msg)
}

// 修改标题
func EditMessageCaption(chatID int64, messageID int, caption string) (tgbotapi.Message, error) {
	msg := tgbotapi.NewEditMessageCaption(chatID, messageID, caption)
	return global.TelegramBot.Send(msg)
}

// 修改内联键盘标记
func EditMessageReplyMarkup(chatID int64, messageID int, markup tgbotapi.InlineKeyboardMarkup) (tgbotapi.Message, error) {
	msg := tgbotapi.NewEditMessageReplyMarkup(chatID, messageID, markup)
	return global.TelegramBot.Send(msg)
}

// 修改内联键盘标题和标记
func EditMessageTextAndMarkup(chatID int64, messageID int, text string, markup tgbotapi.InlineKeyboardMarkup) (tgbotapi.Message, error) {
	msg := tgbotapi.NewEditMessageTextAndMarkup(chatID, messageID, text, markup)
	return global.TelegramBot.Send(msg)
}

// 删除消息
func (m *TelegramService) DeleteMessage(chatID int64, messageID int) (tgbotapi.Message, error) {
	msg := tgbotapi.NewDeleteMessage(chatID, messageID)
	return global.TelegramBot.Send(msg)
}

// 创建群组邀请连接
func CreateChatInviteLink(chatID int64, superGroupUsername string, name string, expireDate int, memberLimit int, createsJoinRequest bool) (tgbotapi.Message, error) {
	config := tgbotapi.CreateChatInviteLinkConfig{
		ChatConfig: tgbotapi.ChatConfig{
			ChatID:             chatID,
			SuperGroupUsername: superGroupUsername,
		},
		Name:               name,
		ExpireDate:         expireDate,
		MemberLimit:        memberLimit,
		CreatesJoinRequest: createsJoinRequest,
	}
	return global.TelegramBot.Send(config)
}

// 修改群组邀请连接
func EditChatInviteLink(chatID int64, superGroupUsername string, name string, expireDate int, memberLimit int, createsJoinRequest bool) (tgbotapi.Message, error) {
	config := tgbotapi.EditChatInviteLinkConfig{
		ChatConfig: tgbotapi.ChatConfig{
			ChatID:             chatID,
			SuperGroupUsername: superGroupUsername,
		},
		Name:               name,
		ExpireDate:         expireDate,
		MemberLimit:        memberLimit,
		CreatesJoinRequest: createsJoinRequest,
	}
	return global.TelegramBot.Send(config)
}

// 撤销群组邀请连接
func RevokeChatInviteLink(chatID int64, superGroupUsername string, inviteLink string) (tgbotapi.Message, error) {
	config := tgbotapi.RevokeChatInviteLinkConfig{
		ChatConfig: tgbotapi.ChatConfig{
			ChatID:             chatID,
			SuperGroupUsername: superGroupUsername,
		},
		InviteLink: inviteLink,
	}
	return global.TelegramBot.Send(config)
}

// 批准聊天加入请求
func (m *TelegramService) ApproveChatJoinRequest(chatID int64, superGroupUsername string, userID int64) (tgbotapi.Message, error) {
	config := tgbotapi.ApproveChatJoinRequestConfig{
		ChatConfig: tgbotapi.ChatConfig{
			ChatID:             chatID,
			SuperGroupUsername: superGroupUsername,
		},
		UserID: userID,
	}
	return global.TelegramBot.Send(config)
}

// 拒绝聊天加入请求
func DeclineChatJoinRequest(chatID int64, superGroupUsername string, userID int64) (tgbotapi.Message, error) {
	config := tgbotapi.DeclineChatJoinRequest{
		ChatConfig: tgbotapi.ChatConfig{
			ChatID:             chatID,
			SuperGroupUsername: superGroupUsername,
		},
		UserID: userID,
	}
	return global.TelegramBot.Send(config)
}

// 离开聊天
func (m *TelegramService) LeaveChat(chatID int64, channelUsername string) (tgbotapi.Message, error) {
	config := tgbotapi.LeaveChatConfig{
		ChatID:          chatID,
		ChannelUsername: channelUsername,
	}
	return global.TelegramBot.Send(config)
}

// 包含有关聊天和用户的信息
func (m *TelegramService) ChatConfigWithUser(chatID int64, superGroupUsername string, userID int64) (tgbotapi.Message, error) {
	config := tgbotapi.GetChatMemberConfig{
		ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
			ChatID:             chatID,
			SuperGroupUsername: superGroupUsername,
			UserID:             userID,
		},
	}
	return global.TelegramBot.Send(config)
}

// 获取我的命令配置
func (m *TelegramService) GetMyCommands(ChatID, UserID int64, Type, LanguageCode string) (tgbotapi.Message, error) {
	config := tgbotapi.GetMyCommandsConfig{
		Scope: &tgbotapi.BotCommandScope{
			Type:   Type,
			ChatID: ChatID,
			UserID: UserID,
		},
		LanguageCode: LanguageCode,
	}
	return global.TelegramBot.Send(config)
}

// 获取贴纸设置配置
func (m *TelegramService) GetStickerSetConfig(name string) (tgbotapi.Message, error) {
	config := tgbotapi.GetStickerSetConfig{
		Name: name,
	}
	return global.TelegramBot.Send(config)
}

// 复制消息
func (m *TelegramService) CopyMessage(chatID, fromChatID int64, messageID int) (tgbotapi.Message, error) {
	config := tgbotapi.NewCopyMessage(chatID, fromChatID, messageID)
	return global.TelegramBot.Send(config)
}

// 删除聊天照片
func (m *TelegramService) DeleteChatPhoto(chatID int64) (tgbotapi.Message, error) {
	config := tgbotapi.NewDeleteChatPhoto(chatID)
	return global.TelegramBot.Send(config)
}

// 删除
func (m *TelegramService) DeleteMessageConfig(ChannelUsername string, chatID int64, messageID int) (tgbotapi.Message, error) {
	config := tgbotapi.DeleteMessageConfig{
		ChannelUsername: ChannelUsername,
		ChatID:          chatID,
		MessageID:       messageID,
	}
	return global.TelegramBot.Send(config)
}

func (m *TelegramService) DeleteMyCommands() (tgbotapi.Message, error) {
	config := tgbotapi.NewDeleteMyCommands()
	return global.TelegramBot.Send(config)
}

// 发送位置
func (m *TelegramService) SendLocation(chatID int64, latitude, longitude string) (tgbotapi.Message, error) {
	config := tgbotapi.NewContact(chatID, latitude, longitude)
	return global.TelegramBot.Send(config)
}

// 消息发送频道
func NewMessageToChannel(username string, text string) (tgbotapi.Message, error) {
	config := tgbotapi.NewMessageToChannel(username, text)
	return global.TelegramBot.Send(config)
}

// 内联键盘按钮数据
func InlineKeyboardButtonData(chatID int64, text, data string) (tgbotapi.Message, error) {
	rows := []tgbotapi.InlineKeyboardButton{tgbotapi.NewInlineKeyboardButtonData(text, data)}
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows)
	return global.TelegramBot.Send(msg)
}

// 内联键盘按钮登录URL
func InlineKeyboardButtonLoginURL(chatID int64, text string, loginURL tgbotapi.LoginURL) (tgbotapi.Message, error) {
	rows := []tgbotapi.InlineKeyboardButton{tgbotapi.NewInlineKeyboardButtonLoginURL(text, loginURL)}
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows)
	return global.TelegramBot.Send(msg)
}

//func InlineKeyboardRow(btn tgbotapi.InlineKeyboardButton) (tgbotapi.Message, error) {
//	rows := []tgbotapi.InlineKeyboardButton{btn}
//
//	msg := tgbotapi.NewMessage(btn.CallbackQuery.Message.Chat.ID, btn.CallbackQuery.Message.Text)
//	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows)
//
//	return global.TelegramBot.Send(msg)
//}
