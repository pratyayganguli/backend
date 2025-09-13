// all the code related to handling of mongo connections will be written over here...
// proper documentation should be done on the source code level itself...

package concurrency

import (
	"context"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const DATABASE = "cidaas"

type Customer struct {
	Id    uint   `bson:"id" json:"id"`
	Email string `bson:"name" json:"email"`
}

func (c *Customer) CreateNewCustomer() {
	// this function will be creating a customer and then storing it in the database
	client := CreateMongoConnectionPool()
	db := client.Database(DATABASE)
	if _, err := db.Collection("users").InsertOne(context.TODO(), c); err != nil {
		log.Fatalf("could not insert data! %s", err.Error())
	}
}

func TestCreateSimpleMongoConnection(t *testing.T) {
	CreateSimpleMongoConnection()
}

func TestCreateMongoConnectionPool(t *testing.T) {
	CreateMongoConnectionPool()
}

func TestCreateUser(t *testing.T) {
	c := &Customer{
		Id:    1,
		Email: "alex@securemail.com",
	}
	c.CreateNewCustomer()
}

// starter code for creating a single mongo connection.
func CreateSimpleMongoConnection() {
	// do not use `@` character in your password.
	uri := "mongodb://developer:cidaas000@localhost:27017/cidaas?authSource=admin"
	options := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if client, err := mongo.Connect(options); err != nil {
		log.Fatalf("could not connect to the database! %s", err)
	} else {
		if err := client.Ping(ctx, nil); err != nil {
			log.Fatalf("could not ping the mongo server! %s", err.Error())
		} else {
			log.Printf("connection successfull!")
		}
	}
}

func CreateMongoConnectionPool() *mongo.Client {
	uri := "mongodb://developer:cidaas000@localhost:27017/cidaas?authSource=admin"
	clientOpts := options.Client().ApplyURI(uri).SetMaxPoolSize(1000).SetMinPoolSize(10).SetMaxConnIdleTime(2 * time.Minute)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if client, err := mongo.Connect(clientOpts); err != nil {
		// handle
		log.Fatalf("could not connect to database! %s", err.Error())
		return nil
	} else {
		if e := client.Ping(ctx, nil); e != nil {
			// handle ping call
			log.Fatalf("could not ping database! %s", e.Error())
			return nil
		} else {
			log.Printf("connection successfull!")
			return client
		}
	}
}
