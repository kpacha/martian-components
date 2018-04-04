package json_schema

import (
	"errors"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

const MIMEJSON = "application/json"

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
