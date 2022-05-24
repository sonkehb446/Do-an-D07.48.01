package user

import (
	"Hybrid-blog/constant"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"Hybrid-blog/helpers/utility"
	"Hybrid-blog/sessions"
	usecase "Hybrid-blog/usecases"
)

type ChangePwInterface interface {
	TemplateChangePassword(c *gin.Context)
	ChangePassword(c *gin.Context)
}

type handlerChangePass struct {
	useCase usecase.UserCaseUser
}

func NewHandlerChangePass() ChangePwInterface {
	return &handlerChangePass{
		useCase: usecase.NewuserCase(),
	}
}
func (h *handlerChangePass) TemplateChangePassword(c *gin.Context) {
	user, _ := sessions.GetInfoUserSession()
	c.HTML(http.StatusOK, "change-password.html", gin.H{
		"title": "Change Password",
		"user":  user,
	})
}
func (h *handlerChangePass) ChangePassword(c *gin.Context) {

	user, _ := sessions.GetInfoUserSession()
	newPassword := c.PostForm("password")
	currentPassword := c.PostForm("currentpw")
	var gmail = fmt.Sprintf("%v", user.Email)
	if currentPassword != "" && newPassword != "" {
		if len(newPassword) <= 225 {
			hashPw, err := utility.HashPassword(newPassword)
			if err != nil {
				return
			}
			user, ok := h.useCase.IsLogin(gmail, currentPassword)
			if ok {
				changePassword := h.useCase.ResetPassword(user.UserID, hashPw)
				if changePassword {
					c.JSON(http.StatusOK, gin.H{
						constant.JSONSuccess: constant.ChangePasswordSuccessfully,
						constant.JSONUrl:     "/logout",
					})
				} else {
					c.JSON(http.StatusAlreadyReported, gin.H{
						constant.JSONError: constant.ErrorChangePassword,
					})
					return
				}
			} else {
				c.JSON(http.StatusAlreadyReported, gin.H{
					constant.JSONError: constant.RequestOldPassword,
				})
				return
			}
		} else {

			c.JSON(http.StatusAlreadyReported, gin.H{
				constant.JSONError: constant.RequestLengthPassword,
			})
			return
		}
	}
}
