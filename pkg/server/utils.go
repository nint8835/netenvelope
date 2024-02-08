package server

import (
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

//nolint:unused
type globalTemplateContext struct {
	currentUser *queries.User
}

//nolint:unused
func (s *Server) getGlobalTemplateContext(c echo.Context) globalTemplateContext {
	_ = getSession(c)

	currentUser := s.getCurrentUser(c)

	return globalTemplateContext{
		currentUser: currentUser,
	}
}
