package json_schema

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

type Verifier struct {
	schema *gojsonschema.Schema
}

func (v *Verifier) Validate(data []byte) error {
	result, err := v.schema.Validate(gojsonschema.NewBytesLoader(data))
	if err != nil {
		return err
	}

	if !result.Valid() {
		errs := make([]string, len(result.Errors()))
		for i, desc := range result.Errors() {
			errs[i] = desc.String()
		}
		return errors.New(strings.Join(errs, "/n"))
	}

	return nil
}

func VerifierFromJSON(b []byte) (*Verifier, error) {
	schema, err := gojsonschema.NewSchema(gojsonschema.NewBytesLoader(b))
	if err != nil {
		return nil, err
	}

	return &Verifier{schema}, nil
}

const MIMEJSON = "application/json"

func readBody(body *io.ReadCloser) ([]byte, error) {
	data, err := ioutil.ReadAll(*body)
	if err != nil {
		return data, err
	}

	(*body).Close()
	*body = ioutil.NopCloser(bytes.NewBuffer(data))
	return data, nil
}
