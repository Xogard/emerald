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

type Contrato struct {
	ID            uuid.UUID       `json:"id" db:"id"`
	CreatedAt     time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at" db:"updated_at"`
	ClienteID     uuid.UUID       `json:"clienteid" db:"clienteid"`
	Exercicio     int             `json:"exercicio" db:"exercicio"`
	Dtcontrato    time.Time       `json:"dtcontrato" db:"dtcontrato"`
	Duracao       int             `json:"duracao" db:"duracao"`
	Valor         decimal.Decimal `json:"valor" db:"valor"`
	Cancelado     bool            `json:"cancelado" db:"cancelado"`
	Faturas       Faturas         `has_many:"faturas" order_by:"dtvencimento" fk_id:"contratoid"`
	DiaVencimento int             `json:"diavencimento" db:"diavencimento"`
}

type DuracaoPermitida struct {
	DtContrato time.Time
	Value      int
}

// String is not required by pop and may be deleted
func (c Contrato) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Contratos is not required by pop and may be deleted
type Contratos []Contrato

// String is not required by pop and may be deleted
func (c Contratos) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Contrato) Validate(tx *pop.Connection) (*validate.Errors, error) {

	finaldate := FinalDate(c.Dtcontrato, c.Duracao)
	nextyear := NextYear(c.Dtcontrato)
	return validate.Validate(
		&validators.UUIDIsPresent{Field: c.ClienteID, Name: "ClienteID"},
		&validators.IntIsPresent{Field: c.Exercicio, Name: "Exercicio"},
		&validators.TimeIsPresent{Field: c.Dtcontrato, Name: "Dtcontrato"},
		&validators.IntIsPresent{Field: c.Duracao, Name: "Duracao"},
		&validators.TimeIsBeforeTime{FirstName: "Final Date of contract", FirstTime: finaldate, SecondName: "Next Year", SecondTime: nextyear},
		//&validators.IntIsPresent{Field: c.Valor, Name: "Valor"},
		//&validators.BytesArePresent{Field: c.Cancelado, Name: "Cancelado"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Contrato) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Contrato) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

//IsValid dasd
func NextYear(dt time.Time) time.Time {

	return time.Date(dt.Year()+1, time.January, 01, 00, 0, 0, 0, dt.Location())

}

func FinalDate(dt time.Time, duracao int) time.Time {
	return dt.AddDate(0, duracao, 0)
}
