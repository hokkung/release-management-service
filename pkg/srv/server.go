package srv

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/hokkung/release-management-service/pkg/srv/config"
)

// Server represents the HTTP server
type Server struct {
	App    *fiber.App
	config *config.Config
}

// New creates a new server instance
func New(customizer ServerCustomizer) *Server {
	fmt.Println(os.Getenv("SRV_APP_NAME"))
	config := config.NewConfig()
	srv := &Server{
		App: fiber.New(fiber.Config{
			AppName: config.AppName,
		}),
		config: config,
	}

	customizer.Register(srv)

	return srv
}

// Start starts the server
func (s *Server) Start() error {
	return s.App.Listen(fmt.Sprintf(":%d", s.config.Port))
}

// Stop stops the server
func (s *Server) Stop() error {
	return s.App.Shutdown()
}
