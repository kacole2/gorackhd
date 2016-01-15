package pollers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewGetPollersLibraryIdentifierParams creates a new GetPollersLibraryIdentifierParams object
// with the default values initialized.
func NewGetPollersLibraryIdentifierParams() *GetPollersLibraryIdentifierParams {
	return &GetPollersLibraryIdentifierParams{}
}

/*GetPollersLibraryIdentifierParams contains all the parameters to send to the API endpoint
for the get pollers library identifier operation typically these are written to a http.Request
*/
type GetPollersLibraryIdentifierParams struct {

	/*Identifier
	  library poller identifier


	*/
	Identifier string
}

// WithIdentifier adds the identifier to the get pollers library identifier params
func (o *GetPollersLibraryIdentifierParams) WithIdentifier(identifier string) *GetPollersLibraryIdentifierParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetPollersLibraryIdentifierParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}