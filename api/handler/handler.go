package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/injoyai/tdx"
)

type Responder func(*gin.Context, any)
type ErrorResponder func(*gin.Context, string)

type Handler struct {
	client  *tdx.Client
	codes   tdx.ICodes
	success Responder
	fail    ErrorResponder
}

func New(client *tdx.Client, codes tdx.ICodes, success Responder, fail ErrorResponder) *Handler {
	return &Handler{
		client:  client,
		codes:   codes,
		success: success,
		fail:    fail,
	}
}
