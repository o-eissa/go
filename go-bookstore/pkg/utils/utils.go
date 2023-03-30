package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, X interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		if err := json.Unmarshal([]byte(body), X); err != nil {
			return
		}
	}

}
