package passgen

type SendOTPReq struct {
	PhoneNumber string `json:"phoneNumber"`
	CountryCode string `json:"countryCode"`
}

type VerifyOTPReq struct {
	PhoneNumber string `json:"phoneNumber"`
	OTP         string `json:"otp"`
}

func (sor *SendOTPReq) Validate() error {
	var e error
	if e = ValidCountryCode(sor.CountryCode); e == nil {
		if e = ValidPhoneNumber(sor.PhoneNumber); e == nil {
			return nil
		}
	}
	return e
}
