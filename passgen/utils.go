package passgen

import "fmt"

const IND = "+91"
const US = "+1"
const CHINA = "+86"
const RUSSIA = "+7"

func ValidCountryCode(countryCode string) error {
	if countryCode == IND || countryCode == US || countryCode == CHINA || countryCode == RUSSIA {
		return nil
	}
	return fmt.Errorf("Invalid country code: %s", countryCode)
}

func ValidPhoneNumber(phoneNumber string) error {
	// write the business logic if the phone number has a whatsapp account or not...
}
