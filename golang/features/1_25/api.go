package features

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//we will be using the same struct for most of the examples in feature comparison examples

type Gender string

type BloodGroup string

var (
	Male        Gender = "MALE"
	Female      Gender = "FEMALE"
	Transgender Gender = "TRANSGENDER"
)

var (
	A_Positive BloodGroup = "A+"
	O_Positive BloodGroup = "O+"
)

// These are the set of data which we will be collecting at the time of registration.
type User struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `json:"name"`
	Age         uint8              `json:"age"`
	Gender      Gender             `json:"gender"`
	BloodGroup  BloodGroup         `json:"bloodGroup"`
	Phone       string             `json:"phone"`
	UserProfile *UserProfile       `json:"userProfile"`
}

// Metadata related to user
type UserProfile struct {
	Id             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Address        string             `json:"address"`
	ProfilePicture string             `json:"profilePicture"`
	EmailAddress   string             `json:"emailAddress"`
}

// Validate all the data related to patient at the time of onboarding.
func (u *User) Validate() {
}

func (u *User) CreateDummy() {
	u = &User{
		Id:   primitive.NewObjectID(),
		Name: gofakeit.Name(),
		// ommit check
		//Age:    gofakeit.Uint8(),
		Gender: Male,
		Phone:  gofakeit.Phone(),
		// ommit check for struct pointers
		// UserProfile: &UserProfile{
		// 	Id:             primitive.NewObjectID(),
		// 	Address:        gofakeit.Address().Address,
		// 	ProfilePicture: gofakeit.Person().Image,
		// 	EmailAddress:   gofakeit.Email(),
		// },
	}
	if byteArr, err := json.Marshal(u, json.OmitZeroStructFields(true), jsontext.WithIndent("\t")); err == nil {
		log.Printf("Encoded res : %s", (string(byteArr)))
	}
}

func (u *User) Create() {
	//business logic here
}

// Create users in batches, sort of a bulk creation function...
func CreateBatchUser([]User) {
	//business logic here
}
