// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/entry/server/gen/models"
)

// GetConfigOKCode is the HTTP code returned for type GetConfigOK
const GetConfigOKCode int = 200

/*GetConfigOK the configuration of entry

swagger:response getConfigOK
*/
type GetConfigOK struct {

	/*
	  In: Body
	*/
	Payload *models.Config `json:"body,omitempty"`
}

// NewGetConfigOK creates GetConfigOK with default headers values
func NewGetConfigOK() *GetConfigOK {

	return &GetConfigOK{}
}

// WithPayload adds the payload to the get config o k response
func (o *GetConfigOK) WithPayload(payload *models.Config) *GetConfigOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config o k response
func (o *GetConfigOK) SetPayload(payload *models.Config) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetConfigDefault generic error response

swagger:response getConfigDefault
*/
type GetConfigDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigDefault creates GetConfigDefault with default headers values
func NewGetConfigDefault(code int) *GetConfigDefault {
	if code <= 0 {
		code = 500
	}

	return &GetConfigDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get config default response
func (o *GetConfigDefault) WithStatusCode(code int) *GetConfigDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get config default response
func (o *GetConfigDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get config default response
func (o *GetConfigDefault) WithPayload(payload *models.Error) *GetConfigDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config default response
func (o *GetConfigDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
