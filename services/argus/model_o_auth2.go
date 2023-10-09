/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

type OAuth2 struct {
	// REQUIRED
	ClientId *string `json:"clientId"`
	// REQUIRED
	ClientSecret *string    `json:"clientSecret"`
	Scopes       *[]string  `json:"scopes,omitempty"`
	TlsConfig    *TLSConfig `json:"tlsConfig,omitempty"`
	// REQUIRED
	TokenUrl *string `json:"tokenUrl"`
}
