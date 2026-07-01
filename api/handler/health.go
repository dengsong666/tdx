package handler

import (
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Health(c *gin.Context) {
	h.success(c, gin.H{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	})
}
