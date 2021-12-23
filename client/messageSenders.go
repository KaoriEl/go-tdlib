// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/kaoriEl/go-tdlib/tdlib"
)

// GetVoiceChatAvailableParticipants Returns list of participant identifiers, which can be used to join voice chats in a chat
// @param chatID Chat identifier
func (client *Client) GetVoiceChatAvailableParticipants(chatID int64) (*tdlib.MessageSenders, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "getVoiceChatAvailableParticipants",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var messageSenders tdlib.MessageSenders
	err = json.Unmarshal(result.Raw, &messageSenders)
	return &messageSenders, err

}

// GetBlockedMessageSenders Returns users and chats that were blocked by the current user
// @param offset Number of users and chats to skip in the result; must be non-negative
// @param limit The maximum number of users and chats to return; up to 100
func (client *Client) GetBlockedMessageSenders(offset int32, limit int32) (*tdlib.MessageSenders, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":  "getBlockedMessageSenders",
		"offset": offset,
		"limit":  limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var messageSenders tdlib.MessageSenders
	err = json.Unmarshal(result.Raw, &messageSenders)
	return &messageSenders, err

}
