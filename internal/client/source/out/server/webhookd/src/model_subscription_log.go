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

import (
	"time"
)

type SubscriptionLog struct {

	// The current attempts
	Attempts int32 `json:"attempts,omitempty"`

	Detail HttpServiceLog `json:"detail,omitempty"`

	EndedAt time.Time `json:"ended_at,omitempty"`

	Event string `json:"event,omitempty"`

	// Limit of number of attempts
	MaxAttempts int32 `json:"max_attempts,omitempty"`

	StartedAt time.Time `json:"started_at,omitempty"`

	Status string `json:"status,omitempty"`

	Uuid string `json:"uuid,omitempty"`
}

// AssertSubscriptionLogRequired checks if the required fields are not zero-ed
func AssertSubscriptionLogRequired(obj SubscriptionLog) error {
	if err := AssertHttpServiceLogRequired(obj.Detail); err != nil {
		return err
	}
	return nil
}

// AssertSubscriptionLogConstraints checks if the values respects the defined constraints
func AssertSubscriptionLogConstraints(obj SubscriptionLog) error {
	return nil
}