/*
 * accent-plugind
 *
 * Accent's plugin management service
 *
 * API version: 0.2
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package plugindserver

type PluginInstallParameters struct {

	// The method used to fetch this plugin
	Method string `json:"method"`

	// Method dependant installation options
	Options map[string]interface{} `json:"options,omitempty"`
}

// AssertPluginInstallParametersRequired checks if the required fields are not zero-ed
func AssertPluginInstallParametersRequired(obj PluginInstallParameters) error {
	elements := map[string]interface{}{
		"method": obj.Method,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertPluginInstallParametersConstraints checks if the values respects the defined constraints
func AssertPluginInstallParametersConstraints(obj PluginInstallParameters) error {
	return nil
}
