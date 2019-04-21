// Package host registers a request modifier for setting a custom Host header
package host

import (
	"github.com/google/martian/parse"
	"github.com/kpacha/martian-components/header/host/modifier"
)

func init() {
	parse.Register("header.Host", FromJSON)
}

func FromJSON(b []byte) (*parse.Result, error) {
	msg, err := modifier.FromJSON(b)
	if err != nil {
		return nil, err
	}

	return parse.NewResult(msg, []parse.ModifierType{parse.Request})
}
