package json_schema

import (
	"errors"
	"net/http"

	"github.com/google/martian/parse"
)

var ErrNoJSONResponse = errors.New("response is not a json message")

func init() {
	parse.Register("body.JSON-SCHEMA.Response", ResponseVerifierFromJSON)
}

type ResponseVerifier struct {
	Verifier
}

// ModifyResponse verifies the body of the response with the given JSON verifier.
func (m *ResponseVerifier) ModifyResponse(res *http.Response) error {
	if contentType := res.Header.Get("Content-Type"); contentType != MIMEJSON || res.Body == nil {
		return ErrNoJSONResponse
	}

	data, err := readBody(&(res.Body))
	if err != nil {
		return err
	}

	return m.Validate(data)
}

func ResponseVerifierFromJSON(b []byte) (*parse.Result, error) {
	v, err := VerifierFromJSON(b)
	if err != nil {
		return nil, err
	}

	return parse.NewResult(&ResponseVerifier{*v}, []parse.ModifierType{parse.Response})
}
