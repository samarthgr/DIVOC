// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostNotificationOKCode is the HTTP code returned for type PostNotificationOK
const PostNotificationOKCode int = 200

/*PostNotificationOK OK

swagger:response postNotificationOK
*/
type PostNotificationOK struct {
}

// NewPostNotificationOK creates PostNotificationOK with default headers values
func NewPostNotificationOK() *PostNotificationOK {

	return &PostNotificationOK{}
}

// WriteResponse to the client
func (o *PostNotificationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
