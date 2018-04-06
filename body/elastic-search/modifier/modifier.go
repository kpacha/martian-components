// Package modifier exposes a request modifier for generating parametrized queries
// to an elastic search service
package modifier

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/martian/parse"
)

type ESQueryModifier struct {
	Keys []string `json:"keys"`
}

type Search struct {
	Query Query `json:"query"`
}

type Query struct {
	Match map[string]string `json:"match"`
}

// ModifyRequest modifies the query string of the request with the given key and value.
func (m *ESQueryModifier) ModifyRequest(req *http.Request) error {
	query := req.URL.Query()
	search := Search{Query{map[string]string{}}}
	for _, k := range m.Keys {
		search.Query.Match[k] = query.Get(k)
		query.Del(k)
	}
	data, err := json.Marshal(search)
	if err != nil {
		return err
	}

	req.Body.Close()

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Method = http.MethodPost
	req.ContentLength = int64(len(data))
	req.Body = ioutil.NopCloser(bytes.NewReader(data))
	req.URL.RawQuery = query.Encode()

	return nil
}

func FromJSON(b []byte) (*parse.Result, error) {
	msg := &ESQueryModifier{}
	if err := json.Unmarshal(b, msg); err != nil {
		return nil, err
	}

	return parse.NewResult(msg, []parse.ModifierType{parse.Request})
}
