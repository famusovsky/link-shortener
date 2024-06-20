// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/": {
            "post": {
                "description": "Add link to db and get its key.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "links"
                ],
                "summary": "Adds input link to the DB.",
                "parameters": [
                    {
                        "description": "Input link",
                        "name": "link",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.input"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Key",
                        "schema": {
                            "$ref": "#/definitions/app.output"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/app.outErr"
                        }
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/app.outErr"
                        }
                    }
                }
            }
        },
        "/{key}": {
            "get": {
                "description": "Get link added to db by the key",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "links"
                ],
                "summary": "Returns link by the shortened.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Key for the link",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Redirects to link",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "Location": {
                                "type": "string",
                                "description": "Url to redirect"
                            }
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/app.outErr"
                        }
                    },
                    "404": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/app.outErr"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.input": {
            "type": "object",
            "properties": {
                "link": {
                    "type": "string"
                }
            }
        },
        "app.outErr": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "app.output": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Link Shortener Api",
	Description:      "This is an api for shortening apis.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
