package cmd

import (
	"github.com/Shopify/goose/logger"
	"github.com/Shopify/goose/safely"
	"github.com/covid-tracing-mongolia/backend-server/pkg/app"
	"github.com/covid-tracing-mongolia/backend-server/pkg/telemetry"
)

var log = logger.New("cmd")

func RunAndWait(appBuilder *app.AppBuilder) {
	defer safely.Recover() // panics -> bugsnag

	log(nil, nil).Info("starting")

	mainApp, db := appBuilder.Build()

	defer telemetry.Initialize(db).Cleanup()

	err := mainApp.RunAndWait()
	defer log(nil, err).Info("final message before shutdown")
}
