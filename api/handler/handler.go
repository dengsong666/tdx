package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/injoyai/tdx"
)

type Responder func(*gin.Context, any)
type ErrorResponder func(*gin.Context, string)
type GbbqProvider func() (tdx.IGbbq, bool)

type Handler struct {
	client  *tdx.Client
	codes   tdx.ICodes
	gbbq    GbbqProvider
	success Responder
	fail    ErrorResponder
}

func New(client *tdx.Client, codes tdx.ICodes, gbbq GbbqProvider, success Responder, fail ErrorResponder) *Handler {
	return &Handler{
		client:  client,
		codes:   codes,
		gbbq:    gbbq,
		success: success,
		fail:    fail,
	}
}
