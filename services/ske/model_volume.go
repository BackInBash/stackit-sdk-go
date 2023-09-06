/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

type Volume struct {
	// REQUIRED
	Size *int32 `json:"size"`
	// For valid values please take a look at /provider-options volumeTypes
	Type *string `json:"type,omitempty"`
}