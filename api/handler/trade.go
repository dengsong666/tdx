package handler

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/injoyai/tdx/protocol"
)

// Trade 获取分时成交
// @Summary 获取分时成交
// @Tags 行情
// @Produce json
// @Param code query string true "股票代码" example(000001)
// @Param date query string false "日期，支持 20260701 或 2026-07-01"
// @Param start query int false "起始偏移，默认 0" minimum(0)
// @Param count query int false "数量，默认 100" minimum(1)
// @Param format query string false "返回格式：simple/raw，默认 simple" Enums(simple, raw)
// @Success 200 {object} ResponseDoc{data=[]SimpleTradeDoc}
// @Router /trade [get]
func (h *Handler) Trade(c *gin.Context) {
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

	date := strings.TrimSpace(c.Query("date"))
	tradeDate, hasDate, validDate := parseOptionalDate(date)
	if !validDate {
		h.fail(c, errInvalidDate.Error())
		return
	}

	start, ok := parseUint16(c.DefaultQuery("start", "0"))
	if !ok {
		h.fail(c, "start must be a number between 0 and 65535")
		return
	}

	count, ok := parseUint16(c.DefaultQuery("count", "100"))
	maxCount := uint16(1800)
	if hasDate && !sameDay(tradeDate, time.Now()) {
		maxCount = 2000
	}
	if !ok || count == 0 || count > maxCount {
		h.fail(c, "count is out of range")
		return
	}

	resp, err := h.trade(code, tradeDate, hasDate, start, count)
	if err != nil {
		h.fail(c, err.Error())
		return
	}

	if format == formatRaw {
		h.success(c, resp)
		return
	}
	h.success(c, gin.H{
		"count": resp.Count,
		"list":  simpleTrades(resp.List),
		"date":  minuteDate(date, time.Now()),
	})
}

func (h *Handler) trade(code string, date time.Time, hasDate bool, start, count uint16) (*protocol.TradeResp, error) {
	if !hasDate {
		return h.client.GetTrade(code, start, count)
	}
	if sameDay(date, time.Now()) {
		return h.client.GetTrade(code, start, count)
	}
	return h.client.GetHistoryTrade(date.Format("20060102"), code, start, count)
}
