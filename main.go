package main

import (
	"goip/app"
	"log"
	"os"
)

func main() {
	cli := app.New()
	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
