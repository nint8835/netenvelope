package server

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	glog "github.com/labstack/gommon/log"
	"github.com/rs/zerolog/log"
	"github.com/ziflex/lecho/v3"
	"golang.org/x/crypto/bcrypt"

	"github.com/nint8835/netenvelope/pkg/database"
	"github.com/nint8835/netenvelope/pkg/database/queries"
	"github.com/nint8835/netenvelope/pkg/server/ui/pages"
)

//go:generate npm run build
//go:embed static
//nolint:typecheck
var staticFS embed.FS

type Config struct {
	BindAddr      string
	SessionSecret string
	DbPath        string
}

type Server struct {
	config   Config
	queries  *queries.Queries
	echoInst *echo.Echo
}

func (s *Server) index(c echo.Context) error {
	return s.renderComponent(c, http.StatusOK, pages.Home())
}

func (s *Server) loginPage(c echo.Context) error {
	currentUser := s.getCurrentUser(c)
	if currentUser != nil {
		return c.Redirect(http.StatusFound, "/")
	}

	return s.renderComponent(c, http.StatusOK, pages.Login(""))
}

type loginFormBody struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (s *Server) login(c echo.Context) error {
	var form loginFormBody

	if err := c.Bind(&form); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user, err := s.queries.GetUserByUsername(c.Request().Context(), form.Username)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Error().Err(err).Msg("Error getting user by username")
		}

		return s.renderComponent(c, http.StatusOK, pages.Login("Invalid username or password"))
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(form.Password))
	if err != nil {
		return s.renderComponent(c, http.StatusOK, pages.Login("Invalid username or password"))
	}

	sess := getSession(c)
	sess.Values[userSessionKey] = user.ID
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error saving session: %s", err))
	}

	return c.Redirect(http.StatusFound, "/")
}

func (s *Server) logout(c echo.Context) error {
	sess := getSession(c)
	delete(sess.Values, userSessionKey)
	err := sess.Save(c.Request(), c.Response())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error saving session: %s", err))
	}

	return c.Redirect(http.StatusFound, "/")
}

func (s *Server) registerRoutes() {
	s.echoInst.GET("/", s.index)

	s.echoInst.GET("/login", s.loginPage)
	s.echoInst.POST("/login", s.login)
	s.echoInst.GET("/logout", s.logout)

	s.echoInst.GET("/static/*", echo.WrapHandler(http.FileServer(http.FS(staticFS))))
}

func (s *Server) Start() error {
	s.registerRoutes()

	return s.echoInst.Start(s.config.BindAddr)
}

func New(config Config) (*Server, error) {
	echoInst := echo.New()
	renderer, err := NewEmbeddedTemplater()
	if err != nil {
		return nil, fmt.Errorf("failed to create renderer: %w", err)
	}

	echoInst.Renderer = renderer

	logger := lecho.From(log.Logger, lecho.WithLevel(glog.INFO))
	echoInst.Logger = logger
	echoInst.Use(lecho.Middleware(lecho.Config{Logger: logger}))

	echoInst.Use(session.Middleware(sessions.NewCookieStore([]byte(config.SessionSecret))))

	dbInst, err := database.New(config.DbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	return &Server{
		config:   config,
		queries:  dbInst,
		echoInst: echoInst,
	}, nil
}
