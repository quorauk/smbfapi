// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/heshoots/smbfapi/gen/models"
)

// SearchEventsOKCode is the HTTP code returned for type SearchEventsOK
const SearchEventsOKCode int = 200

/*SearchEventsOK search results matching criteria

swagger:response searchEventsOK
*/
type SearchEventsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Event `json:"body,omitempty"`
}

// NewSearchEventsOK creates SearchEventsOK with default headers values
func NewSearchEventsOK() *SearchEventsOK {

	return &SearchEventsOK{}
}

// WithPayload adds the payload to the search events o k response
func (o *SearchEventsOK) WithPayload(payload []*models.Event) *SearchEventsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search events o k response
func (o *SearchEventsOK) SetPayload(payload []*models.Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchEventsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Event, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// SearchEventsBadRequestCode is the HTTP code returned for type SearchEventsBadRequest
const SearchEventsBadRequestCode int = 400

/*SearchEventsBadRequest bad input parameter

swagger:response searchEventsBadRequest
*/
type SearchEventsBadRequest struct {
}

// NewSearchEventsBadRequest creates SearchEventsBadRequest with default headers values
func NewSearchEventsBadRequest() *SearchEventsBadRequest {

	return &SearchEventsBadRequest{}
}

// WriteResponse to the client
func (o *SearchEventsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}