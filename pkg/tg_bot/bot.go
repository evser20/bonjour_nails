package tg_bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

type TgBot struct {
	botAPI *tgbotapi.BotAPI
}

func NewAPI(token string) (*TgBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	return &TgBot{
		botAPI: bot,
	}, err
}

func (b *TgBot) StartedFunc() {
	b.botAPI.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.botAPI.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Error().Msg(fmt.Sprintf("[%s] %s", update.Message.From.UserName, update.Message.Text))

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			_, err := b.botAPI.Send(msg)
			if err != nil {
				log.Error().Msg(fmt.Sprintf("error: %s", err))
			}
		}
	}
}
