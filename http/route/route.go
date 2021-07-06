package route

import (
	"fmt"
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

	router.StatusHandler(func(response *goyave.Response, request *goyave.Request) {
		rollbar.LogPanic(response.GetError(), true)
		// rollbar.Info(response.GetError())
		fmt.Print("here")
		fmt.Print(response.GetError())
		fmt.Print(response.GetStatus())
		fmt.Println(request.ContentLength())
		fmt.Println(request.Extra)
		fmt.Println(request.URI().Path)
		fmt.Print("there")
		rollbar.Warning("Bad request", map[string]interface{}{
			"<nil>4040": "yo",
		})
		switch response.GetStatus() {
		case 401:
			rollbar.Warning("Warning unauthorized" + response.GetStacktrace())
		case 403:
			rollbar.Error("Error " + response.GetStacktrace())
		case 500:
			rollbar.Error("Error")
			// rollbar.Message("Error route"+request.URI().Path, response.GetStacktrace())
		}
		goyave.PanicStatusHandler(response, request)
	}, http.StatusNotFound, http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized)

	/*router.StatusHandler(func(response *goyave.Response, request *goyave.Request) {
		//rollbar.Critical(response.GetError())
		// rollbar.ErrorWithExtras("critical", &goyave.Error{}, request.Extra)
		rollbar.LogPanic(response.GetError(), true)
		fmt.Println(request.ContentLength())
		fmt.Println(request.Extra)
		fmt.Println(request.URI().Path)
		if response.GetError() != nil && response.GetStatus() == 500 {
			rollbar.Critical(response.GetError())

			fmt.Printf("response.GetError(): %v\n", response.GetError())
			fmt.Printf("response.GetStacktrace(): %v\n", response.GetStacktrace())
			goyave.ErrLogger.Println(response.GetError())
			fmt.Printf("response.GetStatus(): %v\n", response.GetStatus())
		}
		/*rollbar.Critical("critical", map[string]interface{}{
			"hello": "critical",
		})

		rollbar.Info("info", map[string]interface{}{
			"hello": "info",
		})
		goyave.PanicStatusHandler(response, request)
	}, http.StatusInternalServerError)

	router.StatusHandler(func(response *goyave.Response, request *goyave.Request) {
		response.GetError()
		rollbar.LogPanic(response.GetError(), true)
		rollbar.Error(response.GetError())
		fmt.Printf("response.GetError(): %v\n", response.GetError())
		fmt.Printf("response.GetStacktrace(): %v\n", response.GetStacktrace())
		fmt.Printf("response.GetStatus(): %v\n", response.GetStatus())

		goyave.PanicStatusHandler(response, request)
	}, http.StatusBadRequest)

	router.StatusHandler(func(response *goyave.Response, request *goyave.Request) {
		response.GetError()
		rollbar.Info(response.GetError())
		fmt.Printf("response.GetError(): %v\n", response.GetError())
		fmt.Printf("response.GetStacktrace(): %v\n", response.GetStacktrace())
		fmt.Printf("response.GetStatus(): %v\n", response.GetStatus())

		goyave.PanicStatusHandler(response, request)
	}, http.StatusNotFound)

	router.StatusHandler(func(response *goyave.Response, request *goyave.Request) {
		rollbar.Warning("unhautorized")
		goyave.PanicStatusHandler(response, request)
	}, http.StatusUnauthorized)*/

	/*router.Get("/auth", func(response *goyave.Response, r *goyave.Request) {
		response.String(http.StatusUnauthorized, "KO")
	}).Middleware(middleware.RollbarMiddleware)*/

}
