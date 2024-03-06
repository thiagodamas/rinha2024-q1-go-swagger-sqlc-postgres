// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TransacaoInput transacao input
//
// swagger:model TransacaoInput
type TransacaoInput struct {

	// descricao
	// Required: true
	// Max Length: 10
	// Min Length: 1
	Descricao *string `json:"descricao"`

	// tipo
	// Required: true
	// Enum: [c d]
	Tipo *string `json:"tipo"`

	// valor
	// Required: true
	Valor *int64 `json:"valor"`
}

// Validate validates this transacao input
func (m *TransacaoInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescricao(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTipo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransacaoInput) validateDescricao(formats strfmt.Registry) error {

	if err := validate.Required("descricao", "body", m.Descricao); err != nil {
		return err
	}

	if err := validate.MinLength("descricao", "body", *m.Descricao, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("descricao", "body", *m.Descricao, 10); err != nil {
		return err
	}

	return nil
}

var transacaoInputTypeTipoPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["c","d"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transacaoInputTypeTipoPropEnum = append(transacaoInputTypeTipoPropEnum, v)
	}
}

const (

	// TransacaoInputTipoC captures enum value "c"
	TransacaoInputTipoC string = "c"

	// TransacaoInputTipoD captures enum value "d"
	TransacaoInputTipoD string = "d"
)

// prop value enum
func (m *TransacaoInput) validateTipoEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transacaoInputTypeTipoPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TransacaoInput) validateTipo(formats strfmt.Registry) error {

	if err := validate.Required("tipo", "body", m.Tipo); err != nil {
		return err
	}

	// value enum
	if err := m.validateTipoEnum("tipo", "body", *m.Tipo); err != nil {
		return err
	}

	return nil
}

func (m *TransacaoInput) validateValor(formats strfmt.Registry) error {

	if err := validate.Required("valor", "body", m.Valor); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this transacao input based on context it is used
func (m *TransacaoInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TransacaoInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransacaoInput) UnmarshalBinary(b []byte) error {
	var res TransacaoInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
