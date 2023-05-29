// AUTOGENERATED - DO NOT EDIT

package tdlib

import (
	"encoding/json"
)

// VoiceChat Describes a voice chat
type VoiceChat struct {
	tdCommon
	GroupCallID          int32         `json:"group_call_id"`          // Group call identifier of an active voice chat; 0 if none. Full informationa about the voice chat can be received through the method getGroupCall
	HasParticipants      bool          `json:"has_participants"`       // True, if the voice chat has participants
	DefaultParticipantID MessageSender `json:"default_participant_id"` // Default group call participant identifier to join the voice chat; may be null
}

// MessageType return the string telegram-type of VoiceChat
func (voiceChat *VoiceChat) MessageType() string {
	return "voiceChat"
}

// NewVoiceChat creates a new VoiceChat
//
// @param groupCallID Group call identifier of an active voice chat; 0 if none. Full informationa about the voice chat can be received through the method getGroupCall
// @param hasParticipants True, if the voice chat has participants
// @param defaultParticipantID Default group call participant identifier to join the voice chat; may be null
func NewVoiceChat(groupCallID int32, hasParticipants bool, defaultParticipantID MessageSender) *VoiceChat {
	voiceChatTemp := VoiceChat{
		tdCommon:             tdCommon{Type: "voiceChat"},
		GroupCallID:          groupCallID,
		HasParticipants:      hasParticipants,
		DefaultParticipantID: defaultParticipantID,
	}

	return &voiceChatTemp
}

// UnmarshalJSON unmarshal to json
func (voiceChat *VoiceChat) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		GroupCallID     int32 `json:"group_call_id"`    // Group call identifier of an active voice chat; 0 if none. Full informationa about the voice chat can be received through the method getGroupCall
		HasParticipants bool  `json:"has_participants"` // True, if the voice chat has participants

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	voiceChat.tdCommon = tempObj.tdCommon
	voiceChat.GroupCallID = tempObj.GroupCallID
	voiceChat.HasParticipants = tempObj.HasParticipants

	fieldDefaultParticipantID, _ := unmarshalMessageSender(objMap["default_participant_id"])
	voiceChat.DefaultParticipantID = fieldDefaultParticipantID

	return nil
}