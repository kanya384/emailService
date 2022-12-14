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
        "contact": {
            "name": "API Support",
            "email": "kanya384@mail.ru"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/delivery/": {
            "get": {
                "description": "Рендерит скрытую гифку для письма и помечает прочитанные письма.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "image/gif"
                ],
                "tags": [
                    "delivery"
                ],
                "summary": "Рендерит скрытую гифку для письма и помечает прочитанные письма.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор рассылки",
                        "name": "deliveryId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Идентификатор подписчика",
                        "name": "subscriberId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "post": {
                "description": "Создает рассылку.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "delivery"
                ],
                "summary": "Создает рассылку.",
                "parameters": [
                    {
                        "description": "Данные по рассылке",
                        "name": "delivery",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/delivery.CrateDeliveryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "201": {
                        "description": "Структура рассылки",
                        "schema": {
                            "$ref": "#/definitions/delivery.CreateDeliveryResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "404 Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/template/": {
            "post": {
                "description": "Добавляет шаблон письма.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "templates"
                ],
                "summary": "Добавляет шаблон письма.",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Файл шаблона",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Структура шаблона",
                        "schema": {
                            "$ref": "#/definitions/template.TemplateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "delivery.CrateDeliveryRequest": {
            "type": "object",
            "required": [
                "sendAt",
                "subject",
                "subscribers",
                "template_id"
            ],
            "properties": {
                "sendAt": {
                    "type": "string",
                    "example": "2022-10-25T15:33:35.304895357+03:00"
                },
                "subject": {
                    "type": "string",
                    "example": "Lorem Ipsum"
                },
                "subscribers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/delivery.DeliverySubscriber"
                    }
                },
                "template_id": {
                    "type": "string",
                    "format": "uuid",
                    "example": "00000000-0000-0000-0000-000000000000"
                }
            }
        },
        "delivery.CreateDeliveryResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "modifiedAt": {
                    "type": "string"
                },
                "subject": {
                    "type": "string"
                },
                "templateID": {
                    "type": "string"
                }
            }
        },
        "delivery.DeliverySubscriber": {
            "type": "object",
            "required": [
                "email",
                "name",
                "surname"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "maximum": 100,
                    "minimum": 10,
                    "example": 15
                },
                "email": {
                    "type": "string",
                    "format": "email",
                    "example": "test01@mail.ru"
                },
                "name": {
                    "type": "string",
                    "example": "Ivan"
                },
                "surname": {
                    "type": "string",
                    "example": "Ivanov"
                }
            }
        },
        "http.ErrorResponse": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                }
            }
        },
        "template.TemplateResponse": {
            "type": "object",
            "required": [
                "createdAt",
                "id",
                "modifiedAt",
                "path"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string",
                    "format": "uuid",
                    "example": "00000000-0000-0000-0000-000000000000"
                },
                "modifiedAt": {
                    "type": "string"
                },
                "path": {
                    "type": "string",
                    "example": "/storage/file"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "email service",
	Description:      "email service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
