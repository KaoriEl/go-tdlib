// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/kaoriEl/go-tdlib/tdlib"
)

// GetLogStream Returns information about currently used log stream for internal logging of TDLib. Can be called synchronously
func (client *Client) GetLogStream() (tdlib.LogStream, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getLogStream",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	switch tdlib.LogStreamEnum(result.Data["@type"].(string)) {

	case tdlib.LogStreamDefaultType:
		var logStream tdlib.LogStreamDefault
		err = json.Unmarshal(result.Raw, &logStream)
		return &logStream, err

	case tdlib.LogStreamFileType:
		var logStream tdlib.LogStreamFile
		err = json.Unmarshal(result.Raw, &logStream)
		return &logStream, err

	case tdlib.LogStreamEmptyType:
		var logStream tdlib.LogStreamEmpty
		err = json.Unmarshal(result.Raw, &logStream)
		return &logStream, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
