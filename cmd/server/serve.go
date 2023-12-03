package server

import (
	apg "bonjour_nails/adapter/pg"
	. "bonjour_nails/cmd/config"
	"bonjour_nails/internal/masters"
)

func Run() (*masters.API, error) {
	c := &ServeConfig{}

	if err := Parse(c); err != nil {
		return nil, err
	}

	db, err := apg.NewBonjourDB(apg.DBConfig{
		Address:  c.PostgresAddress,
		Port:     c.PostgresPort,
		User:     c.PostgresUser,
		Password: c.PostgresPassword,
		DBName:   c.PostgresDBName,
	})
	if err != nil {
		return nil, err
	}
	repo, err := apg.NewRepository(db)
	if err != nil {
		return nil, err
	}
	masts, err := masters.NewAPI(masters.APIConfig{
		MastersReader: repo,
	})
	if err != nil {
		return nil, err
	}
	return masts, nil
}
