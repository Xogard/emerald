package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type Pagamento struct {
	ID          uuid.UUID       `json:"id" db:"id"`
	CreatedAt   time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at" db:"updated_at"`
	ClienteID   uuid.UUID       `json:"clienteid" db:"clienteid"`
	ContratoID  uuid.UUID       `json:"contratoid" db:"contratoid"`
	Valor       decimal.Decimal `json:"valor" db:"valor"`
	DtPagamento time.Time       `json:"dtpagamento" db:"dtpagamento"`
	Obs         string          `json:"obs" db:"observacao"`
	Faturas     []uuid.UUID     `json:"faturas,string" db:"-"`
}

// String is not required by pop and may be deleted
func (p Pagamento) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Pagamentos is not required by pop and may be deleted
type Pagamentos []Pagamento

// String is not required by pop and may be deleted
func (p Pagamentos) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Pagamento) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Pagamento) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Pagamento) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
