package json_schema

import (
	"errors"
	"net/http"

	"github.com/google/martian/parse"
)

var ErrNoJSONRequest = errors.New("request is not a json message")

func init() {
	parse.Register("body.JSON-SCHEMA.Request", RequestVerifierFromJSON)
}

type RequestVerifier struct {
	Verifier
}

// ModifyRequest verifies the body of the response with the given JSON verifier.
func (m *RequestVerifier) ModifyRequest(req *http.Request) error {
	if contentType := req.Header.Get("Content-Type"); contentType != MIMEJSON || req.Body == nil {
		return ErrNoJSONRequest
	}

	data, err := readBody(&(req.Body))
	if err != nil {
		return err
	}

	return m.Validate(data)
}

func RequestVerifierFromJSON(b []byte) (*parse.Result, error) {
	v, err := VerifierFromJSON(b)
	if err != nil {
		return nil, err
	}

	return parse.NewResult(&RequestVerifier{*v}, []parse.ModifierType{parse.Request})
}
