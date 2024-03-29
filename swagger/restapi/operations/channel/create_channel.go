// Code generated by go-swagger; DO NOT EDIT.

package channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/berquerant/terraform-slack/swagger/models"
)

// CreateChannelHandlerFunc turns a function with the right signature into a create channel handler
type CreateChannelHandlerFunc func(CreateChannelParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateChannelHandlerFunc) Handle(params CreateChannelParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// CreateChannelHandler interface for that can handle valid create channel params
type CreateChannelHandler interface {
	Handle(CreateChannelParams, *models.Principal) middleware.Responder
}

// NewCreateChannel creates a new http.Handler for the create channel operation
func NewCreateChannel(ctx *middleware.Context, handler CreateChannelHandler) *CreateChannel {
	return &CreateChannel{Context: ctx, Handler: handler}
}

/*
	CreateChannel swagger:route POST /channel channel createChannel

Create a new channel
*/
type CreateChannel struct {
	Context *middleware.Context
	Handler CreateChannelHandler
}

func (o *CreateChannel) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateChannelParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// CreateChannelBody create channel body
//
// swagger:model CreateChannelBody
type CreateChannelBody struct {

	// is private
	// Required: true
	IsPrivate *bool `json:"isPrivate"`

	// name
	// Required: true
	Name *string `json:"name"`

	// team Id
	// Required: true
	TeamID *string `json:"teamId"`
}

// Validate validates this create channel body
func (o *CreateChannelBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIsPrivate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTeamID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateChannelBody) validateIsPrivate(formats strfmt.Registry) error {

	if err := validate.Required("channel"+"."+"isPrivate", "body", o.IsPrivate); err != nil {
		return err
	}

	return nil
}

func (o *CreateChannelBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("channel"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *CreateChannelBody) validateTeamID(formats strfmt.Registry) error {

	if err := validate.Required("channel"+"."+"teamId", "body", o.TeamID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create channel body based on context it is used
func (o *CreateChannelBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateChannelBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateChannelBody) UnmarshalBinary(b []byte) error {
	var res CreateChannelBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
