package app

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type SessionModel struct {
	Username string
	FullName string
	Role     string
}

func SetSession(c echo.Context, model SessionModel) (sess *sessions.Session) {
	sess, _ = session.Get("session", c)
	sess.Values["username"] = model.Username
	sess.Values["fullName"] = model.FullName
	sess.Values["role"] = model.Role
	sess.Save(c.Request(), c.Response())
	return
}

func GetSession(c echo.Context) (sess *sessions.Session, model SessionModel) {
	sess, _ = session.Get("session", c)
	model.Username = sess.Values["username"].(string)
	model.FullName = sess.Values["fullName"].(string)
	model.Role = sess.Values["role"].(string)
	return
}
