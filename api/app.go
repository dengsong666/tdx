package main

import (
	"sync"

	"github.com/injoyai/tdx"
)

type App struct {
	Config Config
	Client *tdx.Client
	Codes  tdx.ICodes

	gbbq   tdx.IGbbq
	gbbqMu sync.Mutex
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

func (a *App) StartGbbqInit() {
	go func() {
		gbbq, err := tdx.NewGbbq(tdx.WithGbbqClient(a.Client))
		if err != nil {
			return
		}
		a.gbbqMu.Lock()
		a.gbbq = gbbq
		a.gbbqMu.Unlock()
	}()
}

func (a *App) Gbbq() (tdx.IGbbq, bool) {
	a.gbbqMu.Lock()
	defer a.gbbqMu.Unlock()
	if a.gbbq != nil {
		return a.gbbq, true
	}
	return nil, false
}
