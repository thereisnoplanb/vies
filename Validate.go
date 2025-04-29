package vies

import (
	"errors"
	"net/http"
)

// Validates EU VAT number.
//
// # Parameters:
//
//	euVat string
//
// EU VAT, for example: PL7251868136.
//
// # Returns:
//
//	isValid bool
//
// True if EU VAT is valid; false, otherwise.
//
// # Error:
//
//	err error
//
// Returns error when country code for EU VAT is not correct, connection with API is broken or API response status is not correct.
func (client *Client) Validate(euVat string) (isValid bool, err error) {

	vies, statusCode, err := client.get(euVat)
	if err != nil {
		return false, err
	}
	if statusCode != http.StatusOK {
		return false, errors.New("expected status code: 200, OK")
	}

	return vies.IsValid, nil
}
