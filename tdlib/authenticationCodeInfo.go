// AUTOGENERATED - DO NOT EDIT

package tdlib

import (
	"encoding/json"
)

// AuthenticationCodeInfo Information about the authentication code that was sent
type AuthenticationCodeInfo struct {
	tdCommon
	PhoneNumber string                 `json:"phone_number"` // A phone number that is being authenticated
	Type        AuthenticationCodeType `json:"type"`         // Describes the way the code was sent to the user
	NextType    AuthenticationCodeType `json:"next_type"`    // Describes the way the next code will be sent to the user; may be null
	Timeout     int32                  `json:"timeout"`      // Timeout before the code can be re-sent, in seconds
}

// MessageType return the string telegram-type of AuthenticationCodeInfo
func (authenticationCodeInfo *AuthenticationCodeInfo) MessageType() string {
	return "authenticationCodeInfo"
}

// NewAuthenticationCodeInfo creates a new AuthenticationCodeInfo
//
// @param phoneNumber A phone number that is being authenticated
// @param typeParam Describes the way the code was sent to the user
// @param nextType Describes the way the next code will be sent to the user; may be null
// @param timeout Timeout before the code can be re-sent, in seconds
func NewAuthenticationCodeInfo(phoneNumber string, typeParam AuthenticationCodeType, nextType AuthenticationCodeType, timeout int32) *AuthenticationCodeInfo {
	authenticationCodeInfoTemp := AuthenticationCodeInfo{
		tdCommon:    tdCommon{Type: "authenticationCodeInfo"},
		PhoneNumber: phoneNumber,
		Type:        typeParam,
		NextType:    nextType,
		Timeout:     timeout,
	}

	return &authenticationCodeInfoTemp
}

// UnmarshalJSON unmarshal to json
func (authenticationCodeInfo *AuthenticationCodeInfo) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		PhoneNumber string `json:"phone_number"` // A phone number that is being authenticated
		Timeout     int32  `json:"timeout"`      // Timeout before the code can be re-sent, in seconds
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	authenticationCodeInfo.tdCommon = tempObj.tdCommon
	authenticationCodeInfo.PhoneNumber = tempObj.PhoneNumber
	authenticationCodeInfo.Timeout = tempObj.Timeout

	fieldType, _ := unmarshalAuthenticationCodeType(objMap["type"])
	authenticationCodeInfo.Type = fieldType

	fieldNextType, _ := unmarshalAuthenticationCodeType(objMap["next_type"])
	authenticationCodeInfo.NextType = fieldNextType

	return nil
}
