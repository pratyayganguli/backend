package mongo

import "testing"

// info: each module should have only one test file and the functions which are exported will be invoked in that particular module

func TestCreateSimpleMongoConnection(t *testing.T) {
	CreateSimpleMongoConnection()
}

func TestCreateMongoConnectionPool(t *testing.T) {
	CreateMongoConnectionPool()
}

func TestCreateNewCustomer(t *testing.T) {
	c := &Customer{
		Id:    1,
		Email: "alex@securemail.com",
	}
	c.CreateNewCustomer()
}
