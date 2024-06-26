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
            "name": "Duy Nguyen",
            "url": "https://github.com/vnworkday/gateway",
            "email": "ntduy.cs@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/tenants": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "List all tenants",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tenant"
                ],
                "summary": "List all tenants",
                "operationId": "ListTenants",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/account.ListTenantsResponse"
                        }
                    },
                    "401": {
                        "description": "when the user is not authenticated",
                        "schema": {
                            "$ref": "#/definitions/shared.Error"
                        }
                    },
                    "403": {
                        "description": "when the user is not authorized",
                        "schema": {
                            "$ref": "#/definitions/shared.Error"
                        }
                    },
                    "500": {
                        "description": "when the server is unable to handle the request",
                        "schema": {
                            "$ref": "#/definitions/shared.Error"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Create a tenant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tenant"
                ],
                "summary": "Create a tenant",
                "operationId": "CreateTenant",
                "parameters": [
                    {
                        "description": "Tenant object that needs to be created",
                        "name": "tenant",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/account.Tenant"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/account.Tenant"
                        }
                    },
                    "401": {
                        "description": "when the user is not authenticated",
                        "schema": {
                            "$ref": "#/definitions/shared.Error"
                        }
                    },
                    "403": {
                        "description": "when the user is not authorized",
                        "schema": {
                            "$ref": "#/definitions/shared.Error"
                        }
                    },
                    "500": {
                        "description": "when the server is unable to handle the request",
                        "schema": {
                            "$ref": "#/definitions/shared.Error"
                        }
                    }
                }
            }
        },
        "/tenants/{id}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    },
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Get a tenant by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tenant"
                ],
                "summary": "Get a tenant",
                "operationId": "GetTenant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tenant ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/account.Tenant"
                        }
                    },
                    "400": {
                        "description": "when the provided id is invalid",
                        "schema": {
                            "$ref": "#/definitions/shared.Error"
                        }
                    },
                    "401": {
                        "description": "when the user is not authenticated",
                        "schema": {
                            "$ref": "#/definitions/shared.Error"
                        }
                    },
                    "403": {
                        "description": "when the user is not authorized",
                        "schema": {
                            "$ref": "#/definitions/shared.Error"
                        }
                    },
                    "404": {
                        "description": "when the tenant is not found",
                        "schema": {
                            "$ref": "#/definitions/shared.Error"
                        }
                    },
                    "500": {
                        "description": "when the server is unable to handle the request",
                        "schema": {
                            "$ref": "#/definitions/shared.Error"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Update an existing tenant by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tenant"
                ],
                "summary": "Update a tenant",
                "operationId": "UpdateTenant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tenant ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Tenant object that needs to be updated",
                        "name": "tenant",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/account.Tenant"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/account.Tenant"
                        }
                    },
                    "401": {
                        "description": "when the user is not authenticated",
                        "schema": {
                            "$ref": "#/definitions/shared.Error"
                        }
                    },
                    "403": {
                        "description": "when the user is not authorized",
                        "schema": {
                            "$ref": "#/definitions/shared.Error"
                        }
                    },
                    "500": {
                        "description": "when the server is unable to handle the request",
                        "schema": {
                            "$ref": "#/definitions/shared.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "account.ListTenantsResponse": {
            "description": "represents a list of tenants with pagination information in the response.",
            "type": "object",
            "required": [
                "items",
                "total",
                "totalPages"
            ],
            "properties": {
                "items": {
                    "description": "Items is a list of tenants",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/account.Tenant"
                    }
                },
                "nextToken": {
                    "description": "NextToken is a token that can be used to retrieve the next page of results.\nIf this field is not present, it means that there are no more results to retrieve.",
                    "type": "string",
                    "format": "base64",
                    "example": "Vk4gV29ya2RheSBFeGFtcGxl"
                },
                "previousToken": {
                    "description": "PreviousToken a token that can be used to retrieve the previous page of results.\nIf this field is not present, it means that there are no more results to retrieve.",
                    "type": "string",
                    "format": "base64",
                    "example": "Vk4gV29ya2RheSBFeGFtcGxl"
                },
                "total": {
                    "description": "Total is the total number of items in the list.",
                    "type": "integer",
                    "minimum": 0,
                    "example": 100
                },
                "totalPages": {
                    "description": "TotalPages is the total number of pages in the list.",
                    "type": "integer",
                    "minimum": 0,
                    "example": 10
                }
            }
        },
        "account.Tenant": {
            "description": "represents a tenant.",
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "id": {
                    "description": "ID of the tenant",
                    "type": "string",
                    "example": "abcxyz"
                },
                "name": {
                    "description": "Name of the tenant",
                    "type": "string",
                    "example": "Tenant Name"
                }
            }
        },
        "shared.Code": {
            "type": "integer",
            "enum": [
                1000,
                1001,
                1002,
                1003,
                2000,
                2001,
                2002,
                2003,
                2004
            ],
            "x-enum-comments": {
                "CodeErrNotFound": "retryable",
                "CodeErrTimeout": "retryable",
                "CodeErrTooManyRequests": "retryable",
                "CodeErrUnavailable": "retryable"
            },
            "x-enum-varnames": [
                "CodeErrInternal",
                "CodeErrTimeout",
                "CodeErrTooManyRequests",
                "CodeErrUnavailable",
                "CodeErrValidation",
                "CodeErrNotFound",
                "CodeErrUnauthorized",
                "CodeErrForbidden",
                "CodeErrTooLarge"
            ]
        },
        "shared.Error": {
            "description": "Represents a client-facing API error.",
            "type": "object",
            "required": [
                "code"
            ],
            "properties": {
                "code": {
                    "description": "Code should be unique and identifiable. It is used to determine the type of error.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/shared.Code"
                        }
                    ],
                    "example": 1000
                },
                "message": {
                    "description": "Message is a human-readable description of the error. Should not be used to display to the user.",
                    "type": "string",
                    "example": "An unexpected error occurred. Please try again later."
                },
                "title": {
                    "description": "Title is a short, human-readable title of the error. It is used to determine the type of error.",
                    "type": "string",
                    "example": "Internal Server Error"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKey": {
            "description": "Please provide a valid API key in the header.",
            "type": "apiKey",
            "name": "x-api-key",
            "in": "header"
        },
        "JWT": {
            "description": "Please provide a valid JWT token in the header.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1.0",
	Host:             "http://localhost:3000",
	BasePath:         "/api/v1",
	Schemes:          []string{"http", "https"},
	Title:            "VN Workday Gateway API",
	Description:      "This is the API documentation for VN Workday Gateway API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
