# martian-components

A collection of martian modifiers, filters and validators

## Using it as a lib

Import the required packages into your project and start using them as any other martian component

	import(
		_ "github.com/kpacha/martian-components/body/elastic-search"
		_ "github.com/kpacha/martian-components/body/json-schema"
	)

## Using the modules as KrakenD plugins

Compile the desired package with the plugin flag

	$ go build -buildmode=plugin -o krakend-martian_json-schema.so ./krakend-plugin/json-schema
	$ go build -buildmode=plugin -o krakend-martian_es.so ./krakend-plugin/elastic-search

And place the plugins into your `plugin` folder, so the KrakenD can load them in runtime.

## List of components:

### ElasticSearch query modifier

Sample DSL:

	{
	    "body.ESQuery": {
	        "keys" : ["foo", "bar", "x"]
	    }
	}


### JSON-Schema request and response verifiers

Sample DSL:

	{
	    "body.JSON-SCHEMA.Request": {
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
	    }
	}

Check the [json-schema](http://json-schema.org/) site for more about schema definitions