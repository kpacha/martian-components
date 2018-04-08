// Package elastic registers a request modifier for generating parametrized queries
// to an elastic search service
package elastic

import (
	"github.com/google/martian/parse"
	"github.com/kpacha/martian-components/body/elastic-search/modifier"
)

func init() {
	parse.Register("body.ESQuery", FromJSON)
}

func FromJSON(b []byte) (*parse.Result, error) {
	msg, err := modifier.FromJSON(b)
	if err != nil {
		return nil, err
	}

	return parse.NewResult(msg, []parse.ModifierType{parse.Request})
}
