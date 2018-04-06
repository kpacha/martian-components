package verifier

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoJSONResponse = errors.New("response is not a json message")

type ResponseVerifier struct {
	verifier
}

// ModifyResponse verifies the body of the response with the given JSON verifier.
func (m *ResponseVerifier) ModifyResponse(res *http.Response) error {
	if !strings.HasPrefix(res.Header.Get("Content-Type"), MIMEJSON) || res.Body == nil {
		return ErrNoJSONResponse
	}

	data, err := readBody(&(res.Body))
	if err != nil {
		return err
	}

	return m.Validate(data)
}

func ResponseVerifierFromJSON(b []byte) (*ResponseVerifier, error) {
	v, err := verifierFromJSON(b)
	if err != nil {
		return nil, err
	}

	return &ResponseVerifier{*v}, nil
}
