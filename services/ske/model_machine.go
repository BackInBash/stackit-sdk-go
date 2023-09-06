/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

type Machine struct {
	// REQUIRED
	Image *Image `json:"image"`
	// For valid types please take a look at /provider-options machineTypes
	// REQUIRED
	Type *string `json:"type"`
}