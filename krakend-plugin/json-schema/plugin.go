package main

import (
	"fmt"

	"github.com/devopsfaith/krakend-martian/register"
	"github.com/kpacha/martian-components/body/json-schema/verifier"
)

var Registrable registrable

type registrable int

func (r *registrable) RegisterExternal(setter func(namespace, name string, v interface{})) error {
	fmt.Println("registrable", r, "from plugin 'krakend_martian-JSON' is registering its components depending on external modules at", register.Namespace)

	requestComponent := register.NewComponent([]string{register.ScopeRequest},
		func(b []byte) (interface{}, error) {
			return verifier.RequestVerifierFromJSON(b)
		})
	setter(register.Namespace, "body.JSON-SCHEMA.Request", requestComponent)

	responseComponent := register.NewComponent([]string{register.ScopeResponse},
		func(b []byte) (interface{}, error) {
			return verifier.ResponseVerifierFromJSON(b)
		})
	setter(register.Namespace, "body.JSON-SCHEMA.Response", responseComponent)

	return nil
}

func init() {
	fmt.Println("plugin 'krakend_martian-JSON' loaded!")
}

func main() {

}
