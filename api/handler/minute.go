package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Minute(c *gin.Context) {
	code := strings.TrimSpace(c.Query("code"))
	if code == "" {
		h.fail(c, "code is required")
		return
	}

	resp, err := h.client.GetMinute(code)
	if err != nil {
		h.fail(c, err.Error())
		return
	}

	h.success(c, resp)
}
