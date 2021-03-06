// Code generated by go-swagger; DO NOT EDIT.

package sessions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListSessionsHandlerFunc turns a function with the right signature into a list sessions handler
type ListSessionsHandlerFunc func(ListSessionsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListSessionsHandlerFunc) Handle(params ListSessionsParams) middleware.Responder {
	return fn(params)
}

// ListSessionsHandler interface for that can handle valid list sessions params
type ListSessionsHandler interface {
	Handle(ListSessionsParams) middleware.Responder
}

// NewListSessions creates a new http.Handler for the list sessions operation
func NewListSessions(ctx *middleware.Context, handler ListSessionsHandler) *ListSessions {
	return &ListSessions{Context: ctx, Handler: handler}
}

/*ListSessions swagger:route GET /api/sessions sessions listSessions

ListSessions list sessions API

*/
type ListSessions struct {
	Context *middleware.Context
	Handler ListSessionsHandler
}

func (o *ListSessions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListSessionsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
