package json_schema

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func ExampleResponseVerifier() {
	verifier, err := ResponseVerifierFromJSON(sampleConfig)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(verifier.ResponseModifier().ModifyResponse(newResponse(`{"firstName": "foo", "lastName": "bar", "age": 42}`)))
	fmt.Println(verifier.ResponseModifier().ModifyResponse(newResponse(`{"firstName": "foo", "lastName": 1, "age": 42}`)))
	fmt.Println(verifier.ResponseModifier().ModifyResponse(newResponse(`{"firstName": "foo", "age": 42}`)))
	fmt.Println(verifier.ResponseModifier().ModifyResponse(newResponse(`{"firstName": "foo", "lastName": "bar", "age": -42}`)))
	fmt.Println(verifier.ResponseModifier().ModifyResponse(newResponse(`{`)))
	fmt.Println(verifier.ResponseModifier().ModifyResponse(&http.Response{Body: ioutil.NopCloser(bytes.NewBufferString("{"))}))
	fmt.Println(verifier.ResponseModifier().ModifyResponse(&http.Response{Header: http.Header{"Content-Type": []string{MIMEJSON}}}))

	// Output:
	// <nil>
	// lastName: Invalid type. Expected: string, given: integer
	// lastName: lastName is required
	// age: Must be greater than or equal to 0
	// unexpected EOF
	// response is not a json message
	// response is not a json message
}

func newResponse(body string) *http.Response {
	return &http.Response{
		Body:   ioutil.NopCloser(bytes.NewBufferString(body)),
		Header: http.Header{"Content-Type": []string{MIMEJSON}},
	}
}

func BenchmarkResponseVerifier_ModifyResponse_ok(b *testing.B) {
	verifier, _ := ResponseVerifierFromJSON(sampleConfig)
	rm := verifier.ResponseModifier()
	res := newResponse(`{"firstName": "foo", "lastName": "bar", "age": 42}`)
	for i := 0; i < b.N; i++ {
		rm.ModifyResponse(res)
	}
}

func BenchmarkResponseVerifier_ModifyResponse_ko(b *testing.B) {
	verifier, _ := ResponseVerifierFromJSON(sampleConfig)
	rm := verifier.ResponseModifier()
	res := newResponse(`{"firstName": "foo", "lastName": "bar", "age": -42}`)
	for i := 0; i < b.N; i++ {
		rm.ModifyResponse(res)
	}
}
