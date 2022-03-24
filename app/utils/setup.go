package utils

import (
	"fmt"
	"go-pos-store/configs"
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// for db connection
func Init() {

	uri := configs.GetEnvVAlues("MongoDbHost")

	err := mgm.SetDefaultConfig(nil, "pos-store", options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to DB.........")
}
