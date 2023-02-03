package main

import (
	"github.com/juliotorresmoreno/SpecialistTalk/db"
	"github.com/juliotorresmoreno/SpecialistTalk/server"
)

func main() {
	db.Migrate()

	e := server.NewServer()

	e.Logger.Fatal(e.Listen())
}
