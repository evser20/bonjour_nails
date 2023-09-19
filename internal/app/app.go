package app

import (
	"bonjour_nails/config"
	tgbot "bonjour_nails/pkg/tg_bot"
	"fmt"
	"github.com/rs/zerolog/log"
)

func Run(cfg *config.Config) {
	//r := chi.NewRouter()

	b, err := tgbot.NewAPI(cfg.TgBot.Token)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("app - Run - tgBot: %s", err))
	}
	b.StartedFunc()
	//err := http.ListenAndServe(":"+cfg.HTTP.Port, r)
	//if err != nil {
	//	l.Error(fmt.Errorf("app - Run - httpServer: %w", err))
	//}
}
