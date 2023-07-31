package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/t67y110v/parser-service/internal/app/parser"
)

// @Summary Parse
// @Description Parser
// @Tags         Parser
//
//	@Accept       json
//
// @Param        category   path      string  true  "Category"
// @Produce json
// @Success 200 {object} responses.Article
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /parse/{querry} [get]
func (h *Handlers) Parse() fiber.Handler {
	return func(c *fiber.Ctx) error {

		querry := c.Params("querry")

		a := parser.Parse(querry)
		return c.JSON(a)
	}
}

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
