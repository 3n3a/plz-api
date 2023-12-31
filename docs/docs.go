// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/plz/all": {
            "get": {
                "description": "Get all Postal Codes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "postal-codes"
                ],
                "summary": "Get all Postal Codes",
                "responses": {}
            }
        },
        "/api/plz/search": {
            "get": {
                "description": "Returns a list of postal codes based on the query.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "postal-codes"
                ],
                "summary": "Finds Postal Codes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "plz search by place",
                        "name": "place",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "plz search by code",
                        "name": "code",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Postal Codes API",
	Description:      "Fast and modern REST-API for Swiss Postal Codes",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
