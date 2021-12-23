// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/kaoriEl/go-tdlib/tdlib"
)

// GetChatInviteLinks Returns invite links for a chat created by specified administrator. Requires administrator privileges and can_invite_users right in the chat to get own links and owner privileges to get other links
// @param chatID Chat identifier
// @param creatorUserID User identifier of a chat administrator. Must be an identifier of the current user for non-owner
// @param isRevoked Pass true if revoked links needs to be returned instead of active or expired
// @param offsetDate Creation date of an invite link starting after which to return invite links; use 0 to get results from the beginning
// @param offsetInviteLink Invite link starting after which to return invite links; use empty string to get results from the beginning
// @param limit The maximum number of invite links to return
func (client *Client) GetChatInviteLinks(chatID int64, creatorUserID int64, isRevoked bool, offsetDate int32, offsetInviteLink string, limit int32) (*tdlib.ChatInviteLinks, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":              "getChatInviteLinks",
		"chat_id":            chatID,
		"creator_user_id":    creatorUserID,
		"is_revoked":         isRevoked,
		"offset_date":        offsetDate,
		"offset_invite_link": offsetInviteLink,
		"limit":              limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chatInviteLinks tdlib.ChatInviteLinks
	err = json.Unmarshal(result.Raw, &chatInviteLinks)
	return &chatInviteLinks, err

}

// RevokeChatInviteLink Revokes invite link for a chat. Available for basic groups, supergroups, and channels. Requires administrator privileges and can_invite_users right in the chat for own links and owner privileges for other links. If a primary link is revoked, then additionally to the revoked link returns new primary link
// @param chatID Chat identifier
// @param inviteLink Invite link to be revoked
func (client *Client) RevokeChatInviteLink(chatID int64, inviteLink string) (*tdlib.ChatInviteLinks, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":       "revokeChatInviteLink",
		"chat_id":     chatID,
		"invite_link": inviteLink,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chatInviteLinks tdlib.ChatInviteLinks
	err = json.Unmarshal(result.Raw, &chatInviteLinks)
	return &chatInviteLinks, err

}
