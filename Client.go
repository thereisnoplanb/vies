package vies

import "net/http"

const address string = "https://ec.europa.eu/taxation_customs/vies/rest-api/ms"

// VIES API client.
type Client struct {
	httpClient http.Client
}
