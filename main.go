package main

import (
	"time"
	"log"
	"context"
)

func main() {
	app := NewApplication()

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	<-app.Done()
}
