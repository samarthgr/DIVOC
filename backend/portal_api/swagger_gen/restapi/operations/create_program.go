// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/divoc/portal-api/swagger_gen/models"
)

// CreateProgramHandlerFunc turns a function with the right signature into a create program handler
type CreateProgramHandlerFunc func(CreateProgramParams, *models.JWTClaimBody) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateProgramHandlerFunc) Handle(params CreateProgramParams, principal *models.JWTClaimBody) middleware.Responder {
	return fn(params, principal)
}

// CreateProgramHandler interface for that can handle valid create program params
type CreateProgramHandler interface {
	Handle(CreateProgramParams, *models.JWTClaimBody) middleware.Responder
}

// NewCreateProgram creates a new http.Handler for the create program operation
func NewCreateProgram(ctx *middleware.Context, handler CreateProgramHandler) *CreateProgram {
	return &CreateProgram{Context: ctx, Handler: handler}
}

/*CreateProgram swagger:route POST /programs createProgram

Create program

*/
type CreateProgram struct {
	Context *middleware.Context
	Handler CreateProgramHandler
}

func (o *CreateProgram) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateProgramParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.JWTClaimBody
	if uprinc != nil {
		principal = uprinc.(*models.JWTClaimBody) // this is really a models.JWTClaimBody, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
