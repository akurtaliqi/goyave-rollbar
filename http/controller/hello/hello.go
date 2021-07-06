package hello

import (
	"errors"
	"net/http"

	"github.com/rollbar/rollbar-go"
	"goyave.dev/goyave/v3"
	"goyave.dev/template/database/model"
)

// Controllers are files containing a collection of Handlers related to a specific feature.
// Each feature should have its own package.
//
// Controller handlers contain the business logic of your application.
// They should be concise and focused on what matters for this particular feature in your application.
// Learn more about controllers here: https://goyave.dev/guide/basics/controllers.html

// ----------------------------------------------------------------------

// SayHi is a controller handler writing "Hi!" as a response.
//
// The Response object is used to write your response.
// https://goyave.dev/guide/basics/responses.html
//
// The Request object contains all the information about the incoming request, including it's parsed body,
// query params and route parameters.
// https://goyave.dev/guide/basics/requests.html
func SayHi(response *goyave.Response, request *goyave.Request) {
	response.String(http.StatusOK, "Hi!")
}

func DummyModel(response *goyave.Response, request *goyave.Request) {
	var u *model.User
	u.Name = "fake name"
}

// Echo is a controller handler writing the input field "text" as a response.
// This route is validated. See "http/request/echorequest/echo.go" for more details.
func Echo(response *goyave.Response, request *goyave.Request) {
	response.String(http.StatusOK, request.String("text"))
}

func PanickyFunction(response *goyave.Response, request *goyave.Request) {
	defer func() {
		err := recover()
		rollbar.LogPanic(err, true) // bool argument sets wait behavior
	}()

	panic(errors.New("critical error "))
}
