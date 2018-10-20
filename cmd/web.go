package cmd

import (
	"github.com/gorilla/mux"
	"git.oxy.sh/oxy/zeal/models"
	"git.oxy.sh/oxy/zeal/routes"
	"github.com/urfave/cli"
	"log"
	"net/http"
)

var Web = cli.Command{
	Name:  "web",
	Usage: "Start Czar web server",
	Description: "Czar server application.",
	Action: runWeb,
	Flags: nil,
}

func runWeb(ctx *cli.Context) error {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	var err error
	err = models.InitDB()

	if err != nil {
		log.Fatal(4, "Failed to start server: %v", err)
		return err
	}

	err = http.ListenAndServe("localhost:8080", r)

	if err != nil {
		log.Fatal(4, "Failed to start server: %v", err)
		return err
	}

	return nil
}