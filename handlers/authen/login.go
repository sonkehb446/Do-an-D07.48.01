package authen

import (
	"Hybrid-blog/constant"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"Hybrid-blog/sessions"
	usecase "Hybrid-blog/usecases"
)

type Authenticator interface {
	TemplateLogin(c *gin.Context)
	Logout(c *gin.Context)
	Login(c *gin.Context)
}

type handlerAuthenticator struct {
	useCase usecase.UserCaseUser
}

func NewHandlerAuthenticator() Authenticator {
	return &handlerAuthenticator{
		useCase: usecase.NewuserCase(),
	}
}

func (u *handlerAuthenticator) TemplateLogin(c *gin.Context) {
	session, _ := sessions.GetSession().Get(c.Request, sessions.IDClient)
	_, ok := session.Values[constant.SessionUserID]
	if ok {
		http.Redirect(c.Writer, c.Request, constant.URLHome, http.StatusSeeOther)
		return
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Login",
		})
		return
	}
}

func (u *handlerAuthenticator) Login(c *gin.Context) {
	email := c.PostForm(constant.FormEmail)
	password := c.PostForm(constant.FormPassword)
	var role = ""
	if email == "" || password == "" {
		c.JSON(http.StatusAlreadyReported, gin.H{
			constant.JSONError: constant.ErrorEmailPasswordEmpty,
		})
	} else {
		user, ok := u.useCase.IsLogin(strings.ToLower(email), password)
		if !ok {
			c.JSON(http.StatusAlreadyReported, gin.H{
				constant.JSONError: constant.ErrorWrongEmailPassword,
			})
		} else {
			if user != nil {
				userID := fmt.Sprintf("%d", user.UserID)
				if *user.RoleUser == true {
					role = constant.DefaultRoleAdmin
				} else {
					role = constant.DefaultRoleUser
				}
				_ = sessions.NewCookies(c, userID, user.FullName, user.Email, user.Url, role)
			}
			c.JSON(http.StatusOK, gin.H{
				constant.JSONUrl: constant.URLHome,
			})
		}
	}
}

func (u *handlerAuthenticator) Logout(c *gin.Context) {
	sessions.ClearSession(c)
	http.Redirect(c.Writer, c.Request, constant.URLLogin, http.StatusFound)
}
