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

type QueueQoSStatistic struct {

	// Start of the statistic interval.
	From string `json:"from,omitempty"`

	QualityOfService []QueueQoSStatisticQualityOfServiceInner `json:"quality_of_service,omitempty"`

	// ID of the corresponding queue.
	QueueId int32 `json:"queue_id,omitempty"`

	// Name of the corresponding queue.
	QueueName string `json:"queue_name,omitempty"`

	// Tenant UUID of the corresponding queue.
	TenantUuid string `json:"tenant_uuid,omitempty"`

	// End of the statistic interval.
	Until string `json:"until,omitempty"`
}

// AssertQueueQoSStatisticRequired checks if the required fields are not zero-ed
func AssertQueueQoSStatisticRequired(obj QueueQoSStatistic) error {
	for _, el := range obj.QualityOfService {
		if err := AssertQueueQoSStatisticQualityOfServiceInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertQueueQoSStatisticConstraints checks if the values respects the defined constraints
func AssertQueueQoSStatisticConstraints(obj QueueQoSStatistic) error {
	return nil
}