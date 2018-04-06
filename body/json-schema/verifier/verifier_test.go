package verifier

import (
	"fmt"
	"testing"
)

func Example_verifier() {
	verifier, err := verifierFromJSON(sampleConfig)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(verifier.Validate([]byte(`{"firstName": "foo", "lastName": "bar", "age": 42}`)))
	fmt.Println(verifier.Validate([]byte(`{"firstName": "foo", "lastName": 1, "age": 42}`)))
	fmt.Println(verifier.Validate([]byte(`{"firstName": "foo", "age": 42}`)))
	fmt.Println(verifier.Validate([]byte(`{"firstName": "foo", "lastName": "bar", "age": -42}`)))
	fmt.Println(verifier.Validate([]byte(`{`)))

	// Output:
	// <nil>
	// lastName: Invalid type. Expected: string, given: integer
	// lastName: lastName is required
	// age: Must be greater than or equal to 0
	// unexpected EOF
}

func Benchmark_verifier_ok(b *testing.B) {
	verifier, _ := verifierFromJSON(sampleConfig)
	data := []byte(`{"firstName": "foo", "lastName": "bar", "age": 42}`)
	for i := 0; i < b.N; i++ {
		verifier.Validate(data)
	}
}

func Benchmark_verifier_ko(b *testing.B) {
	verifier, _ := verifierFromJSON(sampleConfig)
	data := []byte(`{"firstName": "foo", "lastName": "bar", "age": -42}`)
	for i := 0; i < b.N; i++ {
		verifier.Validate(data)
	}
}

var sampleConfig = []byte(`{
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
}`)
