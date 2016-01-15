package obm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewGetNodesIdentifierObmIdentifyParams creates a new GetNodesIdentifierObmIdentifyParams object
// with the default values initialized.
func NewGetNodesIdentifierObmIdentifyParams() *GetNodesIdentifierObmIdentifyParams {
	return &GetNodesIdentifierObmIdentifyParams{}
}

/*GetNodesIdentifierObmIdentifyParams contains all the parameters to send to the API endpoint
for the get nodes identifier obm identify operation typically these are written to a http.Request
*/
type GetNodesIdentifierObmIdentifyParams struct {

	/*Identifier
	  Mac addresses and unique aliases to identify the node by, |
	expect a string or an array of strings.


	*/
	Identifier string
}

// WithIdentifier adds the identifier to the get nodes identifier obm identify params
func (o *GetNodesIdentifierObmIdentifyParams) WithIdentifier(identifier string) *GetNodesIdentifierObmIdentifyParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodesIdentifierObmIdentifyParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

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