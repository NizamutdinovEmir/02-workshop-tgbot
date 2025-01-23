package commands

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsg := "All products: \n\n"
	products := c.productService.List()
	for _, product := range products {
		outputMsg += product.Title
		outputMsg += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)

	serialzedData, _ := json.Marshal(CommandData{
		Offset: 21,
	})

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(serialzedData)),
		),
	)

	c.bot.Send(msg)
}
