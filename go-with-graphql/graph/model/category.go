package model

type Author struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Email       string `json:"email"`
}

// `json:"description", omitempty` means that if the description is empty, it will not be included in the JSON response.
