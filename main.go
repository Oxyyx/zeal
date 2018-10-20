package main

import (
	"git.oxy.sh/oxy/zeal/cmd"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Czar"
	app.Usage = "Go-driven e-commerce"
	app.Commands = []cli.Command{
		cmd.Web,
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(4, "Failed to run app with %s: %v", os.Args, err)
	}
}