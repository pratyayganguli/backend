package sys

import "github.com/gin-gonic/gin"

type CustomerEnquiry struct {
	BusinessEmail string  `json:"email"`
	PhoneNumber   string  `json:"phoneNumber"`
	Query         *string `json:"query"`
}
type User struct {
	Id           uint64        `json:"id" gorm:"primaryKey"`
	Email        string        `json:"email"`
	Active       bool          `json:"active"`
	UserMetadata *UserMetadata `json:"userMetadata"`
	Plan         *Plan         `json:"plan"`
}
type UserMetadata struct {
	Name           string  `json:"name"`
	PhoneNumber    string  `json:"phoneNumber"`
	ProfilePicture string  `json:"profilePicture"`
	Organization   string  `json:"organization"`
	Bio            *string `json:"bio"`
}
type Plan struct {
	Type        string  `json:"type"`
	Price       float64 `json:"price"`
	MaxProjects uint8   `json:"maxProjects"`
}

func SaveCustomerEnquiryDetails(ctx *gin.Context) {
	// controller function for saving the Customer Enquiry Details...
	ce := &CustomerEnquiry{}
	if e := ctx.ShouldBindJSON(ce); e != nil {
		// return bad request
	}
	// some sort of sanitation and then save it
	if e := ce.Create(); e != nil {
		// return internal error
	}
	// compose a free user and have the soft insertion.
	u := composeTrialUser(ce.BusinessEmail, ce.PhoneNumber)
	if e := u.Create(); e != nil {
		// internal error handler
	}
}
func composeTrialUser(email, phoneNumber string) *User {
	return &User{
		Email: email,
		// after verification the user active will be set to true
		Active: false,
		Plan: &Plan{
			Type:        "TRIAL",
			Price:       0.00,
			MaxProjects: 2,
		},
		UserMetadata: &UserMetadata{
			PhoneNumber: phoneNumber,
		},
	}
}

func (u *User) Create() error {
	// gorm business logic over here
	return nil
}
func (ce *CustomerEnquiry) Create() error {
	// gorm business logic over here
	return nil
}

type UserOnboarder interface {
	Onboard() error
}
type TrialUserOnboarder struct {
	TrialUser User
}

// implement this
// todo: write the bare minimum business logic...
func (tuo *TrialUserOnboarder) Onboard() {
}

type PaidUserOnboarder struct {
	PaidUser User
}
