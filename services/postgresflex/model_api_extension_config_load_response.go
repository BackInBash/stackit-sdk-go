/*
STACKIT PostgreSQL Flex API

This is the documentation for the STACKIT postgres service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postgresflex

type ApiExtensionConfigLoadResponse struct {
	// Returns marshalled JSON of the new configuration of whatever extension is called
	Configuration *[]ApiConfiguration `json:"configuration,omitempty"`
}