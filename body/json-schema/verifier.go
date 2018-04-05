package json_schema

import (
	"github.com/google/martian/parse"
	"github.com/kpacha/martian-components/body/json-schema/verifier"
)

func init() {
	parse.Register("body.JSON-SCHEMA.Request", RequestVerifierFromJSON)
	parse.Register("body.JSON-SCHEMA.Response", ResponseVerifierFromJSON)
}

func RequestVerifierFromJSON(b []byte) (*parse.Result, error) {
	v, err := verifier.RequestVerifierFromJSON(b)
	if err != nil {
		return nil, err
	}

	return parse.NewResult(v, []parse.ModifierType{parse.Request})
}

func ResponseVerifierFromJSON(b []byte) (*parse.Result, error) {
	v, err := verifier.ResponseVerifierFromJSON(b)
	if err != nil {
		return nil, err
	}

	return parse.NewResult(v, []parse.ModifierType{parse.Response})
}
