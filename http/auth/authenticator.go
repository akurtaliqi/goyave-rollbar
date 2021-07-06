package auth

import (
	"net/http"
	"reflect"

	"goyave.dev/goyave/v3"
)

// Column matches a column name with a struct field.
type Column struct {
	Field *reflect.StructField
	Name  string
}

// Authenticator is an object in charge of authenticating a model.
type Authenticator interface {

	// Authenticate fetch the user corresponding to the credentials
	// found in the given request and puts the result in the given user pointer.
	// If no user can be authenticated, returns the error detailing why the
	// authentication failed. The error message is already localized.
	Authenticate(request *goyave.Request, user interface{}) error
}

// Middleware create a new authenticator middleware to authenticate
// the given model using the given authenticator.
func Middleware(model interface{}, authenticator Authenticator) goyave.Middleware {
	return func(next goyave.Handler) goyave.Handler {
		return func(response *goyave.Response, r *goyave.Request) {
			userType := reflect.Indirect(reflect.ValueOf(model)).Type()
			user := reflect.New(userType).Interface()
			r.User = user
			if err := authenticator.Authenticate(r, r.User); err != nil {
				response.JSON(http.StatusUnauthorized, map[string]string{"authError": err.Error()})
				return
			}
			next(response, r)
		}
	}
}
