// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAuthorizeParams creates a new AuthorizeParams object
// no default values defined in spec.
func NewAuthorizeParams() AuthorizeParams {

	return AuthorizeParams{}
}

// AuthorizeParams contains all the bound params for the authorize operation
// typically these are obtained from a http.Request
//
// swagger:parameters authorize
type AuthorizeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Authorization code
	  Required: true
	  In: query
	*/
	Code string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAuthorizeParams() beforehand.
func (o *AuthorizeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCode, qhkCode, _ := qs.GetOK("code")
	if err := o.bindCode(qCode, qhkCode, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AuthorizeParams) bindCode(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("code", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("code", "query", raw); err != nil {
		return err
	}

	o.Code = raw

	return nil
}
