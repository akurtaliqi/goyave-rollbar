package main

import (
	"os"

	"goyave.dev/goyave/v3"
	"goyave.dev/template/http/route"
	_ "goyave.dev/template/http/validation"

	"github.com/rollbar/rollbar-go"
)

func main() {
	rollbar.SetToken("MY_TOKEN")
	rollbar.SetEnvironment("production")                          // defaults to "development"
	rollbar.SetCodeVersion("v2")                                  // optional Git hash/branch/tag (required for GitHub integration)
	rollbar.SetServerHost("web.1")                                // optional override; defaults to hostname
	rollbar.SetServerRoot("github.com/akurtaliqi/goyave-rollbar") // path of project (required for GitHub integration and non-project stacktrace collapsing)

	exitCode := 0
	if err := goyave.Start(route.Register); err != nil {
		rollbar.Critical(err)

		exitCode = err.(*goyave.Error).ExitCode
	}

	os.Exit(exitCode)
}
