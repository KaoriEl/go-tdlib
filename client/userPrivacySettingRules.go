// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/kaoriEl/go-tdlib/tdlib"
)

// GetUserPrivacySettingRules Returns the current privacy settings
// @param setting The privacy setting
func (client *Client) GetUserPrivacySettingRules(setting tdlib.UserPrivacySetting) (*tdlib.UserPrivacySettingRules, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "getUserPrivacySettingRules",
		"setting": setting,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var userPrivacySettingRules tdlib.UserPrivacySettingRules
	err = json.Unmarshal(result.Raw, &userPrivacySettingRules)
	return &userPrivacySettingRules, err

}
