// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "s-akhmedoff",
            "url": "https://github.com/s-akhmedoff",
            "email": "sodikjon.akhmedoff@gmail.com"
        },
        "license": {
            "name": "GNU GPLv3",
            "url": "https://www.gnu.org/licenses/gpl-3.0.en.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/config": {
            "get": {
                "description": "expose all configs for debugging",
                "produces": [
                    "application/json"
                ],
                "summary": "Show config",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/views.R"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/config.Config"
                                        }
                                    }
                                }
                            ]
                        },
                        "headers": {
                            "environment": {
                                "type": "string",
                                "description": "Current environment"
                            },
                            "go-os": {
                                "type": "string",
                                "description": "Go OS"
                            },
                            "go-version": {
                                "type": "string",
                                "description": "Version of Golang"
                            }
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        },
                        "headers": {
                            "environment": {
                                "type": "string",
                                "description": "Current environment"
                            },
                            "go-os": {
                                "type": "string",
                                "description": "Go OS"
                            },
                            "go-version": {
                                "type": "string",
                                "description": "Version of Golang"
                            }
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get all products",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update existing product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "product's ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "product params",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.productParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    }
                }
            },
            "post": {
                "description": "create new product in storage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create new product",
                "parameters": [
                    {
                        "description": "product params",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.productParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Read Product By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "product's id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "Delete product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "product's ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    }
                }
            }
        },
        "/products/{ski}": {
            "get": {
                "description": "get product by its ski",
                "produces": [
                    "application/json"
                ],
                "summary": "Read Product By SKI",
                "parameters": [
                    {
                        "type": "string",
                        "description": "product's ski",
                        "name": "ski",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    }
                }
            }
        },
        "/products/{type}": {
            "post": {
                "description": "get product by its type",
                "produces": [
                    "application/json"
                ],
                "summary": "Read product by type",
                "parameters": [
                    {
                        "type": "string",
                        "description": "product's type",
                        "name": "type",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/views.R"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "config.Config": {
            "type": "object",
            "properties": {
                "environment": {
                    "type": "string"
                },
                "httphost": {
                    "type": "string"
                },
                "httpport": {
                    "type": "string"
                },
                "postgresDatabase": {
                    "type": "string"
                },
                "postgresHost": {
                    "type": "string"
                },
                "postgresPassword": {
                    "type": "string"
                },
                "postgresPort": {
                    "type": "string"
                },
                "postgresUser": {
                    "type": "string"
                }
            }
        },
        "rest.productParams": {
            "type": "object",
            "required": [
                "description",
                "is_active",
                "name",
                "product_type",
                "ski",
                "uri"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "product_type": {
                    "type": "string"
                },
                "ski": {
                    "type": "string"
                },
                "uri": {
                    "type": "string"
                }
            }
        },
        "views.R": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "error_code": {
                    "type": "integer",
                    "example": 0
                },
                "error_note": {
                    "type": "string",
                    "example": "Error Note"
                },
                "status": {
                    "type": "string",
                    "example": "Success | Failure"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:7070",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Simple Service REST API",
	Description: "This is a sample server",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
