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

type PluginsStatus struct {
	Endpoints ComponentWithStatus `json:"endpoints,omitempty"`

	Voicemails VoicemailsStatus `json:"voicemails,omitempty"`
}

// AssertPluginsStatusRequired checks if the required fields are not zero-ed
func AssertPluginsStatusRequired(obj PluginsStatus) error {
	if err := AssertComponentWithStatusRequired(obj.Endpoints); err != nil {
		return err
	}
	if err := AssertVoicemailsStatusRequired(obj.Voicemails); err != nil {
		return err
	}
	return nil
}

// AssertPluginsStatusConstraints checks if the values respects the defined constraints
func AssertPluginsStatusConstraints(obj PluginsStatus) error {
	return nil
}