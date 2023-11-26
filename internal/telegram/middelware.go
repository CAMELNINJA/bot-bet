package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *adapter) chekUser(update tgbotapi.Update, msg tgbotapi.MessageConfig) error {
	ply, err := a.service.GetByTelegramID(int(update.Message.Chat.ID))
	if err != nil || ply.UserName == "" {
		msg.Text = StartMsg
		msg.ReplyMarkup = StartKeyboard
		return errors.New("error getting user data")
	}
	return nil
}
