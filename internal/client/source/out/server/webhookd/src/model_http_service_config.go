/*
 * accent-webhookd
 *
 * Control your webhooks from a REST API
 *
 * API version: 1.0.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package webhookdserver

type HttpServiceConfig struct {

	// Jinja2 template, where variables come from the event triggering the webhook. For more details, see https://accentvoice.io/uc-doc/api_sdk/rest_api/webhookd/templates. **Default:** the complete event data, JSON-encoded.
	Body string `json:"body,omitempty"`

	// Content-Type of the body
	ContentType string `json:"content_type,omitempty"`

	Method string `json:"method"`

	// Jinja2 template, where variables come from the event triggering the webhook. For more details, see https://accentvoice.io/uc-doc/api_sdk/rest_api/webhookd/templates
	Url string `json:"url"`

	// May be `true`, `false` or a path to the certificate bundle
	VerifyCertificate string `json:"verify_certificate,omitempty"`
}

// AssertHttpServiceConfigRequired checks if the required fields are not zero-ed
func AssertHttpServiceConfigRequired(obj HttpServiceConfig) error {
	elements := map[string]interface{}{
		"method": obj.Method,
		"url":    obj.Url,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertHttpServiceConfigConstraints checks if the values respects the defined constraints
func AssertHttpServiceConfigConstraints(obj HttpServiceConfig) error {
	return nil
}
