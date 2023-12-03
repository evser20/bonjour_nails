package tg_bot

import (
	"bonjour_nails/internal/masters"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgBot struct {
	bot     *tgbotapi.BotAPI
	masters *masters.API
}

func NewBot(bot *tgbotapi.BotAPI, m *masters.API) *TgBot {
	return &TgBot{bot, m}
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

		b.handleCommand(update.Message)
	}
}
