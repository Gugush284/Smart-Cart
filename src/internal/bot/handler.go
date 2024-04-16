package tgbot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

func msghandler(bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
	logger *logrus.Logger,
	commands []string,
) {
	switch update.Message.Text {
	case "/start":
		logger.Info("/start")
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello! Choose your goods:")
		msg.ReplyMarkup = numericKeyboard
		bot.Send(msg)

	default:
		logger.Info(update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Undef")
		bot.Send(msg)
	}
	logger.Info("Finished processing the request")
}
