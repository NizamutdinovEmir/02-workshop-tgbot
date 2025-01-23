package commands

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	switch {
	case update.CallbackQuery != nil:
		c.handleCallback(update.CallbackQuery)
	case update.Message != nil:
		c.handleMessage(update.Message)
	}
}

func (c *Commander) handleCallback(callback *tgbotapi.CallbackQuery) {
	parsedData := CommandData{}
	json.Unmarshal([]byte(callback.Data), &parsedData)
	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		fmt.Sprintf("Parsed: %+v\n", parsedData),
	)
	c.bot.Send(msg)
}

func (c *Commander) handleMessage(msg *tgbotapi.Message) {
	switch msg.Command() {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	default:
		c.Default(msg)
	}
}
