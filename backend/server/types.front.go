// This contains required bodies for JSON

package deliverd

type B_Courier struct {
	NameFirst       string `json:"name_first" validate:"required"`
	NameMiddle      string `json:"name_middle" validate:"required"`
	NameLast        string `json:"name_last" validate:"required"`
	AddressFirst    string `json:"address_first" validate:"required"`
	AddressSecond   string `json:"address_second" validate:"required"`
	AddressPostcode string `json:"address_postcode" validate:"required"`
	AddressCounty   string `json:"address_county" validate:"required"`
	AddressCountry  string `json:"address_country" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	PhoneNumber     string `json:"phone_number" validate:"required"`
	DateBirth       int    `json:"date_birth" validate:"required"`
}

type B_Partner struct {
	CompanyName     string `json:"company_name" validate:"required"`
	CompanyID       string `json:"company_id" validate:"required"`
	AddressFirst    string `json:"address_first" validate:"required"`
	AddressSecond   string `json:"address_second" validate:"required"`
	AddressPostcode string `json:"address_postcode" validate:"required"`
	AddressCounty   string `json:"address_county" validate:"required"`
	AddressCountry  string `json:"address_country" validate:"required"`
	RepNameFirst    string `json:"rep_name_first" validate:"required"`
	RepNameLast     string `json:"rep_name_last" validate:"required"`
	RepEmail        string `json:"rep_email" validate:"required,email"`
	RepPhone        string `json:"rep_phone" validate:"required"`
}
