/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

type CreateScrapeConfigPayloadMetricsRelabelConfigsInner struct {
	// Action to perform based on regex matching. `Additional Validators:` * if action is replace, targetLabel needs to be in body
	Action *string `json:"action,omitempty"`
	// Modulus to take of the hash of the source label values.
	Modulus *float32 `json:"modulus,omitempty"`
	// Regular expression against which the extracted value is matched.
	Regex *string `json:"regex,omitempty"`
	// Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available.
	Replacement *string `json:"replacement,omitempty"`
	// Separator placed between concatenated source label values.
	Separator *string `json:"separator,omitempty"`
	// The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.
	SourceLabels *[]string `json:"sourceLabels,omitempty"`
	// Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.
	TargetLabel *string `json:"targetLabel,omitempty"`
}
