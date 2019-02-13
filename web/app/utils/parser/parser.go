package parser

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// ParseRequestBody parse request body to dto
func ParseRequestBody(r *http.Request, data interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	r.Body.Close()
	err = json.Unmarshal(body, data)
	return err
}
