package tgbot

import (
	"log"
	"tg_client/src/internal/bot/configs"

	tgbotapi "github.com/Syfaro/telegram-bot-api"

	"github.com/sirupsen/logrus"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("1", "1"),
		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Ready", "Ready"),
	),
)

//func start(config *configs.Config) *logrus.Logger

func configureLogger(config *configs.Config) *logrus.Logger {
	logger := logrus.New()

	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	logger.SetLevel(level)

	return logger
}

func createBot(config *configs.Config, logger *logrus.Logger) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(config.Apitoken)
	if err != nil {
		return nil, err
	}

	bot.Debug = false

	logger.Infof("Authorized on account %s", bot.Self.UserName)

	return bot, nil
}
