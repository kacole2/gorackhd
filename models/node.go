package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*node Node node

swagger:model node
*/
type Node struct {

	/* Catalogs catalogs
	 */
	Catalogs []*Catalog `json:"catalogs,omitempty"`

	/* ID id
	 */
	ID *string `json:"id,omitempty"`

	/* IPAddresses ip addresses
	 */
	IPAddresses interface{} `json:"ipAddresses,omitempty"`

	/* Name name

	Required: true
	*/
	Name string `json:"name,omitempty"`

	/* ObmSettings obm settings
	 */
	ObmSettings interface{} `json:"obmSettings,omitempty"`

	/* Profile profile

	Required: true
	*/
	Profile string `json:"profile,omitempty"`

	/* Sku sku
	 */
	Sku *Sku `json:"sku,omitempty"`

	/* Workflows workflows
	 */
	Workflows []*Graphobject `json:"workflows,omitempty"`
}

// Validate validates this node
func (m *Node) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCatalogs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProfile(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWorkflows(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Node) validateCatalogs(formats strfmt.Registry) error {

	if swag.IsZero(m.Catalogs) { // not required
		return nil
	}

	for i := 0; i < len(m.Catalogs); i++ {

		if m.Catalogs[i] != nil {

			if err := m.Catalogs[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Node) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateProfile(formats strfmt.Registry) error {

	if err := validate.Required("profile", "body", string(m.Profile)); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateWorkflows(formats strfmt.Registry) error {

	if swag.IsZero(m.Workflows) { // not required
		return nil
	}

	for i := 0; i < len(m.Workflows); i++ {

		if m.Workflows[i] != nil {

			if err := m.Workflows[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}