package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"

	"github.com/gobuffalo/x/sessions"
	"github.com/gustavo/emerald/models"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Addr:         "0.0.0.0:3000",
			Env:          ENV,
			SessionStore: sessions.Null{},
			SessionName:  "_emerald_session",
		})
		// Automatically redirect to SSL
		app.Use(ssl.ForceSSL(secure.Options{
			SSLRedirect:     ENV == "production",
			SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		}))

		// Set the request content type to JSON
		app.Use(middleware.SetContentType("application/json"))

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.PopTransaction)
		// Remove to disable this.
		app.Use(middleware.PopTransaction(models.DB))

		api := app.Group("/api")
		clientes := api.Resource("/clientes", ClientesResource{})
		clientes.Resource("/filhos", FilhosResource{})
		clientes.Resource("/contratos", ContratosResource{})
		api.Resource("/escolas", EscolasResource{})

		app.GET("/", HomeHandler)

		//app.Resource("/faturas", FaturasResource{})
		app.Resource("/pagamentos", PagamentosResource{})
	}

	return app
}
