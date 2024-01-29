// Code generated by go-swagger; DO NOT EDIT.

package channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/berquerant/terraform-slack/swagger/models"
)

// CreateChannelCreatedCode is the HTTP code returned for type CreateChannelCreated
const CreateChannelCreatedCode int = 201

/*
CreateChannelCreated Success

swagger:response createChannelCreated
*/
type CreateChannelCreated struct {

	/*
	  In: Body
	*/
	Payload *models.ChannelModel `json:"body,omitempty"`
}

// NewCreateChannelCreated creates CreateChannelCreated with default headers values
func NewCreateChannelCreated() *CreateChannelCreated {

	return &CreateChannelCreated{}
}

// WithPayload adds the payload to the create channel created response
func (o *CreateChannelCreated) WithPayload(payload *models.ChannelModel) *CreateChannelCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create channel created response
func (o *CreateChannelCreated) SetPayload(payload *models.ChannelModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateChannelCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
CreateChannelDefault Error

swagger:response createChannelDefault
*/
type CreateChannelDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewCreateChannelDefault creates CreateChannelDefault with default headers values
func NewCreateChannelDefault(code int) *CreateChannelDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateChannelDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create channel default response
func (o *CreateChannelDefault) WithStatusCode(code int) *CreateChannelDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create channel default response
func (o *CreateChannelDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create channel default response
func (o *CreateChannelDefault) WithPayload(payload *models.ErrorModel) *CreateChannelDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create channel default response
func (o *CreateChannelDefault) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateChannelDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}