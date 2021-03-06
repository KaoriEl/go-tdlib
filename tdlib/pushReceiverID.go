// AUTOGENERATED - DO NOT EDIT

package tdlib

// PushReceiverID Contains a globally unique push receiver identifier, which can be used to identify which account has received a push notification
type PushReceiverID struct {
	tdCommon
	ID JSONInt64 `json:"id"` // The globally unique identifier of push notification subscription
}

// MessageType return the string telegram-type of PushReceiverID
func (pushReceiverID *PushReceiverID) MessageType() string {
	return "pushReceiverId"
}

// NewPushReceiverID creates a new PushReceiverID
//
// @param iD The globally unique identifier of push notification subscription
func NewPushReceiverID(iD JSONInt64) *PushReceiverID {
	pushReceiverIDTemp := PushReceiverID{
		tdCommon: tdCommon{Type: "pushReceiverId"},
		ID:       iD,
	}

	return &pushReceiverIDTemp
}
