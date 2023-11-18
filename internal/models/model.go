package models

type Article struct {
	ID         int    `json:"id,omitempty"`
	Name       string ` json:"name"`
	Annotation string `json:"annotation"`
	Link       string `json:"link"`
}

type Response struct {
	Found    int       `json:"found,omitempty"`
	Articles []Article `json:"articles"`
}
