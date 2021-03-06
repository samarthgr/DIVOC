// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// EnrollRecipientOKCode is the HTTP code returned for type EnrollRecipientOK
const EnrollRecipientOKCode int = 200

/*EnrollRecipientOK OK

swagger:response enrollRecipientOK
*/
type EnrollRecipientOK struct {
}

// NewEnrollRecipientOK creates EnrollRecipientOK with default headers values
func NewEnrollRecipientOK() *EnrollRecipientOK {

	return &EnrollRecipientOK{}
}

// WriteResponse to the client
func (o *EnrollRecipientOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// EnrollRecipientUnauthorizedCode is the HTTP code returned for type EnrollRecipientUnauthorized
const EnrollRecipientUnauthorizedCode int = 401

/*EnrollRecipientUnauthorized Invalid token

swagger:response enrollRecipientUnauthorized
*/
type EnrollRecipientUnauthorized struct {
}

// NewEnrollRecipientUnauthorized creates EnrollRecipientUnauthorized with default headers values
func NewEnrollRecipientUnauthorized() *EnrollRecipientUnauthorized {

	return &EnrollRecipientUnauthorized{}
}

// WriteResponse to the client
func (o *EnrollRecipientUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
