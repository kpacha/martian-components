// Package querystring2body registers a request modifier for generating JSON encoded bodies
// from the querystring params
package querystring2body

import (
	"github.com/google/martian/parse"
	"github.com/kpacha/martian-components/body/querystring2body/modifier"
)

func init() {
	parse.Register("body.FromQuerystring", FromJSON)
}

func FromJSON(b []byte) (*parse.Result, error) {
	msg, err := modifier.FromJSON(b)
	if err != nil {
		return nil, err
	}

	return parse.NewResult(msg, []parse.ModifierType{parse.Request})
}
