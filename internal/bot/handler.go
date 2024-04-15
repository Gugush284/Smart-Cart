package tgbot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

func handler(bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
	logger *logrus.Logger,
) {
	switch update.Message.Text {
	case "/start":
		logger.Info("/start")
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello!")
		bot.Send(msg)

	default:
		logger.Info(update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Undef")
		bot.Send(msg)
	}
	logger.Info("Finished processing the request")
}
