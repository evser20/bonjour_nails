package tg_bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

const (
	commandWriteToMaster = "writeToMaster"
)

func (b *TgBot) handleCommand(message *tgbotapi.Message) {
	var names string
	msg := tgbotapi.NewMessage(message.Chat.ID, "Unknown command! Please, try again!")

	switch message.Text {
	case commandWriteToMaster:
		ms, err := b.masters.GetMasters()
		if err != nil {
			log.Error().Msg(fmt.Sprintf("error: %s", err))
		}
		for _, val := range ms {
			names += fmt.Sprintf("%s (%s) \n", val.FirstName, val.Phone)
		}
		msg = tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Можно осуществить запись к мастерам:\n\n%s", names))
		_, err = b.bot.Send(msg)
		if err != nil {
			log.Error().Msg(fmt.Sprintf("error: %s", err))
		}
	default:
		_, err := b.bot.Send(msg)
		if err != nil {
			log.Error().Msg(fmt.Sprintf("error: %s", err))
		}
	}
	//log.Error().Msg(fmt.Sprintf("[%s] %s", message.From.UserName, message.Text))

}
