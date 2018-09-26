package model

// Dream type declaration
type Dream struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	ImageURI    *string `json:"image_uri"`
	Todo        []Todo  `json:"todo"`
}
