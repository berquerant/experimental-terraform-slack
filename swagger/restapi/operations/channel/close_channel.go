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

	"github.com/berquerant/terraform-slack/swagger/models"
)

// CloseChannelHandlerFunc turns a function with the right signature into a close channel handler
type CloseChannelHandlerFunc func(CloseChannelParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn CloseChannelHandlerFunc) Handle(params CloseChannelParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// CloseChannelHandler interface for that can handle valid close channel params
type CloseChannelHandler interface {
	Handle(CloseChannelParams, *models.Principal) middleware.Responder
}

// NewCloseChannel creates a new http.Handler for the close channel operation
func NewCloseChannel(ctx *middleware.Context, handler CloseChannelHandler) *CloseChannel {
	return &CloseChannel{Context: ctx, Handler: handler}
}

/*
	CloseChannel swagger:route DELETE /channel channel closeChannel

Close a channel
*/
type CloseChannel struct {
	Context *middleware.Context
	Handler CloseChannelHandler
}

func (o *CloseChannel) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCloseChannelParams()
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

// CloseChannelOKBody close channel o k body
//
// swagger:model CloseChannelOKBody
type CloseChannelOKBody struct {
	models.BaseModel

	// already closed
	AlreadyClosed bool `json:"alreadyClosed,omitempty"`

	// noop
	Noop bool `json:"noop,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *CloseChannelOKBody) UnmarshalJSON(raw []byte) error {
	// CloseChannelOKBodyAO0
	var closeChannelOKBodyAO0 models.BaseModel
	if err := swag.ReadJSON(raw, &closeChannelOKBodyAO0); err != nil {
		return err
	}
	o.BaseModel = closeChannelOKBodyAO0

	// CloseChannelOKBodyAO1
	var dataCloseChannelOKBodyAO1 struct {
		AlreadyClosed bool `json:"alreadyClosed,omitempty"`

		Noop bool `json:"noop,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataCloseChannelOKBodyAO1); err != nil {
		return err
	}

	o.AlreadyClosed = dataCloseChannelOKBodyAO1.AlreadyClosed

	o.Noop = dataCloseChannelOKBodyAO1.Noop

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o CloseChannelOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	closeChannelOKBodyAO0, err := swag.WriteJSON(o.BaseModel)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, closeChannelOKBodyAO0)
	var dataCloseChannelOKBodyAO1 struct {
		AlreadyClosed bool `json:"alreadyClosed,omitempty"`

		Noop bool `json:"noop,omitempty"`
	}

	dataCloseChannelOKBodyAO1.AlreadyClosed = o.AlreadyClosed

	dataCloseChannelOKBodyAO1.Noop = o.Noop

	jsonDataCloseChannelOKBodyAO1, errCloseChannelOKBodyAO1 := swag.WriteJSON(dataCloseChannelOKBodyAO1)
	if errCloseChannelOKBodyAO1 != nil {
		return nil, errCloseChannelOKBodyAO1
	}
	_parts = append(_parts, jsonDataCloseChannelOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this close channel o k body
func (o *CloseChannelOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.BaseModel
	if err := o.BaseModel.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this close channel o k body based on the context it is used
func (o *CloseChannelOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.BaseModel
	if err := o.BaseModel.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *CloseChannelOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CloseChannelOKBody) UnmarshalBinary(b []byte) error {
	var res CloseChannelOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
