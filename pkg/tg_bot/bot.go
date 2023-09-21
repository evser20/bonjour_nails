package tg_bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

type TgBot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *TgBot {
	return &TgBot{bot}
}

func (b *TgBot) StartedFunc() {
	updates := b.initUpdatesChannel()

	b.handleUpdates(updates)
}

func (b *TgBot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}

func (b *TgBot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		b.handleMessage(update.Message)
	}
}

func (b *TgBot) handleMessage(message *tgbotapi.Message) {
	log.Error().Msg(fmt.Sprintf("[%s] %s", message.From.UserName, message.Text))

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	_, err := b.bot.Send(msg)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("error: %s", err))
	}
}
