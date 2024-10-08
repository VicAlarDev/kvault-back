// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Victor Alarcon",
            "url": "https://github.com/VicAlarDev/"
        },
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/users": {
            "post": {
                "description": "create a new user account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "Register request",
                        "name": "RegisterRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User created",
                        "schema": {
                            "$ref": "#/definitions/http.userResponse"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Data not found error",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    },
                    "409": {
                        "description": "Data conflict error",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Logs in a registered user and returns an access token if the credentials are valid.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Login and get an access token",
                "parameters": [
                    {
                        "description": "Login request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Succesfully logged in",
                        "schema": {
                            "$ref": "#/definitions/http.authResponse"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/http.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.authResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "v2.local.Gdh5kiOTyyaQ3_bNykYDeYHO21Jg2..."
                }
            }
        },
        "http.errorResponse": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Error message 1",
                        " Error message 2"
                    ]
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "http.userResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "test@example.com"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "username": {
                    "type": "string",
                    "example": "vicalar"
                }
            }
        },
        "request.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "vicalar@gmail.com"
                },
                "password": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 8,
                    "example": "Prueba123"
                }
            }
        },
        "request.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "vicalar@gmail.com"
                },
                "name": {
                    "type": "string",
                    "example": "Victor Alarcon"
                },
                "password": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 8,
                    "example": "Prueba123"
                },
                "username": {
                    "type": "string",
                    "example": "vicalar"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Type \"Bearer\" followed by a space and the access token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{"http", "https"},
	Title:            "KVault API",
	Description:      "Backend API for KVault APP",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
