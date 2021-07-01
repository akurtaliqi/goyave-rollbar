package main

import (
	"os"

	"goyave.dev/template/http/route"
	_ "goyave.dev/template/http/validation"

	"time"

	"github.com/rollbar/rollbar-go"
	"goyave.dev/goyave/v3"
	// Import the appropriate GORM dialect for the database you're using.
	// _ "goyave.dev/goyave/v3/database/dialect/mysql"
	// _ "goyave.dev/goyave/v3/database/dialect/postgres"
	// _ "goyave.dev/goyave/v3/database/dialect/sqlite"
	// _ "goyave.dev/goyave/v3/database/dialect/mssql"
)

func main() {
	// This is the entry point of your application.
	if err := goyave.Start(route.Register); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}

	rollbar.SetToken("MY_TOKEN")
	rollbar.SetEnvironment("dev")                                         // defaults to "development"
	rollbar.SetCodeVersion("v2")                                          // optional Git hash/branch/tag (required for GitHub integration)
	rollbar.SetServerHost("web.1")                                        // optional override; defaults to hostname
	rollbar.SetServerRoot("https://github.com/akurtaliqi/goyave-rollbar") // path of project (required for GitHub integration and non-project stacktrace collapsing)

	rollbar.Info("Message body goes here")
	rollbar.WrapAndWait(doSomething)
}

func doSomething() {
	var timer *time.Timer = nil
	timer.Reset(10) // this will panic
}
