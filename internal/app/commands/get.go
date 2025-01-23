package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)

	if err != nil {
		log.Printf("Invalid product ID: %s", args)
		return
	}

	product, err := c.productService.Get(idx)
	if err != nil {
		log.Printf("Error getting product: %v", err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Error getting product")
		c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		product.Title,
	)

	c.bot.Send(msg)
}
