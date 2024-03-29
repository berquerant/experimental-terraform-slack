// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Channel channel
//
// swagger:model Channel
type Channel struct {

	// id
	ID string `json:"id,omitempty"`

	// is archived
	IsArchived bool `json:"isArchived,omitempty"`

	// is private
	IsPrivate bool `json:"isPrivate,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this channel
func (m *Channel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this channel based on context it is used
func (m *Channel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Channel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Channel) UnmarshalBinary(b []byte) error {
	var res Channel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
