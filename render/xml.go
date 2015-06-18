package render

import (
	"encoding/xml"
	"net/http"
)

func XML(rw http.ResponseWriter, value interface{}) error {

	rw.Header().Set("Content-Type", "text/xml; charset=utf-8")

	return xml.NewEncoder(rw).Encode(value)
}
