package app

import (
	"bonjour_nails/config"
	"bonjour_nails/pkg/tg_bot"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

func Run(cfg *config.Config) {

	bot, err := tgbotapi.NewBotAPI(cfg.TgBot.Token)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("app - Run - tgBot: %s", err))
	}
	bot.Debug = true

	tg := tg_bot.NewBot(bot)
	tg.StartedFunc()
}
