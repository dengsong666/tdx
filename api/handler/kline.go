package handler

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/injoyai/tdx/protocol"
)

func (h *Handler) Kline(c *gin.Context) {
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

	klineType := c.DefaultQuery("type", "day")
	value, ok := klineTypeValue(klineType)
	if !ok {
		h.fail(c, "type must be one of minute, minute5, minute15, minute30, hour, day, week, month, quarter, year")
		return
	}

	adjust := strings.ToLower(strings.TrimSpace(c.DefaultQuery("adjust", "qfq")))
	if adjust != "none" && adjust != "qfq" && adjust != "hfq" {
		h.fail(c, "adjust must be one of none, qfq, hfq")
		return
	}

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

	resp, err := h.kline(code, value, start, count, c.Query("from"), c.Query("to"))
	if err != nil {
		h.fail(c, err.Error())
		return
	}

	klines := protocol.Klines(resp.List)
	if adjust != "none" {
		klines, err = h.adjustKlines(code, klines, adjust)
		if err != nil {
			h.fail(c, err.Error())
			return
		}
		resp.List = klines
		resp.Count = uint16(len(klines))
	}

	if format == formatRaw {
		h.success(c, resp)
		return
	}
	h.success(c, gin.H{
		"count": resp.Count,
		"list":  simpleKlines(klines),
	})
}

func (h *Handler) kline(code string, value uint8, start, count uint16, fromValue, toValue string) (*protocol.KlineResp, error) {
	from, hasFrom, validFrom := parseOptionalDate(fromValue)
	if !validFrom {
		return nil, errInvalidDate
	}
	to, hasTo, validTo := parseOptionalDate(toValue)
	if !validTo {
		return nil, errInvalidDate
	}
	if !hasFrom && !hasTo {
		return h.client.GetKline(value, code, start, count)
	}

	resp, err := h.client.GetKlineAll(value, code)
	if err != nil {
		return nil, err
	}

	list := make([]*protocol.Kline, 0, len(resp.List))
	for _, item := range resp.List {
		if hasFrom && item.Time.Before(from) {
			continue
		}
		if hasTo && item.Time.After(endOfDay(to)) {
			continue
		}
		list = append(list, item)
	}

	return &protocol.KlineResp{
		Count: uint16(len(list)),
		List:  list,
	}, nil
}

func (h *Handler) adjustKlines(code string, klines protocol.Klines, adjust string) (protocol.Klines, error) {
	gbbq, ok := h.gbbq()
	if !ok {
		return nil, errGbbqPreparing
	}
	switch adjust {
	case "hfq":
		return protocol.ApplyHFQ(klines, gbbq.GetFactors(code, klines)), nil
	case "qfq":
		return protocol.ApplyQFQ(klines, gbbq.GetFactors(code, klines)), nil
	default:
		return klines, nil
	}
}

func endOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// klineTypeValue 将接口参数映射为通达信协议常量。
func klineTypeValue(s string) (uint8, bool) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "minute":
		return protocol.TypeKlineMinute, true
	case "minute5":
		return protocol.TypeKline5Minute, true
	case "minute15":
		return protocol.TypeKline15Minute, true
	case "minute30":
		return protocol.TypeKline30Minute, true
	case "hour":
		return protocol.TypeKline60Minute, true
	case "day":
		return protocol.TypeKlineDay, true
	case "week":
		return protocol.TypeKlineWeek, true
	case "month":
		return protocol.TypeKlineMonth, true
	case "quarter":
		return protocol.TypeKlineQuarter, true
	case "year":
		return protocol.TypeKlineYear, true
	default:
		return 0, false
	}
}

// parseUint16 解析协议分页参数。
func parseUint16(s string) (uint16, bool) {
	n, err := strconv.ParseUint(strings.TrimSpace(s), 10, 16)
	if err != nil {
		return 0, false
	}
	return uint16(n), true
}
