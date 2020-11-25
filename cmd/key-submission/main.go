package main

import (
	"github.com/covid-tracing-mongolia/backend-server/pkg/app"
	"github.com/covid-tracing-mongolia/backend-server/pkg/cmd"
)

func main() {
	cmd.RunAndWait(
		app.NewBuilder().
			WithTestTools().
			WithSubmission())
}
