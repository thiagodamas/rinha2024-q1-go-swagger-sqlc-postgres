// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ConsultarExtratoHandlerFunc turns a function with the right signature into a consultar extrato handler
type ConsultarExtratoHandlerFunc func(ConsultarExtratoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ConsultarExtratoHandlerFunc) Handle(params ConsultarExtratoParams) middleware.Responder {
	return fn(params)
}

// ConsultarExtratoHandler interface for that can handle valid consultar extrato params
type ConsultarExtratoHandler interface {
	Handle(ConsultarExtratoParams) middleware.Responder
}

// NewConsultarExtrato creates a new http.Handler for the consultar extrato operation
func NewConsultarExtrato(ctx *middleware.Context, handler ConsultarExtratoHandler) *ConsultarExtrato {
	return &ConsultarExtrato{Context: ctx, Handler: handler}
}

/*
	ConsultarExtrato swagger:route GET /clientes/{id}/extrato consultarExtrato

# Consultar extrato

Consultar extrato
*/
type ConsultarExtrato struct {
	Context *middleware.Context
	Handler ConsultarExtratoHandler
}

func (o *ConsultarExtrato) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewConsultarExtratoParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
