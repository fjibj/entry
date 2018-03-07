// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/entry/server/gen/models"
)

// AuthorizeTemporaryRedirectCode is the HTTP code returned for type AuthorizeTemporaryRedirect
const AuthorizeTemporaryRedirectCode int = 307

/*AuthorizeTemporaryRedirect redirect to homepage of frontend

swagger:response authorizeTemporaryRedirect
*/
type AuthorizeTemporaryRedirect struct {
	/*homepage of frontend

	 */
	Location string `json:"Location"`
	/*set access_token in cookie

	 */
	SetCookie string `json:"Set-Cookie"`
}

// NewAuthorizeTemporaryRedirect creates AuthorizeTemporaryRedirect with default headers values
func NewAuthorizeTemporaryRedirect() *AuthorizeTemporaryRedirect {

	return &AuthorizeTemporaryRedirect{}
}

// WithLocation adds the location to the authorize temporary redirect response
func (o *AuthorizeTemporaryRedirect) WithLocation(location string) *AuthorizeTemporaryRedirect {
	o.Location = location
	return o
}

// SetLocation sets the location to the authorize temporary redirect response
func (o *AuthorizeTemporaryRedirect) SetLocation(location string) {
	o.Location = location
}

// WithSetCookie adds the setCookie to the authorize temporary redirect response
func (o *AuthorizeTemporaryRedirect) WithSetCookie(setCookie string) *AuthorizeTemporaryRedirect {
	o.SetCookie = setCookie
	return o
}

// SetSetCookie sets the setCookie to the authorize temporary redirect response
func (o *AuthorizeTemporaryRedirect) SetSetCookie(setCookie string) {
	o.SetCookie = setCookie
}

// WriteResponse to the client
func (o *AuthorizeTemporaryRedirect) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Location

	location := o.Location
	if location != "" {
		rw.Header().Set("Location", location)
	}

	// response header Set-Cookie

	setCookie := o.SetCookie
	if setCookie != "" {
		rw.Header().Set("Set-Cookie", setCookie)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(307)
}

/*AuthorizeDefault generic error response

swagger:response authorizeDefault
*/
type AuthorizeDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAuthorizeDefault creates AuthorizeDefault with default headers values
func NewAuthorizeDefault(code int) *AuthorizeDefault {
	if code <= 0 {
		code = 500
	}

	return &AuthorizeDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the authorize default response
func (o *AuthorizeDefault) WithStatusCode(code int) *AuthorizeDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the authorize default response
func (o *AuthorizeDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the authorize default response
func (o *AuthorizeDefault) WithPayload(payload *models.Error) *AuthorizeDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the authorize default response
func (o *AuthorizeDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AuthorizeDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
