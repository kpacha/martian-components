package json_schema

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func ExampleRequestVerifier() {
	verifier, err := RequestVerifierFromJSON(sampleConfig)
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

func BenchmarkRequestVerifier_ModifyRequest_ok(b *testing.B) {
	verifier, _ := RequestVerifierFromJSON(sampleConfig)
	rm := verifier.RequestModifier()
	req := newRequest(`{"firstName": "foo", "lastName": "bar", "age": 42}`)
	for i := 0; i < b.N; i++ {
		rm.ModifyRequest(req)
	}
}

func BenchmarkRequestVerifier_ModifyRequest_ko(b *testing.B) {
	verifier, _ := RequestVerifierFromJSON(sampleConfig)
	rm := verifier.RequestModifier()
	req := newRequest(`{"firstName": "foo", "lastName": "bar", "age": -42}`)
	for i := 0; i < b.N; i++ {
		rm.ModifyRequest(req)
	}
}
