package responses

type ParserResult struct {
	Result string `json:"result"`
}

type Error struct {
	Message string `json:"message"`
}

type Article struct {
	ID         int    `json:"id"`
	Name       string ` json:"name"`
	Annotation string `json:"annotation"`
	Link       string `json:"link"`
}

type Response struct {
	Found    int       `json:"found,omitempty"`
	Articles []Article `json:"articles"`
}
