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

	_, _ = conn.Query("CREATE INDEX \"IDX_users_name\" ON public.users USING GIN (name);")
	_, _ = conn.Query("CREATE INDEX \"IDX_users_lastname\" ON public.users USING GIN (lastname);")

	err = conn.Sync2(model.Chat{})
	if err != nil {
		log.Fatal(err)
	}
}
