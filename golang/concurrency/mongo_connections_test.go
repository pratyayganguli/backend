// all the code related to handling of mongo connections will be written over here...
// proper documentation should be done on the source code level itself...

package concurrency

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Customer struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func (c *Customer) CreateNewCustomer() {
	// this function will be creating a customer and then storing it in the database
}

// starter code for creating a single mongo connection.
func CreateSimpleMongoConnection() {
	// do not use `@` character in your password.
	uri := "mongodb://cidaas:cidaas%40000@localhost:27017/cidaas_db"
	options := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if client, err := mongo.Connect(options); err != nil {
		log.Printf("could not connect to the database! %s", err)
		return
	} else {
		if err := client.Ping(ctx, nil); err != nil {
			log.Fatalf("could not ping the mongo server! %s", err.Error())
		} else {
			log.Printf("connection successfull!")
		}
	}
}
