package main

import (
	"github.com/injoyai/tdx"
)

type App struct {
	Config Config
	Client *tdx.Client
	Codes  tdx.ICodes
}

func NewApp(cfg Config) (*App, error) {
	client, err := tdx.DialDefault(tdx.WithDebug(false))
	if err != nil {
		return nil, err
	}

	codes, err := tdx.NewCodes(tdx.WithCodesClient(client))
	if err != nil {
		client.Close()
		return nil, err
	}
	tdx.DefaultCodes = codes

	return &App{
		Config: cfg,
		Client: client,
		Codes:  codes,
	}, nil
}

func (a *App) Close() error {
	if a.Client == nil {
		return nil
	}
	return a.Client.Close()
}
