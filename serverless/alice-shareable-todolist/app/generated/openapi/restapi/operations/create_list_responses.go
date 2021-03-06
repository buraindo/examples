// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/yandex-cloud/examples/serverless/alice-shareable-todolist/app/generated/openapi/models"
)

// CreateListOKCode is the HTTP code returned for type CreateListOK
const CreateListOKCode int = 200

/*CreateListOK Create new list

swagger:response createListOK
*/
type CreateListOK struct {

	/*
	  In: Body
	*/
	Payload *CreateListOKBody `json:"body,omitempty"`
}

// NewCreateListOK creates CreateListOK with default headers values
func NewCreateListOK() *CreateListOK {

	return &CreateListOK{}
}

// WithPayload adds the payload to the create list o k response
func (o *CreateListOK) WithPayload(payload *CreateListOKBody) *CreateListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create list o k response
func (o *CreateListOK) SetPayload(payload *CreateListOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateListDefault error

swagger:response createListDefault
*/
type CreateListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateListDefault creates CreateListDefault with default headers values
func NewCreateListDefault(code int) *CreateListDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create list default response
func (o *CreateListDefault) WithStatusCode(code int) *CreateListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create list default response
func (o *CreateListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create list default response
func (o *CreateListDefault) WithPayload(payload *models.Error) *CreateListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create list default response
func (o *CreateListDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
