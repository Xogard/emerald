package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type Fatura struct {
	ID           uuid.UUID       `json:"id" db:"id"`
	CreatedAt    time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at" db:"updated_at"`
	ClienteID    uuid.UUID       `json:"clienteid" db:"clienteid"`
	ContratoID   uuid.UUID       `json:"contratoid" db:"contratoid"`
	PagamentoID  uuid.UUID       `json:"pagamentoid" db:"pagamentoid"`
	Valor        decimal.Decimal `json:"valor" db:"valor"`
	DtVencimento time.Time       `json:"dtvencimento" db:"dtvencimento"`
	DtPagamento  time.Time       `json:"dtpagamento" db:"dtpagamento"`
}

// String is not required by pop and may be deleted
func (f Fatura) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Faturas is not required by pop and may be deleted
type Faturas []Fatura

// String is not required by pop and may be deleted
func (f Faturas) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (f *Fatura) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (f *Fatura) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.UUIDIsPresent{Field: f.ClienteID, Name: "ClienteID"},
		&validators.UUIDIsPresent{Field: f.ContratoID, Name: "ContratoID"},
		&validators.TimeAfterTime{FirstName: "Due Date", FirstTime: f.DtVencimento, SecondName: "Today", SecondTime: time.Now()},
	), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (f *Fatura) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
