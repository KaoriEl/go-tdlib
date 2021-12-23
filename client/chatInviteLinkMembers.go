// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/kaoriEl/go-tdlib/tdlib"
)

// GetChatInviteLinkMembers Returns chat members joined a chat by an invite link. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links
// @param chatID Chat identifier
// @param inviteLink Invite link for which to return chat members
// @param offsetMember A chat member from which to return next chat members; use null to get results from the beginning
// @param limit The maximum number of chat members to return
func (client *Client) GetChatInviteLinkMembers(chatID int64, inviteLink string, offsetMember *tdlib.ChatInviteLinkMember, limit int32) (*tdlib.ChatInviteLinkMembers, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":         "getChatInviteLinkMembers",
		"chat_id":       chatID,
		"invite_link":   inviteLink,
		"offset_member": offsetMember,
		"limit":         limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chatInviteLinkMembers tdlib.ChatInviteLinkMembers
	err = json.Unmarshal(result.Raw, &chatInviteLinkMembers)
	return &chatInviteLinkMembers, err

}
