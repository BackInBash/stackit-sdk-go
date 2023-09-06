/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

type CertCheckResponse struct {
	// REQUIRED
	CertChecks *[]CertCheckChildResponse `json:"certChecks"`
	// REQUIRED
	Message *string `json:"message"`
}