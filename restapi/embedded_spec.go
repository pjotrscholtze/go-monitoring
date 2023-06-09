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
  "swagger": "2.0",
  "info": {
    "title": "Go monitoring",
    "version": "1.0.0"
  },
  "basePath": "/api",
  "paths": {
    "/check": {
      "get": {
        "summary": "List all checks",
        "operationId": "listAllChecks",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Check"
              }
            }
          }
        }
      }
    },
    "/check/{targetName}/{checkName}": {
      "get": {
        "summary": "Get historical data of a check",
        "operationId": "getTargetChecks",
        "parameters": [
          {
            "type": "string",
            "description": "Name of the target",
            "name": "targetName",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Name of the check",
            "name": "checkName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/CheckResult"
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Check": {
      "required": [
        "TargetName",
        "CheckName",
        "Schedule",
        "LastCheckResult"
      ],
      "properties": {
        "CheckName": {
          "type": "string"
        },
        "LastCheckResult": {
          "$ref": "#/definitions/CheckResult"
        },
        "Schedule": {
          "type": "string"
        },
        "TargetName": {
          "type": "string"
        }
      }
    },
    "CheckResult": {
      "required": [
        "Error",
        "Success",
        "LastCheck",
        "Message"
      ],
      "properties": {
        "Error": {
          "type": "string"
        },
        "LastCheck": {
          "type": "string"
        },
        "Message": {
          "type": "string"
        },
        "Success": {
          "type": "boolean"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "Go monitoring",
    "version": "1.0.0"
  },
  "basePath": "/api",
  "paths": {
    "/check": {
      "get": {
        "summary": "List all checks",
        "operationId": "listAllChecks",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Check"
              }
            }
          }
        }
      }
    },
    "/check/{targetName}/{checkName}": {
      "get": {
        "summary": "Get historical data of a check",
        "operationId": "getTargetChecks",
        "parameters": [
          {
            "type": "string",
            "description": "Name of the target",
            "name": "targetName",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Name of the check",
            "name": "checkName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/CheckResult"
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Check": {
      "required": [
        "TargetName",
        "CheckName",
        "Schedule",
        "LastCheckResult"
      ],
      "properties": {
        "CheckName": {
          "type": "string"
        },
        "LastCheckResult": {
          "$ref": "#/definitions/CheckResult"
        },
        "Schedule": {
          "type": "string"
        },
        "TargetName": {
          "type": "string"
        }
      }
    },
    "CheckResult": {
      "required": [
        "Error",
        "Success",
        "LastCheck",
        "Message"
      ],
      "properties": {
        "Error": {
          "type": "string"
        },
        "LastCheck": {
          "type": "string"
        },
        "Message": {
          "type": "string"
        },
        "Success": {
          "type": "boolean"
        }
      }
    }
  }
}`))
}
