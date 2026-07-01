package handler

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/injoyai/tdx/protocol"
)

// Minute 获取分时数据
// @Summary 获取分时数据
// @Tags 行情
// @Produce json
// @Param code query string true "股票代码" example(000001)
// @Param date query string false "日期，支持 20260701 或 2026-07-01"
// @Param format query string false "返回格式：simple/raw，默认 simple" Enums(simple, raw)
// @Success 200 {object} ResponseDoc{data=[]SimpleMinuteDoc}
// @Router /minute [get]
func (h *Handler) Minute(c *gin.Context) {
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
	var resp *protocol.MinuteResp
	var err error
	if date == "" {
		resp, err = h.client.GetMinute(code)
	} else {
		t, ok := parseDate(date)
		if !ok {
			h.fail(c, "date must be 20260701 or 2026-07-01")
			return
		}
		resp, err = h.client.GetHistoryMinute(t.Format("20060102"), code)
	}
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
		"list":  simpleMinutes(resp.List),
		"date":  minuteDate(date, time.Now()),
	})
}

func minuteDate(value string, fallback time.Time) string {
	if t, ok := parseDate(value); ok {
		return t.Format(time.DateOnly)
	}
	return fallback.Format(time.DateOnly)
}
