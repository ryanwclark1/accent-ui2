/*
 * accent-chatd
 *
 * Control your message and presence from a REST API
 *
 * API version: 1.0.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package chatdserver

type Message struct {

	// Alias/nickname of the sender
	Alias string `json:"alias,omitempty"`

	// The content of the message
	Content string `json:"content,omitempty"`

	// The date of the message's creation
	CreatedAt string `json:"created_at,omitempty"`

	Room RoomRelationBase `json:"room,omitempty"`

	// tenant uuid of the sender
	TenantUuid string `json:"tenant_uuid,omitempty"`

	// user uuid of the sender
	UserUuid string `json:"user_uuid,omitempty"`

	// The UUID of the message
	Uuid string `json:"uuid,omitempty"`

	// accent uuid of the sender
	AccentUuid string `json:"accent_uuid,omitempty"`
}

// AssertMessageRequired checks if the required fields are not zero-ed
func AssertMessageRequired(obj Message) error {
	if err := AssertRoomRelationBaseRequired(obj.Room); err != nil {
		return err
	}
	return nil
}

// AssertMessageConstraints checks if the values respects the defined constraints
func AssertMessageConstraints(obj Message) error {
	return nil
}
