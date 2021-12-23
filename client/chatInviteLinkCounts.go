// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/kaoriEl/go-tdlib/tdlib"
)

// GetChatInviteLinkCounts Returns list of chat administrators with number of their invite links. Requires owner privileges in the chat
// @param chatID Chat identifier
func (client *Client) GetChatInviteLinkCounts(chatID int64) (*tdlib.ChatInviteLinkCounts, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "getChatInviteLinkCounts",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chatInviteLinkCounts tdlib.ChatInviteLinkCounts
	err = json.Unmarshal(result.Raw, &chatInviteLinkCounts)
	return &chatInviteLinkCounts, err

}
