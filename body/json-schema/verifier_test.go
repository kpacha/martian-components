package json_schema

import "fmt"

func ExampleVerifier() {
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
	verifier, err := VerifierFromJSON([]byte(cfg))
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
