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
    "application/vnd.laincloud.entry.v3+json"
  ],
  "produces": [
    "application/vnd.laincloud.entry.v3+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Entry",
    "version": "3.2.2"
  },
  "paths": {
    "/api/authorize": {
      "get": {
        "tags": [
          "auth"
        ],
        "operationId": "authorize",
        "parameters": [
          {
            "type": "string",
            "description": "Authorization code",
            "name": "code",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "307": {
            "description": "redirect to homepage of frontend",
            "headers": {
              "Location": {
                "type": "string",
                "description": "homepage of frontend"
              },
              "Set-Cookie": {
                "type": "string",
                "description": "set access_token in cookie"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/api/commands": {
      "get": {
        "tags": [
          "commands"
        ],
        "operationId": "listCommands",
        "parameters": [
          {
            "type": "string",
            "description": "Cookie with access_token",
            "name": "Cookie",
            "in": "header",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "default": 0,
            "description": "Unix timestamp(unit: second)",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "default": 20,
            "name": "limit",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "default": 0,
            "name": "offset",
            "in": "query"
          },
          {
            "type": "string",
            "description": "MySQL LIKE pattern match",
            "name": "user",
            "in": "query"
          },
          {
            "type": "string",
            "description": "MySQL LIKE pattern match",
            "name": "app_name",
            "in": "query"
          },
          {
            "type": "string",
            "description": "query pattern(MySQL LIKE pattern match)",
            "name": "content",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "name": "session_id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list the commands",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/command"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/api/config": {
      "get": {
        "tags": [
          "config"
        ],
        "operationId": "getConfig",
        "responses": {
          "200": {
            "description": "the configuration of entry",
            "schema": {
              "$ref": "#/definitions/config"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/api/logout": {
      "get": {
        "tags": [
          "auth"
        ],
        "operationId": "logout",
        "responses": {
          "200": {
            "description": "logout",
            "headers": {
              "Set-Cookie": {
                "type": "string",
                "description": "delete access_token in cookie"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/api/ping": {
      "get": {
        "tags": [
          "ping"
        ],
        "operationId": "ping",
        "responses": {
          "200": {
            "description": "this server is healthy",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/api/sessions": {
      "get": {
        "tags": [
          "sessions"
        ],
        "operationId": "listSessions",
        "parameters": [
          {
            "type": "string",
            "description": "Cookie with access_token",
            "name": "Cookie",
            "in": "header",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "name": "session_id",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "default": 0,
            "description": "Unix timestamp(unit: second)",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "default": 20,
            "name": "limit",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "default": 0,
            "name": "offset",
            "in": "query"
          },
          {
            "type": "string",
            "description": "MySQL LIKE pattern match",
            "name": "user",
            "in": "query"
          },
          {
            "type": "string",
            "description": "MySQL LIKE pattern match",
            "name": "app_name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list the sessions",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/session"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/api/sessions/{session_id}/replay": {
      "get": {
        "tags": [
          "sessions"
        ],
        "operationId": "replaySession",
        "responses": {
          "200": {
            "description": "replay the session"
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "session_id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/attach": {
      "get": {
        "tags": [
          "container"
        ],
        "operationId": "attachContainer",
        "responses": {
          "200": {
            "description": "attach to container's stdout/stderr"
          }
        }
      }
    },
    "/enter": {
      "get": {
        "tags": [
          "container"
        ],
        "operationId": "enterContainer",
        "responses": {
          "200": {
            "description": "enter to the container"
          }
        }
      }
    }
  },
  "definitions": {
    "command": {
      "type": "object",
      "properties": {
        "app_name": {
          "type": "string"
        },
        "command_id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "content": {
          "type": "string"
        },
        "created_at": {
          "description": "Unix timestamp(unit: second)",
          "type": "integer",
          "format": "int64"
        },
        "instance_no": {
          "type": "string"
        },
        "proc_name": {
          "type": "string"
        },
        "session_id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "user": {
          "type": "string"
        }
      }
    },
    "config": {
      "type": "object",
      "required": [
        "sso"
      ],
      "properties": {
        "sso": {
          "type": "object",
          "required": [
            "domain",
            "client_id",
            "redirect_uri",
            "scope"
          ],
          "properties": {
            "client_id": {
              "type": "string"
            },
            "domain": {
              "type": "string"
            },
            "redirect_uri": {
              "type": "string"
            },
            "scope": {
              "type": "string"
            }
          }
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "session": {
      "type": "object",
      "properties": {
        "app_name": {
          "type": "string"
        },
        "container_id": {
          "type": "string"
        },
        "created_at": {
          "description": "Unix timestamp(unit: second)",
          "type": "integer",
          "format": "int64"
        },
        "ended_at": {
          "description": "Unix timestamp(unit: second)",
          "type": "integer",
          "format": "int64"
        },
        "instance_no": {
          "type": "string"
        },
        "node_ip": {
          "type": "string"
        },
        "proc_name": {
          "type": "string"
        },
        "session_id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "source_ip": {
          "type": "string"
        },
        "status": {
          "type": "string"
        },
        "user": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/vnd.laincloud.entry.v3+json"
  ],
  "produces": [
    "application/vnd.laincloud.entry.v3+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Entry",
    "version": "3.2.2"
  },
  "paths": {
    "/api/authorize": {
      "get": {
        "tags": [
          "auth"
        ],
        "operationId": "authorize",
        "parameters": [
          {
            "type": "string",
            "description": "Authorization code",
            "name": "code",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "307": {
            "description": "redirect to homepage of frontend",
            "headers": {
              "Location": {
                "type": "string",
                "description": "homepage of frontend"
              },
              "Set-Cookie": {
                "type": "string",
                "description": "set access_token in cookie"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/api/commands": {
      "get": {
        "tags": [
          "commands"
        ],
        "operationId": "listCommands",
        "parameters": [
          {
            "type": "string",
            "description": "Cookie with access_token",
            "name": "Cookie",
            "in": "header",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "default": 0,
            "description": "Unix timestamp(unit: second)",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "default": 20,
            "name": "limit",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "default": 0,
            "name": "offset",
            "in": "query"
          },
          {
            "type": "string",
            "description": "MySQL LIKE pattern match",
            "name": "user",
            "in": "query"
          },
          {
            "type": "string",
            "description": "MySQL LIKE pattern match",
            "name": "app_name",
            "in": "query"
          },
          {
            "type": "string",
            "description": "query pattern(MySQL LIKE pattern match)",
            "name": "content",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "name": "session_id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list the commands",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/command"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/api/config": {
      "get": {
        "tags": [
          "config"
        ],
        "operationId": "getConfig",
        "responses": {
          "200": {
            "description": "the configuration of entry",
            "schema": {
              "$ref": "#/definitions/config"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/api/logout": {
      "get": {
        "tags": [
          "auth"
        ],
        "operationId": "logout",
        "responses": {
          "200": {
            "description": "logout",
            "headers": {
              "Set-Cookie": {
                "type": "string",
                "description": "delete access_token in cookie"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/api/ping": {
      "get": {
        "tags": [
          "ping"
        ],
        "operationId": "ping",
        "responses": {
          "200": {
            "description": "this server is healthy",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/api/sessions": {
      "get": {
        "tags": [
          "sessions"
        ],
        "operationId": "listSessions",
        "parameters": [
          {
            "type": "string",
            "description": "Cookie with access_token",
            "name": "Cookie",
            "in": "header",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "name": "session_id",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "default": 0,
            "description": "Unix timestamp(unit: second)",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "default": 20,
            "name": "limit",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "default": 0,
            "name": "offset",
            "in": "query"
          },
          {
            "type": "string",
            "description": "MySQL LIKE pattern match",
            "name": "user",
            "in": "query"
          },
          {
            "type": "string",
            "description": "MySQL LIKE pattern match",
            "name": "app_name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list the sessions",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/session"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/api/sessions/{session_id}/replay": {
      "get": {
        "tags": [
          "sessions"
        ],
        "operationId": "replaySession",
        "responses": {
          "200": {
            "description": "replay the session"
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "session_id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/attach": {
      "get": {
        "tags": [
          "container"
        ],
        "operationId": "attachContainer",
        "responses": {
          "200": {
            "description": "attach to container's stdout/stderr"
          }
        }
      }
    },
    "/enter": {
      "get": {
        "tags": [
          "container"
        ],
        "operationId": "enterContainer",
        "responses": {
          "200": {
            "description": "enter to the container"
          }
        }
      }
    }
  },
  "definitions": {
    "command": {
      "type": "object",
      "properties": {
        "app_name": {
          "type": "string"
        },
        "command_id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "content": {
          "type": "string"
        },
        "created_at": {
          "description": "Unix timestamp(unit: second)",
          "type": "integer",
          "format": "int64"
        },
        "instance_no": {
          "type": "string"
        },
        "proc_name": {
          "type": "string"
        },
        "session_id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "user": {
          "type": "string"
        }
      }
    },
    "config": {
      "type": "object",
      "required": [
        "sso"
      ],
      "properties": {
        "sso": {
          "$ref": "#/definitions/configSso"
        }
      }
    },
    "configSso": {
      "type": "object",
      "required": [
        "domain",
        "client_id",
        "redirect_uri",
        "scope"
      ],
      "properties": {
        "client_id": {
          "type": "string"
        },
        "domain": {
          "type": "string"
        },
        "redirect_uri": {
          "type": "string"
        },
        "scope": {
          "type": "string"
        }
      },
      "x-go-gen-location": "models"
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "session": {
      "type": "object",
      "properties": {
        "app_name": {
          "type": "string"
        },
        "container_id": {
          "type": "string"
        },
        "created_at": {
          "description": "Unix timestamp(unit: second)",
          "type": "integer",
          "format": "int64"
        },
        "ended_at": {
          "description": "Unix timestamp(unit: second)",
          "type": "integer",
          "format": "int64"
        },
        "instance_no": {
          "type": "string"
        },
        "node_ip": {
          "type": "string"
        },
        "proc_name": {
          "type": "string"
        },
        "session_id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "source_ip": {
          "type": "string"
        },
        "status": {
          "type": "string"
        },
        "user": {
          "type": "string"
        }
      }
    }
  }
}`))
}
