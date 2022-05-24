package middleware

import (
	"Hybrid-blog/constant"
	"Hybrid-blog/helpers/utility"
	"Hybrid-blog/sessions"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := sessions.GetSession().Get(c.Request, sessions.IDClient)
		if err != nil {
			link := fmt.Sprintf("/home")
			http.Redirect(c.Writer, c.Request, link, http.StatusSeeOther)
			return
		}
		_, ok := session.Values[constant.SessionRole]
		if ok == false {
			link := fmt.Sprintf("/home")
			http.Redirect(c.Writer, c.Request, link, http.StatusSeeOther)
			return
		}
		c.Next()
	}
}

func MiddleAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := sessions.GetSession().Get(c.Request, sessions.IDClient)
		role, _ := session.Values[constant.SessionRole]
		log.Print("3223", role)
		if role == constant.DefaultRoleUser {
			utility.RedirectError(c, "You do not have permission to access")
			c.Abort()
			return
		}
		c.Next()
	}
}
