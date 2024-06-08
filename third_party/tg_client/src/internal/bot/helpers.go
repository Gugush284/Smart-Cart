package tgbot

import (
	"log"
	"tg_client/src/internal/bot/configs"

	tgbotapi "github.com/Syfaro/telegram-bot-api"

	"github.com/sirupsen/logrus"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Перловка", "1"),
		tgbotapi.NewInlineKeyboardButtonData("Жвачка", "2"),
		tgbotapi.NewInlineKeyboardButtonData("Молочный коктейль", "3"),
		tgbotapi.NewInlineKeyboardButtonData("Мюсли", "4"),
		tgbotapi.NewInlineKeyboardButtonData("Рис", "5"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Сок", "6"),
		tgbotapi.NewInlineKeyboardButtonData("Булочка", "7"),
		tgbotapi.NewInlineKeyboardButtonData("Торт", "8"),
		tgbotapi.NewInlineKeyboardButtonData("Йогурт", "9"),
		tgbotapi.NewInlineKeyboardButtonData("Семечки", "10"),
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
