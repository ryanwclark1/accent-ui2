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

type HttpServiceLog struct {
	RequestBody string `json:"request_body,omitempty"`

	RequestHeaders map[string]string `json:"request_headers,omitempty"`

	RequestMethod string `json:"request_method,omitempty"`

	RequestUrl string `json:"request_url,omitempty"`

	ResponseBody string `json:"response_body,omitempty"`

	ResponseHeaders map[string]string `json:"response_headers,omitempty"`

	ResponseMethod string `json:"response_method,omitempty"`

	ResponseUrl string `json:"response_url,omitempty"`
}

// AssertHttpServiceLogRequired checks if the required fields are not zero-ed
func AssertHttpServiceLogRequired(obj HttpServiceLog) error {
	return nil
}

// AssertHttpServiceLogConstraints checks if the values respects the defined constraints
func AssertHttpServiceLogConstraints(obj HttpServiceLog) error {
	return nil
}
