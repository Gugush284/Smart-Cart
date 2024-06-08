package tgbot

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"tg_client/src/internal/bot/configs"

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
	config *configs.Config,
) {
	type message struct {
		ID int `json:"id"`
	}

	var m message

	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)

	var addres string

	switch callback.Text {
	case "1":
		m.ID = 1
	case "2":
		m.ID = 2
	case "3":
		m.ID = 3
	case "4":
		m.ID = 4
	case "5":
		m.ID = 5
	case "6":
		m.ID = 6
	case "Ready":
		addres = config.ServerAddr + "/ready/tg/" + strconv.Itoa(int(update.CallbackQuery.Message.Chat.ID))

		resp, err := http.Post(addres, "application/json", nil)
		if err != nil {
			logger.Fatalln(err)
			return
		}
		defer resp.Body.Close()

		logger.Info(resp.Status)
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Ok. Wait for your goods")
		bot.Send(msg)

		return
	default:
		logger.Info(callback.Text)
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Undef")
		bot.Send(msg)
		return
	}

	addres = config.ServerAddr + "/tg/" + strconv.Itoa(int(update.CallbackQuery.Message.Chat.ID))

	bytesRepresentation, err := json.Marshal(m)
	if err != nil {
		log.Fatalln(err)
		return
	}

	resp, err := http.Post(addres, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		logger.Fatalln(err)
		return
	}
	defer resp.Body.Close()

	logger.Info(resp.Status)

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Add good")
	bot.Send(msg)
	logger.Info("Finished processing the request")
}
