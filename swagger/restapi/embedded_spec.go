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
    "title": "Slack API for Terraform provider",
    "version": "0.1.0"
  },
  "host": "127.0.0.1:8030",
  "basePath": "/v0",
  "paths": {
    "/channel": {
      "get": {
        "tags": [
          "channel"
        ],
        "summary": "Fetch a channel",
        "operationId": "fetchChannel",
        "parameters": [
          {
            "type": "string",
            "description": "channel id",
            "name": "id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/ChannelModel"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ErrorModel"
            }
          }
        }
      },
      "put": {
        "tags": [
          "channel"
        ],
        "summary": "Update a channel",
        "operationId": "updateChannel",
        "parameters": [
          {
            "type": "string",
            "description": "channel id",
            "name": "id",
            "in": "query",
            "required": true
          },
          {
            "description": "channel spec",
            "name": "channel",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "isArchived": {
                  "type": "boolean"
                },
                "name": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/ChannelModel"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ErrorModel"
            }
          }
        }
      },
      "post": {
        "tags": [
          "channel"
        ],
        "summary": "Create a new channel",
        "operationId": "createChannel",
        "parameters": [
          {
            "description": "channel spec",
            "name": "channel",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "name",
                "isPrivate",
                "teamId"
              ],
              "properties": {
                "isPrivate": {
                  "type": "boolean"
                },
                "name": {
                  "type": "string"
                },
                "teamId": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/ChannelModel"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ErrorModel"
            }
          }
        }
      }
    },
    "/channels": {
      "get": {
        "tags": [
          "channel"
        ],
        "summary": "Fetch all channels",
        "operationId": "fetchChannels",
        "parameters": [
          {
            "type": "string",
            "description": "encoded team id to list channels in, required if token belongs to org-wide app",
            "name": "team_id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/ChannelsModel"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ErrorModel"
            }
          }
        }
      }
    },
    "/health": {
      "get": {
        "tags": [
          "health"
        ],
        "summary": "Check server alive",
        "operationId": "checkAlive",
        "responses": {
          "200": {
            "description": "Alive",
            "schema": {
              "$ref": "#/definitions/HealthModel"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "BaseModel": {
      "type": "object",
      "required": [
        "ok",
        "request_id"
      ],
      "properties": {
        "ok": {
          "type": "boolean"
        },
        "request_id": {
          "type": "string",
          "format": "uuid"
        }
      }
    },
    "Channel": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "isArchived": {
          "type": "boolean"
        },
        "isPrivate": {
          "type": "boolean"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "ChannelModel": {
      "allOf": [
        {
          "$ref": "#/definitions/BaseModel"
        },
        {
          "type": "object",
          "required": [
            "channel"
          ],
          "properties": {
            "channel": {
              "$ref": "#/definitions/Channel"
            }
          }
        }
      ]
    },
    "ChannelsModel": {
      "allOf": [
        {
          "$ref": "#/definitions/BaseModel"
        },
        {
          "type": "object",
          "required": [
            "channels"
          ],
          "properties": {
            "channels": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Channel"
              }
            }
          }
        }
      ]
    },
    "ErrorModel": {
      "allOf": [
        {
          "$ref": "#/definitions/BaseModel"
        },
        {
          "type": "object",
          "properties": {
            "cause": {
              "type": "string"
            },
            "message": {
              "type": "string"
            }
          }
        }
      ]
    },
    "HealthModel": {
      "allOf": [
        {
          "$ref": "#/definitions/BaseModel"
        }
      ]
    },
    "Principal": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "ApiKey": {
      "type": "apiKey",
      "name": "X-API-Key",
      "in": "header"
    }
  },
  "security": [
    {
      "ApiKey": []
    }
  ],
  "tags": [
    {
      "description": "Health check",
      "name": "health"
    },
    {
      "description": "Manage channels",
      "name": "channel"
    }
  ]
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
    "title": "Slack API for Terraform provider",
    "version": "0.1.0"
  },
  "host": "127.0.0.1:8030",
  "basePath": "/v0",
  "paths": {
    "/channel": {
      "get": {
        "tags": [
          "channel"
        ],
        "summary": "Fetch a channel",
        "operationId": "fetchChannel",
        "parameters": [
          {
            "type": "string",
            "description": "channel id",
            "name": "id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/ChannelModel"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ErrorModel"
            }
          }
        }
      },
      "put": {
        "tags": [
          "channel"
        ],
        "summary": "Update a channel",
        "operationId": "updateChannel",
        "parameters": [
          {
            "type": "string",
            "description": "channel id",
            "name": "id",
            "in": "query",
            "required": true
          },
          {
            "description": "channel spec",
            "name": "channel",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "isArchived": {
                  "type": "boolean"
                },
                "name": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/ChannelModel"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ErrorModel"
            }
          }
        }
      },
      "post": {
        "tags": [
          "channel"
        ],
        "summary": "Create a new channel",
        "operationId": "createChannel",
        "parameters": [
          {
            "description": "channel spec",
            "name": "channel",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "name",
                "isPrivate",
                "teamId"
              ],
              "properties": {
                "isPrivate": {
                  "type": "boolean"
                },
                "name": {
                  "type": "string"
                },
                "teamId": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/ChannelModel"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ErrorModel"
            }
          }
        }
      }
    },
    "/channels": {
      "get": {
        "tags": [
          "channel"
        ],
        "summary": "Fetch all channels",
        "operationId": "fetchChannels",
        "parameters": [
          {
            "type": "string",
            "description": "encoded team id to list channels in, required if token belongs to org-wide app",
            "name": "team_id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/ChannelsModel"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ErrorModel"
            }
          }
        }
      }
    },
    "/health": {
      "get": {
        "tags": [
          "health"
        ],
        "summary": "Check server alive",
        "operationId": "checkAlive",
        "responses": {
          "200": {
            "description": "Alive",
            "schema": {
              "$ref": "#/definitions/HealthModel"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "BaseModel": {
      "type": "object",
      "required": [
        "ok",
        "request_id"
      ],
      "properties": {
        "ok": {
          "type": "boolean"
        },
        "request_id": {
          "type": "string",
          "format": "uuid"
        }
      }
    },
    "Channel": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "isArchived": {
          "type": "boolean"
        },
        "isPrivate": {
          "type": "boolean"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "ChannelModel": {
      "allOf": [
        {
          "$ref": "#/definitions/BaseModel"
        },
        {
          "type": "object",
          "required": [
            "channel"
          ],
          "properties": {
            "channel": {
              "$ref": "#/definitions/Channel"
            }
          }
        }
      ]
    },
    "ChannelsModel": {
      "allOf": [
        {
          "$ref": "#/definitions/BaseModel"
        },
        {
          "type": "object",
          "required": [
            "channels"
          ],
          "properties": {
            "channels": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Channel"
              }
            }
          }
        }
      ]
    },
    "ErrorModel": {
      "allOf": [
        {
          "$ref": "#/definitions/BaseModel"
        },
        {
          "type": "object",
          "properties": {
            "cause": {
              "type": "string"
            },
            "message": {
              "type": "string"
            }
          }
        }
      ]
    },
    "HealthModel": {
      "allOf": [
        {
          "$ref": "#/definitions/BaseModel"
        }
      ]
    },
    "Principal": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "ApiKey": {
      "type": "apiKey",
      "name": "X-API-Key",
      "in": "header"
    }
  },
  "security": [
    {
      "ApiKey": []
    }
  ],
  "tags": [
    {
      "description": "Health check",
      "name": "health"
    },
    {
      "description": "Manage channels",
      "name": "channel"
    }
  ]
}`))
}
