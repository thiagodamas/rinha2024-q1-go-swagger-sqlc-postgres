// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TransacaoOutput transacao output
//
// swagger:model TransacaoOutput
type TransacaoOutput struct {

	// limite
	// Required: true
	Limite *int64 `json:"limite"`

	// saldo
	// Required: true
	Saldo *int64 `json:"saldo"`
}

// Validate validates this transacao output
func (m *TransacaoOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSaldo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransacaoOutput) validateLimite(formats strfmt.Registry) error {

	if err := validate.Required("limite", "body", m.Limite); err != nil {
		return err
	}

	return nil
}

func (m *TransacaoOutput) validateSaldo(formats strfmt.Registry) error {

	if err := validate.Required("saldo", "body", m.Saldo); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this transacao output based on context it is used
func (m *TransacaoOutput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TransacaoOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransacaoOutput) UnmarshalBinary(b []byte) error {
	var res TransacaoOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
