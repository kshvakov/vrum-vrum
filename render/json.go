package render

import (
	"encoding/json"
	"net/http"
)

func JSON(rw http.ResponseWriter, value interface{}) error {

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(rw).Encode(value)
}
