package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type Filho struct {
	ID           uuid.UUID `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	Nome         string    `json:"nome" db:"nome"`
	DtNascimento time.Time `json:"dtnascimento" db:"dtnascimento"`
	Escola       Escola    `belongs_to:"Escola"`
	EscolaID     uuid.UUID `json:"escolaid" db:"escolaid"`
	Cliente      Cliente   `belongs_to:"Cliente"`
	ClienteID    uuid.UUID `json:"clienteid" db:"clienteid"`
}

// String is not required by pop and may be deleted
func (f Filho) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Filhos is not required by pop and may be deleted
type Filhos []Filho

// String is not required by pop and may be deleted
func (f Filhos) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (f *Filho) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.UUIDIsPresent{Field: f.ClienteID, Name: "ClienteID"},
		&validators.StringIsPresent{Field: f.Nome, Name: "Nome"},
		&validators.TimeIsPresent{Field: f.DtNascimento, Name: "DtNascimento"},
		&validators.UUIDIsPresent{Field: f.EscolaID, Name: "EscolaID"}), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (f *Filho) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (f *Filho) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
