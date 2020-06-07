// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Employee Management System",
    "title": "Employee system",
    "version": "0.0.1"
  },
  "host": "localhost:8080",
  "basePath": "/",
  "paths": {
    "/": {
      "get": {
        "description": "get all the employees",
        "operationId": "ListEmployee",
        "responses": {
          "200": {
            "description": "employee response",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/employee"
              }
            }
          },
          "500": {
            "description": "internal server error"
          }
        }
      },
      "post": {
        "operationId": "saveEmployee",
        "parameters": [
          {
            "description": "contains employee information",
            "name": "employee",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/employee"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "employee added",
            "schema": {
              "$ref": "#/definitions/employee"
            }
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/{ID}": {
      "get": {
        "description": "get employee with id provided in url",
        "operationId": "GetEmployeeByID",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the employee",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "employee response",
            "schema": {
              "$ref": "#/definitions/employee"
            }
          },
          "404": {
            "description": "employee not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      },
      "delete": {
        "description": "delete employee with id provided in url",
        "operationId": "DeleteEmployee",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the employee",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "employee deleted"
          },
          "404": {
            "description": "employee not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    }
  },
  "definitions": {
    "employee": {
      "type": "object",
      "properties": {
        "Address": {
          "type": "string"
        },
        "Age": {
          "type": "integer"
        },
        "ID": {
          "type": "string"
        },
        "Name": {
          "type": "string"
        },
        "Phone": {
          "type": "integer"
        },
        "Salary": {
          "type": "integer"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Employee Management System",
    "title": "Employee system",
    "version": "0.0.1"
  },
  "host": "localhost:8080",
  "basePath": "/",
  "paths": {
    "/": {
      "get": {
        "description": "get all the employees",
        "operationId": "ListEmployee",
        "responses": {
          "200": {
            "description": "employee response",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/employee"
              }
            }
          },
          "500": {
            "description": "internal server error"
          }
        }
      },
      "post": {
        "operationId": "saveEmployee",
        "parameters": [
          {
            "description": "contains employee information",
            "name": "employee",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/employee"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "employee added",
            "schema": {
              "$ref": "#/definitions/employee"
            }
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/{ID}": {
      "get": {
        "description": "get employee with id provided in url",
        "operationId": "GetEmployeeByID",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the employee",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "employee response",
            "schema": {
              "$ref": "#/definitions/employee"
            }
          },
          "404": {
            "description": "employee not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      },
      "delete": {
        "description": "delete employee with id provided in url",
        "operationId": "DeleteEmployee",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the employee",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "employee deleted"
          },
          "404": {
            "description": "employee not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    }
  },
  "definitions": {
    "employee": {
      "type": "object",
      "properties": {
        "Address": {
          "type": "string"
        },
        "Age": {
          "type": "integer"
        },
        "ID": {
          "type": "string"
        },
        "Name": {
          "type": "string"
        },
        "Phone": {
          "type": "integer"
        },
        "Salary": {
          "type": "integer"
        }
      }
    }
  }
}`))
}