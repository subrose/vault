// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "fiber@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/collections": {
            "get": {
                "description": "Returns all Collections",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "collections"
                ],
                "summary": "Get all Collections",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/vault.Collection"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a Collection",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "collections"
                ],
                "summary": "Create a Collection",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/vault.Collection"
                        }
                    }
                }
            }
        },
        "/collections/{name}": {
            "get": {
                "description": "Returns a Collection given a name",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "collections"
                ],
                "summary": "Get a Collection by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Collection Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vault.Collection"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a Collection given a name",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "collections"
                ],
                "summary": "Delete a Collection by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Collection Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/collections/{name}/records": {
            "get": {
                "description": "Returns all Records",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "records"
                ],
                "summary": "Get all Records",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Collection Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/vault.Record"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a Record",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "records"
                ],
                "summary": "Create a Record",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Collection Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/collections/{name}/records/search": {
            "post": {
                "description": "Searches for Records",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "records"
                ],
                "summary": "Search Records",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Collection Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Search filters",
                        "name": "filters",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/vault.Record"
                            }
                        }
                    }
                }
            }
        },
        "/collections/{name}/records/{id}": {
            "get": {
                "description": "Returns a Record given an id",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "records"
                ],
                "summary": "Get a Record by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Collection Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Record Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Record formats",
                        "name": "formats",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vault.Record"
                        }
                    }
                }
            },
            "put": {
                "description": "Updates a Record",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "records"
                ],
                "summary": "Update a Record",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Collection Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Record Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a Record",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "records"
                ],
                "summary": "Delete a Record",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Collection Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Record Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/policies": {
            "get": {
                "description": "Returns all Policies",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "policies"
                ],
                "summary": "Get all Policies",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/vault.Policy"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a Policy",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "policies"
                ],
                "summary": "Create a Policy",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/vault.Policy"
                        }
                    }
                }
            }
        },
        "/policies/{policyId}": {
            "get": {
                "description": "Returns a Policy given an id",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "policies"
                ],
                "summary": "Get a Policy by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Policy Id",
                        "name": "policyId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vault.Policy"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a Policy given an id",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "policies"
                ],
                "summary": "Delete a Policy by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Policy Id",
                        "name": "policyId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/principals": {
            "post": {
                "description": "Creates a Principal",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "principals"
                ],
                "summary": "Create a Principal",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.PrincipalResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a Principal given an id",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "principals"
                ],
                "summary": "Delete a Principal by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/principals/{username}": {
            "get": {
                "description": "Returns a Principal given an id",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "principals"
                ],
                "summary": "Get a Prinicipal by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.PrincipalResponse"
                        }
                    }
                }
            }
        },
        "/tokens": {
            "post": {
                "description": "Creates a Token",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tokens"
                ],
                "summary": "Create a Token",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tokens/{tokenId}": {
            "get": {
                "description": "Returns a Token given an id",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tokens"
                ],
                "summary": "Get a Token by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token Id",
                        "name": "tokenId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.PrincipalResponse": {
            "type": "object",
            "required": [
                "username"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "policies": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 3
                }
            }
        },
        "vault.Collection": {
            "type": "object",
            "required": [
                "fields",
                "name"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "fields": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/vault.Field"
                    }
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 3
                },
                "parent": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 3
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "vault.Field": {
            "type": "object",
            "required": [
                "type"
            ],
            "properties": {
                "is_indexed": {
                    "type": "boolean"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "vault.Policy": {
            "type": "object",
            "required": [
                "actions",
                "effect",
                "resources"
            ],
            "properties": {
                "actions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/vault.PolicyAction"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "effect": {
                    "enum": [
                        "allow",
                        "deny"
                    ],
                    "allOf": [
                        {
                            "$ref": "#/definitions/vault.PolicyEffect"
                        }
                    ]
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "resources": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "vault.PolicyAction": {
            "type": "string",
            "enum": [
                "read",
                "write"
            ],
            "x-enum-varnames": [
                "PolicyActionRead",
                "PolicyActionWrite"
            ]
        },
        "vault.PolicyEffect": {
            "type": "string",
            "enum": [
                "deny",
                "allow"
            ],
            "x-enum-varnames": [
                "EffectDeny",
                "EffectAllow"
            ]
        },
        "vault.Record": {
            "type": "object",
            "additionalProperties": {
                "type": "string"
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3001",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Fiber Example API",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}