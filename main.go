package main

import (
	"os"
	"log"
)

func main() {
	cli := NewCli()

	err := cli.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
