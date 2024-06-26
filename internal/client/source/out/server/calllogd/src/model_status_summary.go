/*
 * accent-call-logd
 *
 * Consult call logs from a REST API
 *
 * API version: 1.0.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package calllogdserver

type StatusSummary struct {
	BusConsumer ComponentWithStatus `json:"bus_consumer,omitempty"`

	ServiceToken ComponentWithStatus `json:"service_token,omitempty"`

	TaskQueue ComponentWithStatus `json:"task_queue,omitempty"`
}

// AssertStatusSummaryRequired checks if the required fields are not zero-ed
func AssertStatusSummaryRequired(obj StatusSummary) error {
	if err := AssertComponentWithStatusRequired(obj.BusConsumer); err != nil {
		return err
	}
	if err := AssertComponentWithStatusRequired(obj.ServiceToken); err != nil {
		return err
	}
	if err := AssertComponentWithStatusRequired(obj.TaskQueue); err != nil {
		return err
	}
	return nil
}

// AssertStatusSummaryConstraints checks if the values respects the defined constraints
func AssertStatusSummaryConstraints(obj StatusSummary) error {
	return nil
}
