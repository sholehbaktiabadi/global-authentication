package main

import (
	"global-auth/api"
	"global-auth/config"
	"global-auth/database"
	"net/http"
)

func main() {
	db, err := database.DatabaseConnect()
	if err != nil {
		panic("error database coonection")
	}
	var (
		api    = api.NewRouterApp(db)
		server = http.Server{
			Addr:    ":" + config.Env("SERVER_PORT"),
			Handler: api,
		}
	)
	server.ListenAndServe()
}
