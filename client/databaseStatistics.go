// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/kaoriEl/go-tdlib/tdlib"
)

// GetDatabaseStatistics Returns database statistics
func (client *Client) GetDatabaseStatistics() (*tdlib.DatabaseStatistics, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getDatabaseStatistics",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var databaseStatistics tdlib.DatabaseStatistics
	err = json.Unmarshal(result.Raw, &databaseStatistics)
	return &databaseStatistics, err

}
