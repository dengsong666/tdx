package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Quote(c *gin.Context) {
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

	h.success(c, quotes)
}
