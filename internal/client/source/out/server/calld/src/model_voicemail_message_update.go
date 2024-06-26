/*
 * accent-calld
 *
 * Control your calls from a REST API
 *
 * API version: 1.0.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package calldserver

type VoicemailMessageUpdate struct {

	// The folder's ID
	FolderId int32 `json:"folder_id"`
}

// AssertVoicemailMessageUpdateRequired checks if the required fields are not zero-ed
func AssertVoicemailMessageUpdateRequired(obj VoicemailMessageUpdate) error {
	elements := map[string]interface{}{
		"folder_id": obj.FolderId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertVoicemailMessageUpdateConstraints checks if the values respects the defined constraints
func AssertVoicemailMessageUpdateConstraints(obj VoicemailMessageUpdate) error {
	return nil
}
