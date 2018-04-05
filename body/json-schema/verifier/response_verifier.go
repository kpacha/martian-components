package verifier

import (
	"errors"
	"net/http"
)

var ErrNoJSONResponse = errors.New("response is not a json message")

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

func ResponseVerifierFromJSON(b []byte) (*ResponseVerifier, error) {
	v, err := VerifierFromJSON(b)
	if err != nil {
		return nil, err
	}

	return &ResponseVerifier{*v}, nil
}
