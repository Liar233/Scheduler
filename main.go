package main

import (
	"os"
	"log"
)

func main() {
	var config AppConfig
	cli := NewCli(&config)

	err := cli.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
