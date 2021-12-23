// AUTOGENERATED - DO NOT EDIT

package tdlib

// ChatEventLogFilters Represents a set of filters used to obtain a chat event log
type ChatEventLogFilters struct {
	tdCommon
	MessageEdits       bool `json:"message_edits"`       // True, if message edits should be returned
	MessageDeletions   bool `json:"message_deletions"`   // True, if message deletions should be returned
	MessagePins        bool `json:"message_pins"`        // True, if pin/unpin events should be returned
	MemberJoins        bool `json:"member_joins"`        // True, if members joining events should be returned
	MemberLeaves       bool `json:"member_leaves"`       // True, if members leaving events should be returned
	MemberInvites      bool `json:"member_invites"`      // True, if invited member events should be returned
	MemberPromotions   bool `json:"member_promotions"`   // True, if member promotion/demotion events should be returned
	MemberRestrictions bool `json:"member_restrictions"` // True, if member restricted/unrestricted/banned/unbanned events should be returned
	InfoChanges        bool `json:"info_changes"`        // True, if changes in chat information should be returned
	SettingChanges     bool `json:"setting_changes"`     // True, if changes in chat settings should be returned
	InviteLinkChanges  bool `json:"invite_link_changes"` // True, if changes to invite links should be returned
	VoiceChatChanges   bool `json:"voice_chat_changes"`  // True, if voice chat actions should be returned
}

// MessageType return the string telegram-type of ChatEventLogFilters
func (chatEventLogFilters *ChatEventLogFilters) MessageType() string {
	return "chatEventLogFilters"
}

// NewChatEventLogFilters creates a new ChatEventLogFilters
//
// @param messageEdits True, if message edits should be returned
// @param messageDeletions True, if message deletions should be returned
// @param messagePins True, if pin/unpin events should be returned
// @param memberJoins True, if members joining events should be returned
// @param memberLeaves True, if members leaving events should be returned
// @param memberInvites True, if invited member events should be returned
// @param memberPromotions True, if member promotion/demotion events should be returned
// @param memberRestrictions True, if member restricted/unrestricted/banned/unbanned events should be returned
// @param infoChanges True, if changes in chat information should be returned
// @param settingChanges True, if changes in chat settings should be returned
// @param inviteLinkChanges True, if changes to invite links should be returned
// @param voiceChatChanges True, if voice chat actions should be returned
func NewChatEventLogFilters(messageEdits bool, messageDeletions bool, messagePins bool, memberJoins bool, memberLeaves bool, memberInvites bool, memberPromotions bool, memberRestrictions bool, infoChanges bool, settingChanges bool, inviteLinkChanges bool, voiceChatChanges bool) *ChatEventLogFilters {
	chatEventLogFiltersTemp := ChatEventLogFilters{
		tdCommon:           tdCommon{Type: "chatEventLogFilters"},
		MessageEdits:       messageEdits,
		MessageDeletions:   messageDeletions,
		MessagePins:        messagePins,
		MemberJoins:        memberJoins,
		MemberLeaves:       memberLeaves,
		MemberInvites:      memberInvites,
		MemberPromotions:   memberPromotions,
		MemberRestrictions: memberRestrictions,
		InfoChanges:        infoChanges,
		SettingChanges:     settingChanges,
		InviteLinkChanges:  inviteLinkChanges,
		VoiceChatChanges:   voiceChatChanges,
	}

	return &chatEventLogFiltersTemp
}
