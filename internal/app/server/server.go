package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/t67y110v/parser-service/docs"
	"github.com/t67y110v/parser-service/internal/app/config"
	"github.com/t67y110v/parser-service/internal/app/handlers"
	"github.com/t67y110v/parser-service/internal/app/logging"

	"github.com/gofiber/swagger" // swagger handler
)

type server struct {
	router   *fiber.App
	logger   logging.Logger
	config   *config.Config
	handlers *handlers.Handlers
}

func newServer(config *config.Config, log logging.Logger) *server {
	s := &server{
		router:   fiber.New(fiber.Config{ServerHeader: "software engineering course api", AppName: "Api v1.0.1"}),
		logger:   log,
		config:   config,
		handlers: handlers.NewHandlers(log),
	}
	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func (s *server) configureRouter() {
	s.router.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))
	//api := s.router.Group("/api")
	//api.Use(logger.New())
	// localhost:4000/user/register

	s.router.Get("/swagger/*", swagger.HandlerDefault)
	s.router.Get("/parse/:querry", s.handlers.Parse())
	s.router.Post("/parse/all", s.handlers.ParserAll())
	///////// USER GROUP ///////////////
	////////////////////////////////////

}
