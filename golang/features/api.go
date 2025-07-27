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

// These are the set of data which we will be collecting at the time of app registration. All fields are mandatory
type Patient struct {
	// auto generated and BSON id
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	Age        uint8      `json:"age"`
	Gender     Gender     `json:"gender"`
	BloodGroup BloodGroup `json:"bloodGroup"`
	Phone      string     `json:"phone"`
}

// Validate all the data related to patient at the time of onboarding.
func (p *Patient) validate() {
	// check for all the nil checks and handle the erors
}

type PatientProfile struct {
}
