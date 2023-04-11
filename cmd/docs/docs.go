// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "sethukumarj.com",
            "email": "sethukumarj.76@gmail.com"
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
        "/account/create": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account-Service"
                ],
                "summary": "Create account",
                "operationId": "Create Account",
                "parameters": [
                    {
                        "description": "Create Account",
                        "name": "CreateAccount",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.CreateAccountRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.CreateAccountResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/pb.CreateAccountResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/pb.CreateAccountResponse"
                        }
                    }
                }
            }
        },
        "/account/get/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account-Service"
                ],
                "summary": "Find user by id",
                "operationId": "Find account by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Find account by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.FindAccountResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/pb.FindAccountResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/pb.FindAccountResponse"
                        }
                    }
                }
            }
        },
        "/account/update/{id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account-Service"
                ],
                "summary": "Update account",
                "operationId": "Update Account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Create Account",
                        "name": "UpdateAccount",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.UpdateAccountRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.UpdateAccountResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/pb.UpdateAccountResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/pb.UpdateAccountResponse"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication-Service"
                ],
                "summary": "Login user",
                "operationId": "User Login",
                "parameters": [
                    {
                        "description": "User Login",
                        "name": "LoginUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.LoginRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.LoginResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/pb.LoginResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/pb.LoginResponse"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication-Service"
                ],
                "summary": "Register new user",
                "operationId": "User Registration",
                "parameters": [
                    {
                        "description": "User registration",
                        "name": "RegisterUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.RegisterRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.RegisterResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/pb.RegisterResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/pb.RegisterResponse"
                        }
                    }
                }
            }
        },
        "/auth/token-refresh": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication-Service"
                ],
                "summary": "Refresh token for users",
                "operationId": "User RefreshToken",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.TokenRefreshResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/pb.TokenRefreshResponse"
                        }
                    }
                }
            }
        },
        "/user/delete/{id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Section"
                ],
                "summary": "Delete user",
                "operationId": "Delete user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of the user to be deleted",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.RegisterResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/pb.RegisterResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/pb.RegisterResponse"
                        }
                    }
                }
            }
        },
        "/user/finduser/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Section"
                ],
                "summary": "Find user by id",
                "operationId": "Find user by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Find user by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.FindUserResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/pb.FindUserResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/pb.FindUserResponse"
                        }
                    }
                }
            }
        },
        "/user/getusers": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Section"
                ],
                "summary": "Get users",
                "operationId": "Get users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.GetUsersResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/pb.GetUsersResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/pb.GetUsersResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "pb.CreateAccountResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "pb.FindAccountData": {
            "type": "object",
            "properties": {
                "accountHolder": {
                    "type": "string"
                },
                "account_number": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "pb.FindAccountResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/pb.FindAccountData"
                },
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "pb.FindUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "pb.FindUserResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/pb.FindUser"
                },
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "pb.GetUsersResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "user": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pb.User"
                    }
                }
            }
        },
        "pb.LoginResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                },
                "refressh_token": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "pb.RegisterResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "pb.TokenRefreshResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "pb.UpdateAccountResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "pb.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "routes.CreateAccountRequestBody": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "routes.LoginRequestBody": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "routes.RegisterRequestBody": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "routes.UpdateAccountRequestBody": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BearerAuth": {
            "type": "apiKey",
            "name": "authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3005",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "CashierX API",
	Description:      "This is Money management project. You can visit the GitHub repository at https://github.com/SethukumarJ/CashierX-API-Gateway",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
