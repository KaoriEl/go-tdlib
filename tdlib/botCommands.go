// AUTOGENERATED - DO NOT EDIT

package tdlib

// BotCommands Contains a list of bot commands
type BotCommands struct {
	tdCommon
	BotUserID int64        `json:"bot_user_id"` // Bot's user identifier
	Commands  []BotCommand `json:"commands"`    // List of bot commands
}

// MessageType return the string telegram-type of BotCommands
func (botCommands *BotCommands) MessageType() string {
	return "botCommands"
}

// NewBotCommands creates a new BotCommands
//
// @param botUserID Bot's user identifier
// @param commands List of bot commands
func NewBotCommands(botUserID int64, commands []BotCommand) *BotCommands {
	botCommandsTemp := BotCommands{
		tdCommon:  tdCommon{Type: "botCommands"},
		BotUserID: botUserID,
		Commands:  commands,
	}

	return &botCommandsTemp
}
