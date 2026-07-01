package handler

import (
	"errors"
	"strings"
	"time"
)

var errInvalidDate = errors.New("date must be 20260701 or 2026-07-01")
var errGbbqPreparing = errors.New("复权数据准备中")

// splitCodes 解析逗号分隔的股票代码。
func splitCodes(s string) []string {
	parts := strings.Split(s, ",")
	codes := make([]string, 0, len(parts))
	for _, part := range parts {
		code := strings.TrimSpace(part)
		if code != "" {
			codes = append(codes, code)
		}
	}
	return codes
}

// parseDate 支持 20260701 和 2026-07-01。
func parseDate(s string) (time.Time, bool) {
	s = strings.TrimSpace(s)
	if s == "" {
		return time.Time{}, false
	}
	for _, layout := range []string{"20060102", time.DateOnly} {
		t, err := time.ParseInLocation(layout, s, time.Local)
		if err == nil {
			return t, true
		}
	}
	return time.Time{}, false
}

func parseOptionalDate(s string) (time.Time, bool, bool) {
	s = strings.TrimSpace(s)
	if s == "" {
		return time.Time{}, false, true
	}
	t, ok := parseDate(s)
	return t, true, ok
}

func sameDay(a, b time.Time) bool {
	ay, am, ad := a.Date()
	by, bm, bd := b.Date()
	return ay == by && am == bm && ad == bd
}
