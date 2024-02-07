package server

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	glog "github.com/labstack/gommon/log"
	"github.com/rs/zerolog/log"
	"github.com/ziflex/lecho/v3"
)

//go:generate npm run build
//go:embed static
var staticFS embed.FS

type Config struct {
	BindAddr      string
	SessionSecret string
}

type Server struct {
	config Config

	echoInst *echo.Echo
}

func (s *Server) index(c echo.Context) error {
	sess := getSession(c)
	sess.Values["test"] = "test"
	err := sess.Save(c.Request(), c.Response())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error saving session: %s", err))
	}

	return c.Render(http.StatusOK, "index.gohtml", nil)
}

func (s *Server) registerRoutes() {
	s.echoInst.GET("/", s.index)
	s.echoInst.GET("/static/*", echo.WrapHandler(http.FileServer(http.FS(staticFS))))
}

func (s *Server) Start() error {
	s.registerRoutes()

	return s.echoInst.Start(s.config.BindAddr)
}

func New(config Config) *Server {
	echoInst := echo.New()
	echoInst.Renderer = NewEmbeddedTemplater()

	logger := lecho.From(log.Logger, lecho.WithLevel(glog.INFO))
	echoInst.Logger = logger
	echoInst.Use(lecho.Middleware(lecho.Config{Logger: logger}))

	echoInst.Use(session.Middleware(sessions.NewCookieStore([]byte(config.SessionSecret))))

	return &Server{
		config:   config,
		echoInst: echoInst,
	}
}
