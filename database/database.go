package database

import (
	"context"

	"global-auth/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseConnect() (mongo.Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Env("MONGO_DB_URL")))
	if err != nil {
		panic(err)
	}
	db := client.Database("momofin_dev")
	return *db, nil
}
