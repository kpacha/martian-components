package json_schema

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ExampleRequestVerifier() {
	cfg := `{
    "title": "Person",
    "type": "object",
    "properties": {
        "firstName": {
            "type": "string"
        },
        "lastName": {
            "type": "string"
        },
        "age": {
            "description": "Age in years",
            "type": "integer",
            "minimum": 0
        }
    },
    "required": ["firstName", "lastName"]
}`
	verifier, err := RequestVerifierFromJSON([]byte(cfg))
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(verifier.RequestModifier().ModifyRequest(newRequest(`{"firstName": "foo", "lastName": "bar", "age": 42}`)))
	fmt.Println(verifier.RequestModifier().ModifyRequest(newRequest(`{"firstName": "foo", "lastName": 1, "age": 42}`)))
	fmt.Println(verifier.RequestModifier().ModifyRequest(newRequest(`{"firstName": "foo", "age": 42}`)))
	fmt.Println(verifier.RequestModifier().ModifyRequest(newRequest(`{"firstName": "foo", "lastName": "bar", "age": -42}`)))
	fmt.Println(verifier.RequestModifier().ModifyRequest(newRequest("{")))
	fmt.Println(verifier.RequestModifier().ModifyRequest(&http.Request{Body: ioutil.NopCloser(bytes.NewBufferString("{"))}))
	fmt.Println(verifier.RequestModifier().ModifyRequest(&http.Request{Header: http.Header{"Content-Type": []string{MIMEJSON}}}))

	// Output:
	// <nil>
	// lastName: Invalid type. Expected: string, given: integer
	// lastName: lastName is required
	// age: Must be greater than or equal to 0
	// unexpected EOF
	// request is not a json message
	// request is not a json message
}

func newRequest(body string) *http.Request {
	req, _ := http.NewRequest("POST", "/", ioutil.NopCloser(bytes.NewBufferString(body)))
	req.Header.Set("Content-Type", MIMEJSON)
	return req
}
