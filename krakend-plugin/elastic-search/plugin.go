package main

import (
	"github.com/devopsfaith/krakend-martian/register"
	"github.com/kpacha/martian-components/body/elastic-search/modifier"
)

func init() {
	register.Set("body.ESQuery", []register.Scope{register.ScopeRequest}, func(b []byte) (interface{}, error) {
		return modifier.FromJSON(b)
	})
}

func main() {

}
