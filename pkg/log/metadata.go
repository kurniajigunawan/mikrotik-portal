package log

import (
	"net/http"
)

func Metadata(code int, message string) map[string]interface{} {
	if code == 0 {
		code = http.StatusInternalServerError
	}
	return map[string]interface{}{
		"code":    code,
		"message": message,
	}
}
