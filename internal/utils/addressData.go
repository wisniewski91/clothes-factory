package utils

import "strings"

type AddressData struct {
	Country     string
	City        string
	District    string
	Street      string
	PostalCode  string
	HouseNumber string
	FlatNumber  string
}

func (a AddressData) Validate() error {
	var e []string
	if a.Country == "" {
		e = append(e, "The country field cannot be empty.")
	}
	if a.City == "" {
		e = append(e, "The city field cannot be empty.")
	}
	if a.District == "" {
		e = append(e, "District field cannot be empty.")
	}
	if a.Street == "" {
		e = append(e, "Street field cannot be empty.")
	}
	if a.PostalCode == "" {
		e = append(e, "Postal code field cannot be empty.")
	}
	if a.HouseNumber == "" {
		e = append(e, "House number field cannot be empty.")
	}

	if len(e) == 0 {
		return nil
	}

	return ValidateError{
		Code:    -10,
		Object:  "AddressData",
		Message: strings.Join(e, "\n"),
	}
}
