package tgbot

import (
	"cart/src/internal/bot/configs"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

func Start(config *configs.Config) {
	logger := configureLogger(config)

	logger.Info("Create bot")

	bot, err := createBot(config, logger)
	if err != nil {
		logger.Error(err)
	}

	logger.Info("Get bot update")

	updatesConfig := tgbotapi.NewUpdate(0)
	updatesConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updatesConfig)
	if err != nil {
		logger.Error(err)
	}

	updates.Clear()

	Serve(bot, updates, logger)
}

func Serve(
	bot *tgbotapi.BotAPI,
	updates tgbotapi.UpdatesChannel,
	logger *logrus.Logger,
) {
	logger.Info("Serve")

	commands := []string{"/start"}

	for update := range updates {
		logger.Info(update)

		if update.Message != nil {
			logger.Info(update.Message)

			if update.Message.Text != "" {
				msghandler(bot, update, logger, commands)
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Undef")
				bot.Send(msg)
			}
		} else if update.CallbackQuery != nil {
			logger.Info(update.CallbackQuery)

			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			logger.Info(callback.Text)
		}
	}
}
