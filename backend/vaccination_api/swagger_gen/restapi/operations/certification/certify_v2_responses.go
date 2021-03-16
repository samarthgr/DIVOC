// Code generated by go-swagger; DO NOT EDIT.

package certification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// CertifyV2OKCode is the HTTP code returned for type CertifyV2OK
const CertifyV2OKCode int = 200

/*CertifyV2OK OK

swagger:response certifyV2OK
*/
type CertifyV2OK struct {
}

// NewCertifyV2OK creates CertifyV2OK with default headers values
func NewCertifyV2OK() *CertifyV2OK {

	return &CertifyV2OK{}
}

// WriteResponse to the client
func (o *CertifyV2OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// CertifyV2BadRequestCode is the HTTP code returned for type CertifyV2BadRequest
const CertifyV2BadRequestCode int = 400

/*CertifyV2BadRequest Invalid input

swagger:response certifyV2BadRequest
*/
type CertifyV2BadRequest struct {

	/*
	  In: Body
	*/
	Payload *CertifyV2BadRequestBody `json:"body,omitempty"`
}

// NewCertifyV2BadRequest creates CertifyV2BadRequest with default headers values
func NewCertifyV2BadRequest() *CertifyV2BadRequest {

	return &CertifyV2BadRequest{}
}

// WithPayload adds the payload to the certify v2 bad request response
func (o *CertifyV2BadRequest) WithPayload(payload *CertifyV2BadRequestBody) *CertifyV2BadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the certify v2 bad request response
func (o *CertifyV2BadRequest) SetPayload(payload *CertifyV2BadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CertifyV2BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}