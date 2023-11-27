package telegram

import (
	"log/slog"
	"strconv"

	"github.com/CAMELNINJA/bot-bet.git/internal/models"
	tgbotapi "github.com/matterbridge/telegram-bot-api/v6"
)

func (a *adapter) Listener() error {
	bot, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		a.log.Error("error creating bot", err)
		return err
	}
	a.bot = bot

	a.bot.Debug = true

	a.log.Info("Authorized on account ", slog.String("name", bot.Self.UserName))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := a.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.Chat.Type != "private" {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				if err := a.chekUser(update, msg); err == nil {
					msg.ReplyMarkup = a.getMainKeyboard(int(update.Message.Chat.ID))
					msg.Text = AlreadyRegistered
				}
			case "addbalance":
				if err := a.chekUser(update, msg); err == nil {
					balanceStr := update.Message.CommandArguments()
					if balanceStr == "" {
						msg.Text = "Введите сумму пополнения"
					} else {
						ply, err := a.service.GetByTelegramID(int(update.Message.Chat.ID))
						if err != nil {
							a.log.Error("error getting user data", err)
							msg.Text = "Ошибка"
						}
						balance, err := strconv.Atoi(balanceStr)
						if err != nil {
							a.log.Error("error converting string to int", err)
							msg.Text = "Ошибка"
						}

						ply.FactBalance = ply.FactBalance + balance
						if err := a.service.Update(ply); err != nil {
							a.log.Error("error updating user data", err)
							msg.Text = "Ошибка"
						}
						msg.Text = "Баланс пополнен, дождитесь подтверждения администратора"

						msg.ReplyMarkup = tgbotapi.NewKeyboardButtonWebApp("Сделать ставку", tgbotapi.WebAppInfo{URL: a.webAppUrl})
					}
				}
			default:
				msg.Text = "Я не знаю такую команду, мне известны только [/start , /addbalance]"
			}
		} else {
			switch update.Message.Text {
			case RegisterButton:
				ply, err := a.service.GetByTelegramID(int(update.Message.Chat.ID))
				if err != nil || ply.UserName == "" {

					user := models.User{
						UserName: update.Message.Chat.UserName,
						ChatID:   int(update.Message.Chat.ID),
					}
					if err = a.service.Create(&user); err != nil {
						msg.Text = "Анлаки чет не получилось зарегаться"
						msg.ReplyMarkup = StartKeyboard
					} else {
						msg.Text = RegistrationSuccessful
						msg.ReplyMarkup = AddBalanceKeyboard
					}

				} else {
					msg.Text = AlreadyRegistered
				}
			case SupportButton:
				msg.Text = SupportText
			case AddBalanceKeyboardButton:
				if err := a.chekUser(update, msg); err == nil {
					msg.Text = "Пополнить баланс /addbalance + сумма"
				}

			default:
				msg.Text = "Я не знаю такую команду, мне известны только [/start , /addbalance]"
			}
		}
		bot.Send(msg)
	}
	return nil
}

func (a *adapter) StopBot() {
	a.bot.StopReceivingUpdates()
}
