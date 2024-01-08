/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

type CredentialsRotationState struct {
	LastCompletionTime *string `json:"lastCompletionTime,omitempty"`
	LastInitiationTime *string `json:"lastInitiationTime,omitempty"`
	// Phase of the credentials rotation. \"Never\" indicates that no credentials rotation has been performed using the new credentials rotation endpoints yet. Using the deprecated \"/v1/projects/{project_id}/clusters/{cluster_name}/rotate-credentials\" endpoint will not update this status field.
	Phase *string `json:"phase,omitempty"`
}