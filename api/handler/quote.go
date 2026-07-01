package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Quote(c *gin.Context) {
	format, ok := responseFormat(c.DefaultQuery("format", formatSimple))
	if !ok {
		h.fail(c, "format must be one of raw, simple")
		return
	}

	code := strings.TrimSpace(c.Query("code"))
	if code == "" {
		h.fail(c, "code is required")
		return
	}

	codes := splitCodes(code)
	quotes, err := h.client.GetQuote(codes...)
	if err != nil {
		h.fail(c, err.Error())
		return
	}

	if format == formatRaw {
		h.success(c, quotes)
		return
	}
	h.success(c, simpleQuotes(quotes))
}
