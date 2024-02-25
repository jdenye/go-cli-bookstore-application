package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(res *http.Request, x interface{}) {
	if body, err := io.ReadAll(res.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
