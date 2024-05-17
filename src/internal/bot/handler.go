package tgbot

import (
	"bytes"
	"cart/src/internal/bot/configs"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
		First  int `json:"first"`
		Second int `json:"second"`
		Third  int `json:"third"`
	}

	var m message

	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)

	var addres string

	switch callback.Text {
	case "1":
		m.First = 1
		m.Second = 0
		m.Third = 0
	case "2":
		m.First = 0
		m.Second = 1
		m.Third = 0
	case "3":
		m.First = 0
		m.Second = 0
		m.Third = 1
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
