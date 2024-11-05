// Package doc GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package doc

import "github.com/swaggo/swag"

const docTemplateAPI = `{
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
        "/api/users": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.UserResponse"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.UserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfoAPI holds exported Swagger Info so clients can modify it
var SwaggerInfoAPI = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "API",
	SwaggerTemplate:  docTemplateAPI,
}

func init() {
	swag.Register(SwaggerInfoAPI.InstanceName(), SwaggerInfoAPI)
}
