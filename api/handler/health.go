package handler

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Health 健康检查
// @Summary 健康检查
// @Tags 系统
// @Produce json
// @Success 200 {object} ResponseDoc
// @Router /health [get]
func (h *Handler) Health(c *gin.Context) {
	h.success(c, gin.H{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	})
}
