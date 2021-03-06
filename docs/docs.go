// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
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
            "name": "1362",
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
        "/age": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AgeController 年龄相关"
                ],
                "summary": "获取最新age",
                "responses": {}
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AgeController 年龄相关"
                ],
                "summary": "设置最新age",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "object",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.AgeModel"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/buffer": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CarInventoryController 库存相关"
                ],
                "summary": "获取最新库存数量",
                "responses": {}
            }
        },
        "/car": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CarInventoryController 库存相关"
                ],
                "summary": "获取最新carid",
                "responses": {}
            }
        },
        "/rate": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CarInventoryController 库存相关"
                ],
                "summary": "获取最新售卖率",
                "responses": {}
            }
        }
    },
    "definitions": {
        "models.AgeModel": {
            "type": "object",
            "properties": {
                "age": {
                    "description": "年龄",
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "124.220.12.138:8888",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "ExamCodeAboutTeslaByGin",
	Description:      "ExamCodeAboutTeslaByGin",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
