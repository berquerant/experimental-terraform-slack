// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ErrorModel error model
//
// swagger:model ErrorModel
type ErrorModel struct {
	BaseModel

	// cause
	Cause string `json:"cause,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ErrorModel) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseModel
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseModel = aO0

	// AO1
	var dataAO1 struct {
		Cause string `json:"cause,omitempty"`

		Message string `json:"message,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Cause = dataAO1.Cause

	m.Message = dataAO1.Message

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ErrorModel) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseModel)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Cause string `json:"cause,omitempty"`

		Message string `json:"message,omitempty"`
	}

	dataAO1.Cause = m.Cause

	dataAO1.Message = m.Message

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this error model
func (m *ErrorModel) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseModel
	if err := m.BaseModel.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this error model based on the context it is used
func (m *ErrorModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseModel
	if err := m.BaseModel.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorModel) UnmarshalBinary(b []byte) error {
	var res ErrorModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
