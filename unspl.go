package main

import (
	"log"
	"os"
	"unspl/cmd"
)

func main() {
	app := cmd.App()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
