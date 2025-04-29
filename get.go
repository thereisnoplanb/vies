package vies

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/thereisnoplanb/vies/enums/CountryCode"
)

type vies struct {
	IsValid           bool            `json:"isValid"`
	RequestDate       time.Time       `json:"requestDate"`
	UserError         string          `json:"userError"`
	Name              string          `json:"name"`
	Address           string          `json:"address"`
	RequestIdentifier string          `json:"requestIdentifier"`
	OriginalVatNumber string          `json:"originalVatNumber"`
	VatNumber         string          `json:"vatNumber"`
	ViesApproximate   viesApproximate `json:"viesApproximate"`
}

type viesApproximate struct {
	Name             string `json:"name"`
	Street           string `json:"street"`
	PostalCode       string `json:"postalCode"`
	City             string `json:"city"`
	CompanyType      string `json:"companyType"`
	MatchName        int    `json:"matchName"`
	MatchStreet      int    `json:"matchStreet"`
	MatchPostalCode  int    `json:"matchPostalCode"`
	MatchCity        int    `json:"matchCity"`
	MatchCompanyType int    `json:"matchCompanyType"`
}

func (client *Client) get(euVat string) (vies vies, statusCode int, err error) {

	countryCode := CountryCode.Enum("")
	if len(euVat) > 1 {
		countryCode = CountryCode.Enum(euVat[:2])
	}
	vatNumber := ""
	if len(euVat) > 2 {
		vatNumber = euVat[2:]
	}

	if !countryCode.IsDefined() {
		return vies, statusCode, errors.New("country code is not defined")
	}

	response, err := client.httpClient.Get(fmt.Sprintf("%s/%s/vat/%s", address, countryCode, vatNumber))
	if err != nil {
		return vies, statusCode, err
	}

	statusCode = response.StatusCode
	if statusCode != http.StatusOK {
		return vies, statusCode, nil
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return vies, statusCode, err
	}

	err = json.Unmarshal(body, &vies)
	return vies, statusCode, err
}
