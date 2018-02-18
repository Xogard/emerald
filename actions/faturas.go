package actions

import (
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gustavo/emerald/models"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Fatura)
// DB Table: Plural (faturas)
// Resource: Plural (Faturas)
// Path: Plural (/faturas)
// View Template Folder: Plural (/templates/faturas/)

// FaturasResource is the resource for the Fatura model
type FaturasResource struct {
	buffalo.Resource
}

// List gets all Faturas. This function is mapped to the path
// GET /faturas
func (v FaturasResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	faturas := &models.Faturas{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Faturas from the DB
	if err := q.All(faturas); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the headers so clients know how to paginate.
	c.Response().Header().Set("X-Pagination", q.Paginator.String())

	return c.Render(200, r.JSON(faturas))
}

// Show gets the data for one Fatura. This function is mapped to
// the path GET /faturas/{fatura_id}
func (v FaturasResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Fatura
	fatura := &models.Fatura{}

	// To find the Fatura the parameter fatura_id is used.
	if err := tx.Find(fatura, c.Param("fatura_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.JSON(fatura))
}

// New default implementation. Returns a 404
func (v FaturasResource) New(c buffalo.Context) error {
	return c.Error(404, errors.New("not available"))
}

// CreateFaturas ok
func (v FaturasResource) CreateFaturas(c buffalo.Context, contrato *models.Contrato) error {
	// Allocate an empty Fatura
	//fatura := &models.Fatura{}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	year := contrato.Dtcontrato.Year()
	month := contrato.Dtcontrato.Month()
	day := contrato.DiaVencimento
	location := contrato.Dtcontrato.Location()

	DtInicial := time.Date(year, month, day, 0, 0, 0, 0, location)
	for index := 0; index < contrato.Duracao; index++ {

		// Allocate an empty Fatura
		fatura := &models.Fatura{}

		fatura.ClienteID = contrato.ClienteID
		fatura.ContratoID = contrato.ID
		fatura.Valor = contrato.Valor
		fatura.DtVencimento = DtInicial.AddDate(0, index, 0)

		// Validate the data from the html form
		verrs, err := tx.ValidateAndCreate(fatura)
		if err != nil {
			return errors.WithMessage(err, "Erro")
		}

		if verrs.HasAny() {
			// Render errors as JSON
			return c.Render(400, r.JSON(verrs))
		}

	}
	erro := tx.Load(contrato, "Faturas")
	if erro != nil {
		return errors.WithMessage(erro, "Erro")
	}

	return c.Render(201, r.JSON(contrato))
}

// Create adds a Fatura to the DB. This function is mapped to the
// path POST /faturas
func (v FaturasResource) Create(c buffalo.Context) error {
	return c.Error(404, errors.New("not available"))
}

// Edit default implementation. Returns a 404
func (v FaturasResource) Edit(c buffalo.Context) error {
	return c.Error(404, errors.New("not available"))
}

// Update changes a Fatura in the DB. This function is mapped to
// the path PUT /faturas/{fatura_id}
func (v FaturasResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Fatura
	fatura := &models.Fatura{}

	if err := tx.Find(fatura, c.Param("fatura_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Fatura to the html form elements
	if err := c.Bind(fatura); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(fatura)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Render errors as JSON
		return c.Render(400, r.JSON(verrs))
	}

	return c.Render(200, r.JSON(fatura))
}

// Destroy deletes a Fatura from the DB. This function is mapped
// to the path DELETE /faturas/{fatura_id}
func (v FaturasResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Fatura
	fatura := &models.Fatura{}

	// To find the Fatura the parameter fatura_id is used.
	if err := tx.Find(fatura, c.Param("fatura_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(fatura); err != nil {
		return errors.WithStack(err)
	}

	return c.Render(200, r.JSON(fatura))
}