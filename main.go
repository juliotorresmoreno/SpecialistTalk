package main

import (
	"github.com/juliotorresmoreno/freelive/configs"
	"github.com/juliotorresmoreno/freelive/db"
	"github.com/juliotorresmoreno/freelive/server"
)

func main() {

	configs.Init()
	db.Migrate()

	e := server.NewServer()

	e.Logger.Fatal(e.Listen())
}
