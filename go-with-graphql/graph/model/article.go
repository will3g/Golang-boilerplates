package model

type Article struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Body     string `json:"body"`
	AuthorID string `json:"authorId"`
}

// `json:"description", omitempty` means that if the description is empty, it will not be included in the JSON response.
