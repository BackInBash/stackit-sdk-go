/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

type BackupRetentionResponse struct {
	// REQUIRED
	AlertConfigBackupRetention *string `json:"alertConfigBackupRetention"`
	// REQUIRED
	AlertRulesBackupRetention *string `json:"alertRulesBackupRetention"`
	// REQUIRED
	GrafanaBackupRetention *string `json:"grafanaBackupRetention"`
	// REQUIRED
	Message *string `json:"message"`
	// REQUIRED
	ScrapeConfigBackupRetention *string `json:"scrapeConfigBackupRetention"`
}
