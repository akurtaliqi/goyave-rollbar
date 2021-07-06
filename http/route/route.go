package route

import (
	"net/http"

	"github.com/rollbar/rollbar-go"
	"goyave.dev/template/http/controller/hello"
	"goyave.dev/template/http/middleware"

	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/auth"
	"goyave.dev/goyave/v3/cors"
)

// Routing is an essential part of any Goyave application.
// Routes definition is the action of associating a URI, sometimes having
// parameters, with a handler which will process the request and respond to it.

// Routes are defined in routes registrer functions.
// The main route registrer is passed to "goyave.Start()" and is executed
// automatically with a newly created root-level router.

// Register all the application routes. This is the main route registrer.
func Register(router *goyave.Router) {

	// Applying default CORS settings (allow all methods and all origins)
	// Learn more about CORS options here: https://goyave.dev/guide/advanced/cors.html
	router.CORS(cors.Default())

	// Register your routes here

	// Route without validation
	router.Get("/hello", hello.SayHi)

	// Route with validation
	router.Middleware(middleware.RollbarMiddleware, auth.ConfigBasicAuth())

	router.Get("/dummy", hello.DummyModel)

	router.Get("/panic", hello.PanickyFunction)

	router.StatusHandler(func(response *goyave.Response, request *goyave.Request) {

		rollbar.Warning("Bad request", map[string]interface{}{
			"<nil>4040": response.GetStacktrace(),
		})

		switch response.GetStatus() {
		case 401:
			rollbar.Warning("Warning unauthorized 2" + response.GetStacktrace())
		case 403:
			rollbar.Error("Error "+request.URI().Path, response.GetStacktrace())
		case 500:
			rollbar.Critical("Critical ", response.GetStacktrace())
			// rollbar.Message("Error route"+request.URI().Path, response.GetStacktrace())
		}
		goyave.PanicStatusHandler(response, request)
	}, http.StatusNotFound, http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden)

}
