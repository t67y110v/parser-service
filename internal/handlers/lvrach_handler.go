package handlers

import (
	"net/http"

	"github.com/t67y110v/parser-service/internal/handlers/requests"
	parser "github.com/t67y110v/parser-service/internal/parser"

	"github.com/gofiber/fiber/v2"
)

// @Summary Parse page by current query
// @Description pars site to get informaion about nr
// @Tags         Parser
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.Body true "create new user"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /parser/lvrach [post]
func (h *Handlers) HandleLvrachParsePage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := &requests.Body{}
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		p := parser.NewLvrach(h.logger)

		articles, err := p.Parse.Page(req.Query, 1)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(articles)
	}
}

// @Summary Parse all pages by current query
// @Description pars site to get informaion about nr
// @Tags         Parser
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.Body  true "create new user"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /parser/lvrach/all [post]
func (h *Handlers) HandleLvrachParseAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := requests.Body{}
		if err := c.BodyParser(&req); err != nil {
			h.logger.Error(err)
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})

		}
		p := parser.NewLvrach(h.logger)
		articles, err := p.Parse.All(req.Query)
		if err != nil {
			h.logger.Error(err)
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(articles)
	}
}
