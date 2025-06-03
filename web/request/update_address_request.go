package request

type UpdateAddressRequest struct {
	ContactID  string `json:"_;validate:required"`
	AddressID  string `json:"_;validate:required"`
	Street     string `json:"street;validate:max=200"`
	City       string `json:"city;validate:max=100"`
	Province   string `json:"province;validate:max=100"`
	Country    string `json:"country;validate:required;max=100"`
	PostalCode string `json:"postal_code;validate:max=100"`
}
