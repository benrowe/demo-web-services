// Code generated by go-swagger; DO NOT EDIT.

package employee

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListEmployeesHandlerFunc turns a function with the right signature into a list employees handler
type ListEmployeesHandlerFunc func(ListEmployeesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListEmployeesHandlerFunc) Handle(params ListEmployeesParams) middleware.Responder {
	return fn(params)
}

// ListEmployeesHandler interface for that can handle valid list employees params
type ListEmployeesHandler interface {
	Handle(ListEmployeesParams) middleware.Responder
}

// NewListEmployees creates a new http.Handler for the list employees operation
func NewListEmployees(ctx *middleware.Context, handler ListEmployeesHandler) *ListEmployees {
	return &ListEmployees{Context: ctx, Handler: handler}
}

/*ListEmployees swagger:route GET /employee employee listEmployees

ListEmployees list employees API

*/
type ListEmployees struct {
	Context *middleware.Context
	Handler ListEmployeesHandler
}

func (o *ListEmployees) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListEmployeesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
