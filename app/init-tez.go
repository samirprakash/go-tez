package main

import (
	"log"
	"os"

	"github.com/samirprakash/tez"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	t := &tez.Tez{}
	err = t.New(path)
	if err != nil {
		log.Fatal(err)
	}

	t.AppName = "app"
	t.Debug = true

	app := &application{
		App: t,
	}

	return app
}
