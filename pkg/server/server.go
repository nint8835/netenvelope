package server

import (
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:generate npm run build
//go:embed static
var staticFS embed.FS

type Server struct {
	echoInst *echo.Echo
}

func (s *Server) index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.gohtml", nil)
}

func (s *Server) registerRoutes() {
	s.echoInst.GET("/", s.index)
	s.echoInst.GET("/static/*", echo.WrapHandler(http.FileServer(http.FS(staticFS))))
}

func (s *Server) Start(bindAddr string) error {
	s.registerRoutes()

	return s.echoInst.Start(bindAddr)
}

func New() *Server {
	echoInst := echo.New()
	echoInst.Renderer = NewEmbeddedTemplater()

	return &Server{
		echoInst: echoInst,
	}
}
