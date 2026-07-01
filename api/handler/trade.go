package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Trade(c *gin.Context) {
	code := strings.TrimSpace(c.Query("code"))
	if code == "" {
		h.fail(c, "code is required")
		return
	}

	start, ok := parseUint16(c.DefaultQuery("start", "0"))
	if !ok {
		h.fail(c, "start must be a number between 0 and 65535")
		return
	}

	count, ok := parseUint16(c.DefaultQuery("count", "100"))
	if !ok || count == 0 || count > 1800 {
		h.fail(c, "count must be a number between 1 and 1800")
		return
	}

	resp, err := h.client.GetTrade(code, start, count)
	if err != nil {
		h.fail(c, err.Error())
		return
	}

	h.success(c, resp)
}
