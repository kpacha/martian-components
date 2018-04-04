package json_schema

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/google/martian/parse"
)

var ErrNoJSONRequest = errors.New("request is not a json message")

func init() {
	parse.Register("body.REQUEST-JSON-SCHEMA", RequestVerifierFromJSON)
}

type RequestVerifier struct {
	Verifier
}

// ModifyRequest verifies the body of the response with the given JSON verifier.
func (m *RequestVerifier) ModifyRequest(req *http.Request) error {
	if contentType := req.Header.Get("Content-Type"); contentType != MIMEJSON || req.Body == nil {
		return ErrNoJSONRequest
	}

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	req.Body.Close()
	req.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	return m.Validate(data)
}

func RequestVerifierFromJSON(b []byte) (*parse.Result, error) {
	v, err := VerifierFromJSON(b)
	if err != nil {
		return nil, err
	}

	return parse.NewResult(&RequestVerifier{*v}, []parse.ModifierType{parse.Request})
}
