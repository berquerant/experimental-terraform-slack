package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

func NewErrorModel(cause, message string) ErrorModel {
	return ErrorModel{
		BaseModel: NewBaseModel(false),
		Cause:     cause,
		Message:   message,
	}
}

func NewUUID() string {
	return uuid.New().String()
}

func NewBaseModel(ok bool) BaseModel {
	id := strfmt.UUID(NewUUID())
	return BaseModel{
		Ok:        &ok,
		RequestID: &id,
	}
}
