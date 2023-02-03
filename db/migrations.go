package db

import (
	"log"

	"github.com/juliotorresmoreno/SpecialistTalk/configs"
	"github.com/juliotorresmoreno/SpecialistTalk/model"
)

// Migrate s
func Migrate() {
	conf := configs.GetConfig()
	conn, err := NewEngigne(conf.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	err = conn.Sync2(model.User{})
	if err != nil {
		log.Fatal(err)
	}
}
