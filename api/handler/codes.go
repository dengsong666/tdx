package handler

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Codes(c *gin.Context) {
	kind := strings.ToLower(strings.TrimSpace(c.DefaultQuery("type", "stocks")))
	limit := parseLimit(c.DefaultQuery("limit", "200"))

	switch kind {
	case "stock", "stocks":
		h.success(c, h.codes.GetStocks(limit))
	case "etf", "etfs":
		h.success(c, h.codes.GetETFs(limit))
	case "index", "indexes", "indices":
		h.success(c, h.codes.GetIndexes(limit))
	case "all":
		data := gin.H{
			"stocks":  h.codes.GetStocks(limit),
			"etfs":    h.codes.GetETFs(limit),
			"indexes": h.codes.GetIndexes(limit),
		}
		h.success(c, data)
	default:
		h.fail(c, "type must be one of stocks, etfs, indexes, all")
	}
}

func parseLimit(s string) int {
	limit, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil || limit <= 0 {
		return 200
	}
	if limit > 5000 {
		return 5000
	}
	return limit
}
