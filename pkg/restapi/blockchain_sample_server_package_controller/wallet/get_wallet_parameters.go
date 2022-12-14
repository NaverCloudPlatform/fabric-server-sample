// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetWalletParams creates a new GetWalletParams object
//
// There are no default values defined in the spec.
func NewGetWalletParams() GetWalletParams {

	return GetWalletParams{}
}

// GetWalletParams contains all the bound params for the get wallet operation
// typically these are obtained from a http.Request
//
// swagger:parameters getWallet
type GetWalletParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: query
	*/
	WalletID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetWalletParams() beforehand.
func (o *GetWalletParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qWalletID, qhkWalletID, _ := qs.GetOK("walletId")
	if err := o.bindWalletID(qWalletID, qhkWalletID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindWalletID binds and validates parameter WalletID from query.
func (o *GetWalletParams) bindWalletID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("walletId", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("walletId", "query", raw); err != nil {
		return err
	}
	o.WalletID = raw

	return nil
}
