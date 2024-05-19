package tgbot

import (
	"tg_client/src/internal/bot/configs"

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

	Serve(bot, updates, logger, config)
}

func Serve(
	bot *tgbotapi.BotAPI,
	updates tgbotapi.UpdatesChannel,
	logger *logrus.Logger,
	config *configs.Config,
) {
	logger.Info("Serve")

	for update := range updates {
		logger.Info(update)

		if update.Message != nil {
			logger.Info(update.Message)

			if update.Message.Text != "" {
				msghandler(bot, update, logger)
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Undef")
				bot.Send(msg)
			}
		} else if update.CallbackQuery != nil {
			logger.Info(update.CallbackQuery)

			callbackhandler(bot, update, logger, config)
		}
	}
}
