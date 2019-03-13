// Package querystring2jsonbody registers a request modifier for generating parametrized queries
// to an querystring2jsonbody search service
package querystring2jsonbody

import (
	"github.com/google/martian/parse"
	"github.com/kpacha/martian-components/body/querystring2jsonbody/modifier"
)

func init() {
	parse.Register("body.JSONFromQuerystring", FromJSON)
}

func FromJSON(b []byte) (*parse.Result, error) {
	msg, err := modifier.FromJSON(b)
	if err != nil {
		return nil, err
	}

	return parse.NewResult(msg, []parse.ModifierType{parse.Request})
}
