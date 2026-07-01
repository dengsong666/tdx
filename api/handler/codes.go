package handler

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/injoyai/tdx"
)

// Codes 获取证券代码
// @Summary 获取证券代码
// @Tags 代码
// @Produce json
// @Param type query string false "类型，默认 stocks" Enums(stocks, etfs, indexes, all)
// @Param limit query int false "数量限制，默认 200，最大 5000" minimum(1) maximum(5000)
// @Param keyword query string false "关键词，匹配代码、完整代码或名称"
// @Success 200 {object} ResponseDoc
// @Router /codes [get]
func (h *Handler) Codes(c *gin.Context) {
	kind := strings.ToLower(strings.TrimSpace(c.DefaultQuery("type", "stocks")))
	limit := parseLimit(c.DefaultQuery("limit", "200"))
	keyword := strings.TrimSpace(c.Query("keyword"))

	switch kind {
	case "stocks":
		h.success(c, limitCodes(filterCodes(h.codes.GetStocks(), keyword), limit))
	case "etfs":
		h.success(c, limitCodes(filterCodes(h.codes.GetETFs(), keyword), limit))
	case "indexes":
		h.success(c, limitCodes(filterCodes(h.codes.GetIndexes(), keyword), limit))
	case "all":
		data := gin.H{
			"stocks":  limitCodes(filterCodes(h.codes.GetStocks(), keyword), limit),
			"etfs":    limitCodes(filterCodes(h.codes.GetETFs(), keyword), limit),
			"indexes": limitCodes(filterCodes(h.codes.GetIndexes(), keyword), limit),
		}
		h.success(c, data)
	default:
		h.fail(c, "type must be one of stocks, etfs, indexes, all")
	}
}

// filterCodes 按代码、完整代码或名称过滤。
func filterCodes(items tdx.CodeModels, keyword string) tdx.CodeModels {
	keyword = strings.ToLower(strings.TrimSpace(keyword))
	if keyword == "" {
		return items
	}
	result := make(tdx.CodeModels, 0, len(items))
	for _, item := range items {
		if strings.Contains(strings.ToLower(item.Code), keyword) ||
			strings.Contains(strings.ToLower(item.FullCode()), keyword) ||
			strings.Contains(strings.ToLower(item.Name), keyword) {
			result = append(result, item)
		}
	}
	return result
}

func limitCodes(items tdx.CodeModels, limit int) tdx.CodeModels {
	if limit <= 0 || len(items) <= limit {
		return items
	}
	return items[:limit]
}

// parseLimit 限制列表接口返回量。
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
