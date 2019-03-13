// Package modifier exposes a request modifier for generating parametrized queries
// to an elastic search service
package modifier

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"
)

type Config struct {
	KeysToExtract []string `json:"keys_to_extract"`
	Template      string   `json:"template"`
	Method        string   `json:"method"`
}

type Query2BodyModifier struct {
	keysToExtract []string
	template      *template.Template
	method        string
}

func (m *Query2BodyModifier) ModifyRequest(req *http.Request) error {
	query := req.URL.Query()

	buf := new(bytes.Buffer)
	if err := m.template.Execute(buf, query); err != nil {
		return err
	}

	for _, k := range m.keysToExtract {
		query.Del(k)
	}

	req.Method = m.method
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.ContentLength = int64(buf.Len())
	req.Body = ioutil.NopCloser(buf)
	req.URL.RawQuery = query.Encode()

	return nil
}

func FromJSON(b []byte) (*Query2BodyModifier, error) {
	cfg := &Config{}
	if err := json.Unmarshal(b, cfg); err != nil {
		return nil, err
	}

	tmpl, err := template.New("query2jsonbody_modifier").Parse(cfg.Template)
	if err != nil {
		return nil, err
	}
	if cfg.Method == "" {
		cfg.Method = http.MethodPost
	}

	return &Query2BodyModifier{
		keysToExtract: cfg.KeysToExtract,
		template:      tmpl,
		method:        cfg.Method,
	}, nil
}
