// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetTargetChecksParams creates a new GetTargetChecksParams object
//
// There are no default values defined in the spec.
func NewGetTargetChecksParams() GetTargetChecksParams {

	return GetTargetChecksParams{}
}

// GetTargetChecksParams contains all the bound params for the get target checks operation
// typically these are obtained from a http.Request
//
// swagger:parameters getTargetChecks
type GetTargetChecksParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Name of the check
	  Required: true
	  In: path
	*/
	CheckName string
	/*Name of the target
	  Required: true
	  In: path
	*/
	TargetName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetTargetChecksParams() beforehand.
func (o *GetTargetChecksParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rCheckName, rhkCheckName, _ := route.Params.GetOK("checkName")
	if err := o.bindCheckName(rCheckName, rhkCheckName, route.Formats); err != nil {
		res = append(res, err)
	}

	rTargetName, rhkTargetName, _ := route.Params.GetOK("targetName")
	if err := o.bindTargetName(rTargetName, rhkTargetName, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCheckName binds and validates parameter CheckName from path.
func (o *GetTargetChecksParams) bindCheckName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.CheckName = raw

	return nil
}

// bindTargetName binds and validates parameter TargetName from path.
func (o *GetTargetChecksParams) bindTargetName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.TargetName = raw

	return nil
}
