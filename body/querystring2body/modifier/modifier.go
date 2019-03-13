// Package modifier exposes a request modifier for generating bodies
// from the querystring params
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
	ContentType   string   `json:"content_type"`
}

type Query2BodyModifier struct {
	keysToExtract []string
	template      *template.Template
	method        string
	contentType   string
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

	if m.method != "" {
		req.Method = m.method
	}
	if m.contentType != "" {
		req.Header.Set("Content-Type", m.contentType)
	} else {
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}
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

	tmpl, err := template.New("query2body_modifier").Parse(cfg.Template)
	if err != nil {
		return nil, err
	}

	return &Query2BodyModifier{
		keysToExtract: cfg.KeysToExtract,
		template:      tmpl,
		method:        cfg.Method,
	}, nil
}
