package constants

import (
	http "net/http"
	"strings"
)

func GetHttpStatusText(httpStatus int) string {
	if text := http.StatusText(httpStatus); text != "" {
		upper := strings.ToUpper(text)
		return strings.ReplaceAll(upper, " ", "_")
	}

	return "INTERNAL_SERVER_ERROR"
}
