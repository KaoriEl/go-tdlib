// AUTOGENERATED - DO NOT EDIT

package tdlib

// TMeURLs Contains a list of t.me URLs
type TMeURLs struct {
	tdCommon
	URLs []TMeURL `json:"urls"` // List of URLs
}

// MessageType return the string telegram-type of TMeURLs
func (tMeURLs *TMeURLs) MessageType() string {
	return "tMeUrls"
}

// NewTMeURLs creates a new TMeURLs
//
// @param uRLs List of URLs
func NewTMeURLs(uRLs []TMeURL) *TMeURLs {
	tMeURLsTemp := TMeURLs{
		tdCommon: tdCommon{Type: "tMeUrls"},
		URLs:     uRLs,
	}

	return &tMeURLsTemp
}
