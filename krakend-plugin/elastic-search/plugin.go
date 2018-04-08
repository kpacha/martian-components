package main

import (
	"fmt"

	"github.com/devopsfaith/krakend-martian/register"
	"github.com/kpacha/martian-components/body/elastic-search/modifier"
)

var Registrable registrable

type registrable int

func (r *registrable) RegisterExternal(setter func(namespace, name string, v interface{})) error {
	fmt.Println("registrable", r, "from plugin 'krakend_martian-ES' is registering its components depending on external modules at", register.Namespace)

	component := register.NewComponent([]string{register.ScopeRequest},
		func(b []byte) (interface{}, error) {
			return modifier.FromJSON(b)
		})
	setter(register.Namespace, "body.ESQuery", component)

	return nil
}

func init() {
	fmt.Println("plugin 'krakend_martian-ES' loaded!")
}

func main() {

}
