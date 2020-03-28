// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-03-29 08:07:22.689326 +0900 JST m=+0.043163531

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
            "name": "Contact us",
            "url": "https://datascientist-toolbox.com/",
            "email": "datascientist.toolbox.com@gmail.com"
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
        "/companies": {
            "get": {
                "description": "Dashboardアプリに登録されている企業の基本情報を取得します。",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "companies"
                ],
                "summary": "Dashboardアプリに登録されている企業情報一覧を取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CompanyWrapper"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Company": {
            "type": "object",
            "properties": {
                "companyId": {
                    "description": "システム内で設定する企業毎に一意となるコード",
                    "type": "string"
                },
                "companyName": {
                    "description": "企業名",
                    "type": "string"
                },
                "countryCode": {
                    "description": "国コード",
                    "type": "string"
                },
                "foundationDate": {
                    "description": "設立日(月)",
                    "type": "string"
                },
                "latitude": {
                    "description": "本社所在地緯度",
                    "type": "number"
                },
                "listedMarket": {
                    "description": "上場市場",
                    "type": "string"
                },
                "longitude": {
                    "description": "本社所在地経度",
                    "type": "number"
                },
                "stockCode": {
                    "description": "銘柄コード",
                    "type": "string"
                }
            }
        },
        "models.CompanyWrapper": {
            "type": "object",
            "properties": {
                "companies": {
                    "description": "企業名・銘柄コード・設立日などの会社の基本情報を保持するクラスの配列",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Company"
                    }
                },
                "timestamp": {
                    "description": "クライアントへレスポンスを送信する日時",
                    "type": "string"
                }
            }
        },
        "models.HTTPError": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "エラーに対する付加的な情報",
                    "type": "string"
                },
                "status": {
                    "description": "HTTP Status Code",
                    "type": "integer"
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
	Host:        "34.84.101.196:30080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Stock Prediction API",
	Description: "Demonstration api for dshack project",
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
