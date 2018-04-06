package verifier

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoJSONRequest = errors.New("request is not a json message")

type RequestVerifier struct {
	verifier
}

// ModifyRequest verifies the body of the response with the given JSON verifier.
func (m *RequestVerifier) ModifyRequest(req *http.Request) error {
	if !strings.HasPrefix(req.Header.Get("Content-Type"), MIMEJSON) || req.Body == nil {
		return ErrNoJSONRequest
	}

	data, err := readBody(&(req.Body))
	if err != nil {
		return err
	}

	return m.Validate(data)
}

func RequestVerifierFromJSON(b []byte) (*RequestVerifier, error) {
	v, err := verifierFromJSON(b)
	if err != nil {
		return nil, err
	}

	return &RequestVerifier{*v}, nil
}
