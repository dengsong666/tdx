package handler

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/injoyai/tdx/protocol"
)

func (h *Handler) Kline(c *gin.Context) {
	code := strings.TrimSpace(c.Query("code"))
	if code == "" {
		h.fail(c, "code is required")
		return
	}

	klineType := c.DefaultQuery("type", "day")
	start, ok := parseUint16(c.DefaultQuery("start", "0"))
	if !ok {
		h.fail(c, "start must be a number between 0 and 65535")
		return
	}

	count, ok := parseUint16(c.DefaultQuery("count", "200"))
	if !ok || count == 0 || count > 800 {
		h.fail(c, "count must be a number between 1 and 800")
		return
	}

	resp, err := h.client.GetKline(klineTypeValue(klineType), code, start, count)
	if err != nil {
		h.fail(c, err.Error())
		return
	}

	h.success(c, resp)
}

func klineTypeValue(s string) uint8 {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "minute", "minute1", "1m":
		return protocol.TypeKlineMinute
	case "minute5", "5m":
		return protocol.TypeKline5Minute
	case "minute15", "15m":
		return protocol.TypeKline15Minute
	case "minute30", "30m":
		return protocol.TypeKline30Minute
	case "hour", "minute60", "60m":
		return protocol.TypeKline60Minute
	case "week":
		return protocol.TypeKlineWeek
	case "month":
		return protocol.TypeKlineMonth
	case "quarter":
		return protocol.TypeKlineQuarter
	case "year":
		return protocol.TypeKlineYear
	default:
		return protocol.TypeKlineDay
	}
}

func parseUint16(s string) (uint16, bool) {
	n, err := strconv.ParseUint(strings.TrimSpace(s), 10, 16)
	if err != nil {
		return 0, false
	}
	return uint16(n), true
}
