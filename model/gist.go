package model

// SaveGistRequest represents the request body for saving a Gist.
type SaveGistRequest struct {
	Filename    string `json:"filename" validate:"required"`
	Content     string `json:"content" validate:"required"`
	Description string `json:"description"`
	Public      bool   `json:"public"` // Default to false if not provided
}

// SaveGistResponse represents the response body after saving a Gist.
type SaveGistResponse struct {
	GistURL string `json:"gistUrl"`
	GistID  string `json:"gistId"`
}