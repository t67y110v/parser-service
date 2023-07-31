package responses

type Article struct {
	Name        string   ` json:"name"`
	Annotation  string   `json:"annotation"`
	Link        string   `json:"link"`
	Authors     []string `json:"authors,omitempty"`
	Year        int      `json:"year,omitempty"`
	Journal     string   `json:"journal,omitempty"`
	JournalLink string   `json:"journal_link,omitempty"`
	OCR         []string `json:"ocr,omitempty"`
}
