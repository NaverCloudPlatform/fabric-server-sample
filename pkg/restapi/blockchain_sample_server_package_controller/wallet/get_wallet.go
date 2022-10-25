// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetWalletHandlerFunc turns a function with the right signature into a get wallet handler
type GetWalletHandlerFunc func(GetWalletParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetWalletHandlerFunc) Handle(params GetWalletParams) middleware.Responder {
	return fn(params)
}

// GetWalletHandler interface for that can handle valid get wallet params
type GetWalletHandler interface {
	Handle(GetWalletParams) middleware.Responder
}

// NewGetWallet creates a new http.Handler for the get wallet operation
func NewGetWallet(ctx *middleware.Context, handler GetWalletHandler) *GetWallet {
	return &GetWallet{Context: ctx, Handler: handler}
}

/* GetWallet swagger:route GET /getWallet Wallet getWallet

GetWallet get wallet API

*/
type GetWallet struct {
	Context *middleware.Context
	Handler GetWalletHandler
}

func (o *GetWallet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetWalletParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetWalletOKBody get wallet o k body
//
// swagger:model GetWalletOKBody
type GetWalletOKBody struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// token
	Token int64 `json:"token"`
}

// Validate validates this get wallet o k body
func (o *GetWalletOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get wallet o k body based on context it is used
func (o *GetWalletOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWalletOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWalletOKBody) UnmarshalBinary(b []byte) error {
	var res GetWalletOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}