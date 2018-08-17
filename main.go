package main

import (
	"os"
	"log"
)

func main() {
	var err error

	app := Application{}
	err = app.Bootstrap(os.Args)

	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
