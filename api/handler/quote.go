package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// Quote 获取五档行情
// @Summary 获取五档行情
// @Tags 行情
// @Produce json
// @Param code query string true "股票代码，多个用逗号分隔" example(000001,600519)
// @Param format query string false "返回格式：simple/raw，默认 simple" Enums(simple, raw)
// @Success 200 {object} ResponseDoc{data=[]SimpleQuoteDoc}
// @Router /quote [get]
func (h *Handler) Quote(c *gin.Context) {
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

	codes := splitCodes(code)
	quotes, err := h.client.GetQuote(codes...)
	if err != nil {
		h.fail(c, err.Error())
		return
	}

	if format == formatRaw {
		h.success(c, quotes)
		return
	}
	h.success(c, simpleQuotes(quotes))
}
