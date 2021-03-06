// AUTOGENERATED - DO NOT EDIT

package tdlib

// LocalizationTargetInfo Contains information about the current localization target
type LocalizationTargetInfo struct {
	tdCommon
	LanguagePacks []LanguagePackInfo `json:"language_packs"` // List of available language packs for this application
}

// MessageType return the string telegram-type of LocalizationTargetInfo
func (localizationTargetInfo *LocalizationTargetInfo) MessageType() string {
	return "localizationTargetInfo"
}

// NewLocalizationTargetInfo creates a new LocalizationTargetInfo
//
// @param languagePacks List of available language packs for this application
func NewLocalizationTargetInfo(languagePacks []LanguagePackInfo) *LocalizationTargetInfo {
	localizationTargetInfoTemp := LocalizationTargetInfo{
		tdCommon:      tdCommon{Type: "localizationTargetInfo"},
		LanguagePacks: languagePacks,
	}

	return &localizationTargetInfoTemp
}
