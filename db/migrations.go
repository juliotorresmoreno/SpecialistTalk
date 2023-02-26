package db

import (
	"fmt"
	"log"
	"strings"

	"github.com/juliotorresmoreno/SpecialistTalk/configs"
	"github.com/juliotorresmoreno/SpecialistTalk/model"
	"xorm.io/xorm"
)

func createIndex(conn *xorm.Engine, alg, indexname, table, columnname string) {
	sql := fmt.Sprintf(
		"CREATE INDEX \"%v\" ON public.%v USING %v (lower(%v));",
		indexname, table, alg, columnname,
	)
	_, _ = conn.Query(sql)
}

func configureUsers(conn *xorm.Engine) {
	err := conn.Sync2(&model.User{})
	if err != nil && !strings.HasPrefix(err.Error(), "Unknown col lower") {
		log.Fatal(err)
	}

	createIndex(conn, "IDX_users_name", "GIN", "users", "name")
	createIndex(conn, "IDX_users_lastname", "GIN", "users", "lastname")
}

func Migrate() {
	conf := configs.GetConfig()
	conn, err := NewEngigne(conf.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	configureUsers(conn)

	if err = conn.Sync2(&model.Chat{}); err != nil {
		log.Fatal(err)
	}
	if err = conn.Sync2(&model.Group{}); err != nil {
		log.Fatal(err)
	}
	if err = conn.Sync2(&model.Gallery{}); err != nil {
		log.Fatal(err)
	}
}
