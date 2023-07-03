// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/api/createOrder": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "CreateOrder",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "CreateOrder",
                "parameters": [
                    {
                        "type": "string",
                        "description": "productId",
                        "name": "productId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/createProduct": {
            "post": {
                "description": "CreateProduct",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "CreateProduct",
                "parameters": [
                    {
                        "description": "name,desc,type is require",
                        "name": "Product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/product.Product"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/deleteOrder": {
            "delete": {
                "description": "DeleteOrder",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "DeleteOrder",
                "parameters": [
                    {
                        "type": "string",
                        "description": "productId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/getOrder": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "GetOrder",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "GetOrder",
                "parameters": [
                    {
                        "type": "string",
                        "description": "optional",
                        "name": "orderId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "optional",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "optional",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "optional",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Bearer",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/order.Order"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/getProduct": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "GetProduct",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "GetProduct",
                "parameters": [
                    {
                        "type": "string",
                        "description": "optional",
                        "name": "productId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "optional",
                        "name": "productType",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "optional",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "optional",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Bearer",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/product.Product"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/getUser": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "GetUser",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "GetUser",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/register": {
            "post": {
                "description": "CreateUser",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "CreateUser",
                "parameters": [
                    {
                        "description": "except created every field are require and validate length",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authen"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "require",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "require",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authen.LoginRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/validatePinCode": {
            "get": {
                "description": "logic test - ValidatePinCode",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "logictest"
                ],
                "summary": "ValidatePinCode",
                "parameters": [
                    {
                        "type": "string",
                        "description": "input",
                        "name": "input",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/logictest.Result"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/logictest.Result"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "authen.LoginRes": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "logictest.Result": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "string"
                }
            }
        },
        "order.Order": {
            "type": "object",
            "properties": {
                "createdDate": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "productDesc": {
                    "type": "string"
                },
                "productId": {
                    "type": "string"
                },
                "productName": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "product.Product": {
            "type": "object",
            "required": [
                "productName",
                "productType"
            ],
            "properties": {
                "createdDate": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "productDesc": {
                    "type": "string",
                    "maxLength": 1000
                },
                "productId": {
                    "type": "string"
                },
                "productName": {
                    "type": "string",
                    "maxLength": 255
                },
                "productType": {
                    "type": "string",
                    "maxLength": 50
                }
            }
        },
        "user.User": {
            "type": "object",
            "required": [
                "firstname",
                "lastname",
                "password",
                "rePassword",
                "username"
            ],
            "properties": {
                "createdDate": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string",
                    "maxLength": 50
                },
                "id": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string",
                    "maxLength": 50
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 6
                },
                "rePassword": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 6
                },
                "username": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 6
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{"http"},
	Title:            "NeversitupTest",
	Description:      "This is a sample server server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
