// Code generated by go-swagger; DO NOT EDIT.

package channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/berquerant/terraform-slack/swagger/models"
)

// UpdateChannelHandlerFunc turns a function with the right signature into a update channel handler
type UpdateChannelHandlerFunc func(UpdateChannelParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateChannelHandlerFunc) Handle(params UpdateChannelParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// UpdateChannelHandler interface for that can handle valid update channel params
type UpdateChannelHandler interface {
	Handle(UpdateChannelParams, *models.Principal) middleware.Responder
}

// NewUpdateChannel creates a new http.Handler for the update channel operation
func NewUpdateChannel(ctx *middleware.Context, handler UpdateChannelHandler) *UpdateChannel {
	return &UpdateChannel{Context: ctx, Handler: handler}
}

/*
	UpdateChannel swagger:route PUT /channel channel updateChannel

Update a channel
*/
type UpdateChannel struct {
	Context *middleware.Context
	Handler UpdateChannelHandler
}

func (o *UpdateChannel) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateChannelParams()
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

// UpdateChannelBody update channel body
//
// swagger:model UpdateChannelBody
type UpdateChannelBody struct {

	// is archived
	IsArchived bool `json:"isArchived,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this update channel body
func (o *UpdateChannelBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update channel body based on context it is used
func (o *UpdateChannelBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateChannelBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateChannelBody) UnmarshalBinary(b []byte) error {
	var res UpdateChannelBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
