/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

type MachineImageVersion struct {
	Cri            *[]CRI  `json:"cri,omitempty"`
	ExpirationDate *string `json:"expirationDate,omitempty"`
	State          *string `json:"state,omitempty"`
	Version        *string `json:"version,omitempty"`
}