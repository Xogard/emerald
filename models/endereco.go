package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/pop/nulls"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type Endereco struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
	Enderecoid  uuid.UUID    `json:"enderecoid" db:"enderecoid"`
	Logradouro  nulls.String `json:"logradouro" db:"logradouro"`
	Complemento string       `json:"complemento" db:"complemento"`
	Cep         string       `json:"cep" db:"cep"`
	Bairro      string       `json:"bairro" db:"bairro"`
	Cidade      string       `json:"cidade" db:"cidade"`
	Estado      string       `json:"estado" db:"estado"`
	ClienteID   uuid.UUID    `db:"clienteid"`
}

// String is not required by pop and may be deleted
func (e Endereco) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Enderecos is not required by pop and may be deleted
type Enderecos []Endereco

// String is not required by pop and may be deleted
func (e Enderecos) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (e *Endereco) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: e.Complemento, Name: "Complemento"},
		&validators.StringIsPresent{Field: e.Cep, Name: "Cep"},
		&validators.StringIsPresent{Field: e.Bairro, Name: "Bairro"},
		&validators.StringIsPresent{Field: e.Cidade, Name: "Cidade"},
		&validators.StringIsPresent{Field: e.Estado, Name: "Estado"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (e *Endereco) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (e *Endereco) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
