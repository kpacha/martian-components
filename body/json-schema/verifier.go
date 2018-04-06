// Package json_schema registers a couple of body request and response modifiers
// for validating the json messages using json-schema (http://json-schema.org/)
package json_schema

import (
	"github.com/google/martian/parse"
	"github.com/kpacha/martian-components/body/json-schema/verifier"
)

func init() {
	parse.Register("body.JSON-SCHEMA.Request", RequestVerifierFromJSON)
	parse.Register("body.JSON-SCHEMA.Response", ResponseVerifierFromJSON)
}

// RequestVerifierFromJSON creates a new RequestVerifier with the received specs
// Sample DSL:
// 	{
// 	    "body.JSON-SCHEMA.Request": {
// 	        "title": "Person",
// 	        "type": "object",
// 	        "properties": {
// 	            "firstName": {
// 	                "type": "string"
// 	            },
// 	            "lastName": {
// 	                "type": "string"
// 	            },
// 	            "age": {
// 	                "description": "Age in years",
// 	                "type": "integer",
// 	                "minimum": 0
// 	            }
// 	        },
// 	        "required": ["firstName", "lastName"]
// 	    }
// 	}
func RequestVerifierFromJSON(b []byte) (*parse.Result, error) {
	v, err := verifier.RequestVerifierFromJSON(b)
	if err != nil {
		return nil, err
	}

	return parse.NewResult(v, []parse.ModifierType{parse.Request})
}

// ResponseVerifierFromJSON creates a new ResponseVerifier with the received specs
// Sample DSL:
// 	{
// 	    "body.JSON-SCHEMA.Response": {
// 	        "title": "Person",
// 	        "type": "object",
// 	        "properties": {
// 	            "firstName": {
// 	                "type": "string"
// 	            },
// 	            "lastName": {
// 	                "type": "string"
// 	            },
// 	            "age": {
// 	                "description": "Age in years",
// 	                "type": "integer",
// 	                "minimum": 0
// 	            }
// 	        },
// 	        "required": ["firstName", "lastName"]
// 	    }
// 	}
func ResponseVerifierFromJSON(b []byte) (*parse.Result, error) {
	v, err := verifier.ResponseVerifierFromJSON(b)
	if err != nil {
		return nil, err
	}

	return parse.NewResult(v, []parse.ModifierType{parse.Response})
}
