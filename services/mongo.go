package services

import (
	"context"

	"github.com/juliotorresmoreno/SpecialistTalk/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var poolMongo *mongo.Client

func GetPoolMongo() (*mongo.Client, error) {
	var err error
	if poolMongo == nil {
		conf := configs.GetConfig()
		uri := conf.Mongo.DSN

		ctx := context.Background()
		opts := options.Client().ApplyURI(uri)
		poolMongo, err = mongo.Connect(ctx, opts)
	}

	return poolMongo, err
}
