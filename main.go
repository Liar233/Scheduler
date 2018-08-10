package main

import (
	"os"
	"log"
	"fmt"
)


func main() {
	var config AppConfig
	cli := NewCli(&config)

	err := cli.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", config)
}
