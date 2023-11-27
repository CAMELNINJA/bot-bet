package telegram

import (
	"strconv"

	tgbotapi "github.com/matterbridge/telegram-bot-api/v6"
)

const (
	StartMsg               = `Привет! Это бот по ставкам телеграмме`
	RegistrationSuccessful = "Бот активирован! Пополните баланс для начала работы с ботом."
	AlreadyRegistered      = "С возвращением тебя!"
	MuteModeActivated      = "Уведомления отключены до конца дня!"
	SupportText            = "Напишите нам: \n\n@L9camel"
)

const (
	RegisterButton           = "🔑 Регистрация"
	SupportButton            = "💬 Служба поддержки"
	AddBalanceKeyboardButton = "💰 Пополнить баланс"
)

const (
	StartKeyboardType = iota + 1
	MainKeyboardType
)

var StartKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(RegisterButton),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(SupportButton),
	))

var AddBalanceKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(AddBalanceKeyboardButton),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(SupportButton),
	))

func (a *adapter) getMainKeyboard(sessionID int) tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonWebApp("Сделать ставку", tgbotapi.WebAppInfo{
				URL: a.webAppUrl + "?session_id=" + strconv.Itoa(sessionID),
			}),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(SupportButton),
		),
	)
}
