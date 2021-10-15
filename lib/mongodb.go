package lib

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// var Db mongo.Database

func MongodbConnection() mongo.Database {

	MONGO_URL := viper.GetString("mongodb.url")
	MONGO_PORT := viper.GetString("mongodb.port")
	MONGO_USENAME := viper.GetString("mongodb.username")
	MONGO_PASSWORD := viper.GetString("mongodb.password")
	MONGO_DBNAME := viper.GetString("mongodb.dbname")

	url := "mongodb://" + MONGO_USENAME + ":" + MONGO_PASSWORD + "@" + MONGO_URL + ":" + MONGO_PORT + "/" + MONGO_DBNAME

	client, err := mongo.NewClient(options.Client().ApplyURI(url).SetMaxPoolSize(2000))

	if err != nil {

		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {

		fmt.Println("could not ping to mongo db service: \n", err)

		defer func() {

			ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

			err := client.Disconnect(ctx)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Mongo disconnected")
		}()
	}

	db := *client.Database(MONGO_DBNAME)

	fmt.Println("Mongo connected")

	return db
}
