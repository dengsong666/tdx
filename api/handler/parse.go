package handler

import "strings"

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
