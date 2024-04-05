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

type ApplicationSnoopPut struct {
	WhisperMode string `json:"whisper_mode,omitempty"`
}

// AssertApplicationSnoopPutRequired checks if the required fields are not zero-ed
func AssertApplicationSnoopPutRequired(obj ApplicationSnoopPut) error {
	return nil
}

// AssertApplicationSnoopPutConstraints checks if the values respects the defined constraints
func AssertApplicationSnoopPutConstraints(obj ApplicationSnoopPut) error {
	return nil
}
