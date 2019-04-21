// Package modifier exposes a request modifier for setting a custom Host header
package modifier

import (
	"encoding/json"
	"net/http"
)

type Config struct {
	Host string `json:"host"`
}

type HostModifier string

func (m *HostModifier) ModifyRequest(req *http.Request) error {
	req.Host = string(*m)

	return nil
}

func FromJSON(b []byte) (*HostModifier, error) {
	cfg := &Config{}
	if err := json.Unmarshal(b, cfg); err != nil {
		return nil, err
	}

	m := HostModifier(cfg.Host)
	return &m, nil
}
