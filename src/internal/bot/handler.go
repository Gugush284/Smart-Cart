package tgbot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

func msghandler(bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
	logger *logrus.Logger,
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

func callbackhandler(bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
	logger *logrus.Logger,
) {
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)

	switch callback.Text {
	case "1":
		logger.Info("First")
	default:
		logger.Info(callback.Text)
		msg := tgbotapi.NewMessage(update.CallbackQuery.ID, "Undef")
		bot.Send(msg)
		return
	}
	msg := tgbotapi.NewMessage(update.CallbackQuery.ID, "Add good")
	bot.Send(msg)
	logger.Info("Finished processing the request")
}