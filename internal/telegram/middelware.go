package telegram

import (
	"errors"

	tgbotapi "github.com/matterbridge/telegram-bot-api/v6"
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
