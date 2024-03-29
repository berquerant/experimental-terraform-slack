// Code generated by go-swagger; DO NOT EDIT.

package channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/berquerant/terraform-slack/swagger/models"
)

// FetchChannelsOKCode is the HTTP code returned for type FetchChannelsOK
const FetchChannelsOKCode int = 200

/*
FetchChannelsOK Success

swagger:response fetchChannelsOK
*/
type FetchChannelsOK struct {

	/*
	  In: Body
	*/
	Payload *models.ChannelsModel `json:"body,omitempty"`
}

// NewFetchChannelsOK creates FetchChannelsOK with default headers values
func NewFetchChannelsOK() *FetchChannelsOK {

	return &FetchChannelsOK{}
}

// WithPayload adds the payload to the fetch channels o k response
func (o *FetchChannelsOK) WithPayload(payload *models.ChannelsModel) *FetchChannelsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the fetch channels o k response
func (o *FetchChannelsOK) SetPayload(payload *models.ChannelsModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FetchChannelsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
FetchChannelsDefault Error

swagger:response fetchChannelsDefault
*/
type FetchChannelsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewFetchChannelsDefault creates FetchChannelsDefault with default headers values
func NewFetchChannelsDefault(code int) *FetchChannelsDefault {
	if code <= 0 {
		code = 500
	}

	return &FetchChannelsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the fetch channels default response
func (o *FetchChannelsDefault) WithStatusCode(code int) *FetchChannelsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the fetch channels default response
func (o *FetchChannelsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the fetch channels default response
func (o *FetchChannelsDefault) WithPayload(payload *models.ErrorModel) *FetchChannelsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the fetch channels default response
func (o *FetchChannelsDefault) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FetchChannelsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
