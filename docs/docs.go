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
            "name": "pilaoban",
            "url": "https://github.com/jkuup",
            "email": "alphkiee@gmail.com"
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
        "/api/ping": {
            "get": {
                "description": "查看调用接口是否能够ping通",
                "produces": [
                    "application/json"
                ],
                "summary": "Get Ping",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    }
                }
            }
        },
        "/api/v1/operlog/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据ID批量删除日志",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete /api/v1/operlog/delete",
                "parameters": [
                    {
                        "description": "ids",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.IdsReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    }
                }
            }
        },
        "/api/v1/operlog/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "列出所有操作日志",
                "produces": [
                    "application/json"
                ],
                "summary": "Get /api/v1/operlog/list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "method",
                        "name": "method",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "path",
                        "name": "path",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "ip",
                        "name": "ip",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    }
                }
            }
        },
        "/api/v1/role/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "创建角色",
                "produces": [
                    "application/json"
                ],
                "summary": "Post /api/v1/role/create",
                "parameters": [
                    {
                        "description": "name, keyword, desc, creator",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateRoleReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    }
                }
            }
        },
        "/api/v1/role/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据ID批量删除角色",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete /api/v1/role/delete",
                "parameters": [
                    {
                        "description": "ids",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.IdsReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    }
                }
            }
        },
        "/api/v1/role/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "列出所有角色",
                "produces": [
                    "application/json"
                ],
                "summary": "Get /api/v1/role/list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "keyword",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "status",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "creator",
                        "name": "creator",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    }
                }
            }
        },
        "/api/v1/role/perms/:roleId": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取当前用户信息",
                "produces": [
                    "application/json"
                ],
                "summary": "Get /api/v1/role/perms/:roleId",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "roleId",
                        "name": "roleId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    }
                }
            }
        },
        "/api/v1/role/perms/update/{roleId}": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据角色 ID来更新角色权限信息",
                "produces": [
                    "application/json"
                ],
                "summary": "Patch /api/v1/role/perms/update/:roleId",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "roleId",
                        "name": "roleId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "menus_id, apis_id",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateRolePermsReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    }
                }
            }
        },
        "/api/v1/role/update/{roleId}": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据role ID来更新角色基本信息",
                "produces": [
                    "application/json"
                ],
                "summary": "Patch /api/v1/role//update/:roleId",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "roleId",
                        "name": "roleId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    }
                }
            }
        },
        "/api/v1/user/changePwd": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "更改用户的密码",
                "produces": [
                    "application/json"
                ],
                "summary": "Put /api/v1/user/changePwd",
                "parameters": [
                    {
                        "description": "old_password, new_password",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ChangePwdReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    }
                }
            }
        },
        "/api/v1/user/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "创建用户",
                "produces": [
                    "application/json"
                ],
                "summary": "Post /api/v1/user/create",
                "parameters": [
                    {
                        "description": "username, password, name, role_id",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    }
                }
            }
        },
        "/api/v1/user/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据ID批量删除用户",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete /api/v1/user/delete",
                "parameters": [
                    {
                        "description": "ids",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.IdsReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    }
                }
            }
        },
        "/api/v1/user/info": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取当前用户信息",
                "produces": [
                    "application/json"
                ],
                "summary": "Get /api/v1/user/info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    }
                }
            }
        },
        "/api/v1/user/info/update/{userId}": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据user ID来更新用户基本信息",
                "produces": [
                    "application/json"
                ],
                "summary": "Patch /api/v1/user/info/update/:userId",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "userId",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "mobile, name, email",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateUserBaseInfoReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    }
                }
            }
        },
        "/api/v1/user/info/uploadImg": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "上传头像",
                "produces": [
                    "application/json"
                ],
                "summary": "Post /api/v1/user/info/uploadImg",
                "parameters": [
                    {
                        "type": "file",
                        "description": "avatar",
                        "name": "avatar",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    }
                }
            }
        },
        "/api/v1/user/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "列出所有用户",
                "produces": [
                    "application/json"
                ],
                "summary": "Get /api/v1/user/list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "mobile",
                        "name": "mobile",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "status",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "creator",
                        "name": "creator",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "dept_id",
                        "name": "dept_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    }
                }
            }
        },
        "/api/v1/user/update/{userId}": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "更新用户根据 user ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Patch /api/v1/user/update/:userId",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "userId",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "mobile, name, email, password",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.RespInfo"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.ChangePwdReq": {
            "type": "object",
            "required": [
                "new_password",
                "old_password"
            ],
            "properties": {
                "new_password": {
                    "type": "string"
                },
                "old_password": {
                    "type": "string"
                }
            }
        },
        "request.CreateRoleReq": {
            "type": "object",
            "required": [
                "keyword",
                "name"
            ],
            "properties": {
                "creator": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "keyword": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "request.CreateUserReq": {
            "type": "object",
            "required": [
                "name",
                "password",
                "role_id",
                "username"
            ],
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "creator": {
                    "type": "string"
                },
                "dept_id": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "mobile": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role_id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "request.IdsReq": {
            "type": "object",
            "properties": {
                "ids": {
                    "description": "传多个id",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "request.UpdateRolePermsReq": {
            "type": "object",
            "properties": {
                "apis_id": {
                    "description": "传多个id",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "menus_id": {
                    "description": "传多个id",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "request.UpdateUserBaseInfoReq": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "mobile": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "request.UpdateUserReq": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "dept_id": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "mobile": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role_id": {
                    "type": "integer"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "response.RespInfo": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "错误代码代码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据内容",
                    "type": "object"
                },
                "message": {
                    "description": "消息提示",
                    "type": "string"
                },
                "status": {
                    "description": "状态",
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	Version:     "2.0",
	Host:        "127.0.0.1:9000",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Go-Xops",
	Description: "Go-Xops swagger接口文档",
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
