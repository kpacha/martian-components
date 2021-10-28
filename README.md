# martian-components

A collection of martian modifiers, filters and validators

## Using it as a lib

Import the required packages into your project and start using them as any other martian component

	import(
		_ "github.com/kpacha/martian-components/body/elastic-search"
		_ "github.com/kpacha/martian-components/body/json-schema"
	)

## Using the modules as KrakenD plugins

**THIS IS A DEPRECATED EXPERIMENT. KRAKEND DOES NOT ACCEPT MARTIAN PLUGINS**

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
