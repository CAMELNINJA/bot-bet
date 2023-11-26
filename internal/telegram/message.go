package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	StartMsg               = `Привет! Это бот по ставкам телеграмме`
	RegistrationSuccessful = "Бот активирован! Пополните баланс для начала работы с ботом."
	AlreadyRegistered      = "С возвращением тебя!"
	MuteModeActivated      = "Уведомления отключены до конца дня!"
	SupportText            = "Напишите нам: \n\n@came1l"
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
var MainKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(SupportButton),
	),
)
