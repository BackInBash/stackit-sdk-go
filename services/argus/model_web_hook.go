/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

type WebHook struct {
	MsTeams      *bool `json:"msTeams,omitempty"`
	SendResolved *bool `json:"sendResolved,omitempty"`
	// REQUIRED
	Url *string `json:"url"`
}