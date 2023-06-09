// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pjotrscholtze/go-monitoring/models"
)

// ListAllChecksOKCode is the HTTP code returned for type ListAllChecksOK
const ListAllChecksOKCode int = 200

/*ListAllChecksOK Successful operation

swagger:response listAllChecksOK
*/
type ListAllChecksOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Check `json:"body,omitempty"`
}

// NewListAllChecksOK creates ListAllChecksOK with default headers values
func NewListAllChecksOK() *ListAllChecksOK {

	return &ListAllChecksOK{}
}

// WithPayload adds the payload to the list all checks o k response
func (o *ListAllChecksOK) WithPayload(payload []*models.Check) *ListAllChecksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all checks o k response
func (o *ListAllChecksOK) SetPayload(payload []*models.Check) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllChecksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Check, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
