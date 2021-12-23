// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/kaoriEl/go-tdlib/tdlib"
)

// GetScopeNotificationSettings Returns the notification settings for chats of a given type
// @param scope Types of chats for which to return the notification settings information
func (client *Client) GetScopeNotificationSettings(scope tdlib.NotificationSettingsScope) (*tdlib.ScopeNotificationSettings, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getScopeNotificationSettings",
		"scope": scope,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var scopeNotificationSettings tdlib.ScopeNotificationSettings
	err = json.Unmarshal(result.Raw, &scopeNotificationSettings)
	return &scopeNotificationSettings, err

}
