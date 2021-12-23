// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/kaoriEl/go-tdlib/tdlib"
)

// GetChatFilter Returns information about a chat filter by its identifier
// @param chatFilterID Chat filter identifier
func (client *Client) GetChatFilter(chatFilterID int32) (*tdlib.ChatFilter, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":          "getChatFilter",
		"chat_filter_id": chatFilterID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chatFilterDummy tdlib.ChatFilter
	err = json.Unmarshal(result.Raw, &chatFilterDummy)
	return &chatFilterDummy, err

}
