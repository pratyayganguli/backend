package features

// we will be using the same struct for most of the examples in feature comparison examples

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
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Age         uint8       `json:"age"`
	Gender      Gender      `json:"gender"`
	BloodGroup  BloodGroup  `json:"bloodGroup"`
	Phone       string      `json:"phone"`
	UserProfile UserProfile `json:"userProfile"`
}

// Metadata related to user
type UserProfile struct {
	Id             string `json:"id"`
	Address        string `json:"address"`
	ProfilePicture string `json:"profilePicture"`
	EmailAddress   string `json:"emailAddress"`
}

// Validate all the data related to patient at the time of onboarding.
func (p *User) Validate() {
}

func (u *User) Create() {
}

// Create users in batches, sort of a bulk creation function...
func CreateBatchUser([]User) {
}
