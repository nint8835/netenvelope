package server

import (
	"context"

	"github.com/a-h/templ"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/nint8835/netenvelope/pkg/database/queries"
)

var userSessionKey = "user"

func getSession(c echo.Context) *sessions.Session {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   24 * 60 * 60,
		HttpOnly: true,
	}

	return sess
}

func (s *Server) getCurrentUser(c echo.Context) *queries.User {
	sess := getSession(c)
	userIdInterface, hasUser := sess.Values[userSessionKey]
	if !hasUser {
		return nil
	}

	userId := userIdInterface.(int64)

	user, err := s.queries.GetUserById(c.Request().Context(), userId)
	if err != nil {
		log.Error().Err(err).Msg("Error getting current user")
		return nil
	}

	return &user
}

func (s *Server) getRenderContext(c echo.Context) context.Context {
	_ = getSession(c)

	currentUser := s.getCurrentUser(c)

	//nolint:staticcheck
	return context.WithValue(c.Request().Context(), "current_user", currentUser)
}

func (s *Server) renderComponent(c echo.Context, status int, component templ.Component) error {
	c.Response().Writer.WriteHeader(status)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return component.Render(s.getRenderContext(c), c.Response().Writer)
}

func (s *Server) requireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		currentUser := s.getCurrentUser(c)
		if currentUser == nil {
			return c.Redirect(302, "/login")
		}

		return next(c)
	}
}
