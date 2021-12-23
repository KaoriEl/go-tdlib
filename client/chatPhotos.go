// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/kaoriEl/go-tdlib/tdlib"
)

// GetUserProfilePhotos Returns the profile photos of a user. The result of this query may be outdated: some photos might have been deleted already
// @param userID User identifier
// @param offset The number of photos to skip; must be non-negative
// @param limit The maximum number of photos to be returned; up to 100
func (client *Client) GetUserProfilePhotos(userID int64, offset int32, limit int32) (*tdlib.ChatPhotos, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "getUserProfilePhotos",
		"user_id": userID,
		"offset":  offset,
		"limit":   limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chatPhotos tdlib.ChatPhotos
	err = json.Unmarshal(result.Raw, &chatPhotos)
	return &chatPhotos, err

}
