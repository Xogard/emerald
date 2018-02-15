package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type Cliente struct {
	ID           uuid.UUID `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	Nome         string    `json:"nome" db:"nome"`
	DtNascimento time.Time `json:"dtnascimento" db:"dtnascimento"`
	Cpf          string    `json:"cpf" db:"cpf"`
	Rg           string    `json:"rg" db:"rg"`
	OrgaoExp     string    `json:"orgaoexp" db:"orgaoexp"`
	Email        string    `json:"email" db:"email"`
	Fone         string    `json:"fone" db:"fone"`
	Celular      string    `json:"celular" db:"celular"`
	Endereco     Endereco  `has_one:"endereco" fk_id:"clienteid"`
	Filhos       Filhos    `has_many:"filhos"`
}

// String is not required by pop and may be deleted
func (c Cliente) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Clientes is not required by pop and may be deleted
type Clientes []Cliente

// String is not required by pop and may be deleted
func (c Clientes) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Cliente) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.Nome, Name: "Nome"}), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Cliente) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Cliente) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
