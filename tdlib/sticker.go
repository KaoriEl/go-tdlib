// AUTOGENERATED - DO NOT EDIT

package tdlib

// Sticker Describes a sticker
type Sticker struct {
	tdCommon
	SetID        JSONInt64          `json:"set_id"`        // The identifier of the sticker set to which the sticker belongs; 0 if none
	Width        int32              `json:"width"`         // Sticker width; as defined by the sender
	Height       int32              `json:"height"`        // Sticker height; as defined by the sender
	Emoji        string             `json:"emoji"`         // Emoji corresponding to the sticker
	IsAnimated   bool               `json:"is_animated"`   // True, if the sticker is an animated sticker in TGS format
	IsMask       bool               `json:"is_mask"`       // True, if the sticker is a mask
	MaskPosition *MaskPosition      `json:"mask_position"` // Position where the mask should be placed; may be null
	Outline      []ClosedVectorPath `json:"outline"`       // Sticker's outline represented as a list of closed vector paths; may be empty. The coordinate system origin is in the upper-left corner
	Thumbnail    *Thumbnail         `json:"thumbnail"`     // Sticker thumbnail in WEBP or JPEG format; may be null
	Sticker      *File              `json:"sticker"`       // File containing the sticker
}

// MessageType return the string telegram-type of Sticker
func (sticker *Sticker) MessageType() string {
	return "sticker"
}

// NewSticker creates a new Sticker
//
// @param setID The identifier of the sticker set to which the sticker belongs; 0 if none
// @param width Sticker width; as defined by the sender
// @param height Sticker height; as defined by the sender
// @param emoji Emoji corresponding to the sticker
// @param isAnimated True, if the sticker is an animated sticker in TGS format
// @param isMask True, if the sticker is a mask
// @param maskPosition Position where the mask should be placed; may be null
// @param outline Sticker's outline represented as a list of closed vector paths; may be empty. The coordinate system origin is in the upper-left corner
// @param thumbnail Sticker thumbnail in WEBP or JPEG format; may be null
// @param sticker File containing the sticker
func NewSticker(setID JSONInt64, width int32, height int32, emoji string, isAnimated bool, isMask bool, maskPosition *MaskPosition, outline []ClosedVectorPath, thumbnail *Thumbnail, sticker *File) *Sticker {
	stickerTemp := Sticker{
		tdCommon:     tdCommon{Type: "sticker"},
		SetID:        setID,
		Width:        width,
		Height:       height,
		Emoji:        emoji,
		IsAnimated:   isAnimated,
		IsMask:       isMask,
		MaskPosition: maskPosition,
		Outline:      outline,
		Thumbnail:    thumbnail,
		Sticker:      sticker,
	}

	return &stickerTemp
}
