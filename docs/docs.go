// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
		"/calculator/add": {
			post: {
			summary: Add two numbers,		
			  description: Calculate the addition of the two entered numbers,
			  produces: [
				application/json
			  ],
			  "parameters": [
				{
				  "in": "body",
				  "name": "Numbers",
				  "description": "Enter value of number1 & number2",
				  "properties":{
				  "number1":{
					  "type":"integer",
					  "format":"json"
				  },
	
					"number2":{
						"type":"integer",
						"format":"json"
					},
					},
				   
				  },
				  ],
			  "responses": {
				"200": {
				  "description": "Addition of numbers is successful"
				},
			},
		},
		},
			"/calculator/subtract": {
				post: {
				summary: Subtract two numbers,		
				  description: Calculate the value of the subtraction of the two entered numbers,
				  produces: [
					application/json
				  ],
				  "parameters":[
					{
					  "in": "body",
					  "name": "Numbers",
					  "description": "Enter value of number1 & number2",
					  "properties":{
					  "number1":{
						  "type":"integer",
						  "format":"json"
					  },
		
						"number2":{
							"type":"integer",
							"format":"json"
						},
						},
					   
					  },
					  ],
				  "responses": {
					"200": {
					  "description": "Subtraction of numbers is successful"
					},
				},
			},
			},
				"/calculator/multiply": {
					post: {
					summary: Multiply two numbers,		
					  description: Calculate the value of the multiplicaion of the two entered numbers,
					  produces: [
						application/json
					  ],
					  "parameters":[
						{
						  "in": "body",
						  "name": "Numbers",
						  "description": "Enter value of number1 & number2",
						  "properties":{
						  "number1":{
							  "type":"integer",
							  "format":"json"
						  },
			
							"number2":{
								"type":"integer",
								"format":"json"
							},
							},
						   
						  }
						  ],
					  "responses": {
						"200": {
						  "description": "Multiplication of numbers is successful"
						},
					},
				},
				},
				"/calculator/divide": {
					post: {
					summary: Divide two numbers,		
					  description: Calculate the value of the division of the two entered numbers,
					  produces: [
						application/json
					  ],
					  "parameters":[
						{
						  "in": "body",
						  "name": "Numbers",
						  "description": "Enter value of number1 & number2",
						  "properties":{
						  "number1":{
							  "type":"integer",
							  "format":"json"
						  },
			
							"number2":{
								"type":"integer",
								"format":"json"
							},
							},
						   
						  }
						  ],
					  "responses": {
						"200": {
						  "description": "Division of numbers is successful"
						},
					},
				},
				},
				"/calculator/modulus": {
					post: {
					summary: Modulus of two numbers,		
					  description: Calculate the value of the modulus of the two entered numbers,
					  produces: [
						application/json
					  ],
					  "parameters":[
						{
						  "in": "body",
						  "name": "Numbers",
						  "description": "Enter value of number1 & number2",
						  "properties":{
						  "number1":{
							  "type":"integer",
							  "format":"json"
						  },
			
							"number2":{
								"type":"integer",
								"format":"json"
							},
							},
						   
						  }
						  ],
					  "responses": {
						"200": {
						  "description": "Modulus of numbers is successful"
						},
					},
				},
				},
				"/calculator/square": {
					post: {
					summary: Square of a number,		
					  description: Calculate the value of the square of a number,
					  produces: [
						application/json
					  ],
					  "parameters":[
						{
						  "in": "body",
						  "name": "Numbers",
						  "description": "Enter value of number1",
						  "properties":{
						  "number1":{
							  "type":"integer",
							  "format":"json"
						  },
							},
						   
						  }
						  ],
					  "responses": {
						"200": {
						  "description": "Square of a number is successful"
						},
					},
				},
				},
				"/calculator/power": {
					post: {
					summary: Power of number,		
					  description: Calculate the value of the number1 raised to the power number2,
					  produces: [
						application/json
					  ],
					  "parameters":[
						{
						  "in": "body",
						  "name": "Numbers",
						  "description": "Enter value of number1 & number2",
						  "properties":{
						  "number1":{
							  "type":"integer",
							  "format":"json"
						  },
			
							"number2":{
								"type":"integer",
								"format":"json"
							},
							},
						   
						  }
						  ],
					  "responses": {
						"200": {
						  "description": "Power of numbers is successful"
						},
					},
				},
				},
				"/calculator/sqrt": {
					post: {
					summary: Suqare root of a number,		
					  description: Calculate the value of the square root of the entered numbers,
					  produces: [
						application/json
					  ],
					  "parameters":[
						{
						  "in": "body",
						  "name": "Numbers",
						  "description": "Enter value of number1",
						  "properties":{
						  "number1":{
							  "type":"integer",
							  "format":"json"
						  },
			
							},
						   
						  }
						  ],
					  "responses": {
						"200": {
						  "description": "Square root of number is successful"
						},
					},
				},
				},
				"/calculator/GetRecord": {
					post: {
					summary: Post record by user id,		
					  description: Returns the reocrd by user id,
					  produces: [
						application/json
					  ],
					  "parameters":[
						{
						  "in": "body",
						  "name": "ID",
						  "description": "Enter user id",
						  "properties":{
						  "id":{
							  "type":"integer",
							  "format":"json"
						  },
							},
						   
						  },
						  ],
					  "responses": {
						"200": {
						  "description": "Record has been displayed"
						},
					},
				},
				},
		}, 
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
	Host:        "localhost:1323",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Calculator",
	Description: "Different operations of calculator",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
