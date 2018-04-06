package main

import (
	"github.com/devopsfaith/krakend-martian/register"
	"github.com/kpacha/martian-components/body/json-schema/verifier"
)

func init() {
	register.Set("body.JSON-SCHEMA.Request", []register.Scope{register.ScopeRequest}, func(b []byte) (interface{}, error) {
		return verifier.RequestVerifierFromJSON(b)
	})
	register.Set("body.JSON-SCHEMA.Response", []register.Scope{register.ScopeResponse}, func(b []byte) (interface{}, error) {
		return verifier.ResponseVerifierFromJSON(b)
	})
}

func main() {

}
