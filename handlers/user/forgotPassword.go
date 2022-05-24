package user

import (
	"Hybrid-blog/constant"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"Hybrid-blog/helpers/utility"
	"Hybrid-blog/models"
	usecase "Hybrid-blog/usecases"
)

var Address = constant.AddressResetPassword

type ForgotPwInterface interface {
	ForgotPassword(c *gin.Context)
	TemplateResetPass(c *gin.Context)
	ResetPassword(c *gin.Context)
}

type handlerForgotPassword struct {
	useCase usecase.UserCaseUser
}

func NewHandlerForgotPassword() ForgotPwInterface {
	return &handlerForgotPassword{
		useCase: usecase.NewuserCase(),
	}
}

func (fw *handlerForgotPassword) ForgotPassword(c *gin.Context) {
	email := c.PostForm(constant.FormEmail)
	if email == "" {
		c.JSON(http.StatusAlreadyReported, gin.H{
			constant.JSONError: constant.ErrorWrongEmail,
		})
	} else {
		user, _ := fw.useCase.FindUserByEmail(email)
		if user != nil {
			fwpass, errFind := fw.useCase.FindFwPassword(int(user.ID))
			MD5Email := utility.HashStringMD5(email)
			link := fmt.Sprintf("%s?&successful=%s&id=%v", Address, MD5Email, user.ID)
			log.Printf("link Email-> %v", link)
			if errFind != nil {
				timeResetPassword := time.Now().Local().Add(time.Hour * time.Duration(24))
				var newResetPass = models.Reset_password{
					UserID:     int64(user.ID),
					Time:       utility.AddTime(5),
					Expired_at: timeResetPassword,
					Token:      MD5Email,
				}
				err := fw.useCase.CreateFwPassword(&newResetPass)
				if err == nil {
					c.JSON(http.StatusOK, gin.H{
						constant.JSONSuccess: constant.NotificationCheckEmail,
					})
					sendEmail(email, link)
					return
				}
			} else {
				timeCurrent := fwpass.Time
				Now := utility.GetTime()
				checkTime := utility.TimeLimit(timeCurrent, Now)
				if checkTime {
					c.JSON(http.StatusOK, gin.H{
						constant.JSONSuccess: constant.NotificationCheckEmail,
					})
					fw.useCase.UpdateFwPassword(fwpass, utility.AddTime(5))
					sendEmail(email, link)
					return
				} else {
					c.JSON(http.StatusOK, gin.H{
						constant.JSONError: constant.WaitTimeSendEmail,
					})
					return
				}
			}
		} else {
			c.JSON(http.StatusAlreadyReported, gin.H{
				constant.JSONError: constant.ErrorWrongEmail,
			})
			return
		}

	}
}

func sendEmail(email string, link string) {
	go func(email string) {
		utility.SendEmail(email, link)
	}(email)
}

func (u *handlerForgotPassword) TemplateResetPass(c *gin.Context) {
	id := c.Query("id")
	sucsse := c.Query("successful")
	check := false
	if id == "" && sucsse == "" {
		return
	} else {
		Id, _ := strconv.Atoi(id)
		fw, err := u.useCase.FindFwPassword(int(Id))
		if err == nil {
			if fw.Token == sucsse && fw.UserID == int64(Id) {
				Time := time.Now()
				TimeCurrent := fw.Expired_at
				if Time.Before(TimeCurrent) {
					c.HTML(http.StatusOK, "resetPassword.html", gin.H{
						"title": "Reset password",
					})
				} else {
					check = true
				}
			} else {
				check = true
			}
		} else {
			check = true
		}
	}
	if check {
		c.JSON(http.StatusAlreadyReported, gin.H{
			constant.JSONError: constant.TokenExpiration,
		})
	}
}

func (u *handlerForgotPassword) ResetPassword(c *gin.Context) {
	id := c.Query("id")
	sucsse := c.Query("successful")
	pw := c.PostForm("password")

	if id == "" && sucsse == "" {
		return
	} else {
		Id, _ := strconv.Atoi(id)
		fw, err := u.useCase.FindFwPassword(int(Id))
		if err == nil {
			if fw.Token == sucsse && fw.UserID == int64(Id) {
				Time := time.Now()
				TimeCurrent := fw.Expired_at
				if Time.Before(TimeCurrent) {
					if pw == "" {
						c.JSON(http.StatusAlreadyReported, gin.H{
							constant.JSONError: constant.RequestPasswordConfirm,
						})
						return
					} else {
						if len(pw) <= 225 {
							hashPw, err := utility.HashPassword(pw)
							if err != nil {
								log.Print(err)
								return
							} else {
								err := u.useCase.DeleteFwPassword(int(fw.ID))
								if err == nil {
									go func(id int, hashPw string) {
										u.resetPassword(id, hashPw)
									}(int(fw.UserID), hashPw)
									c.JSON(http.StatusOK, gin.H{
										constant.JSONSuccess: constant.ChangePasswordSuccessfully,
										constant.JSONUrl:     "http://" + constant.IpAddress + ":" + constant.PortWeb + "/login",
									})
								}
							}
						} else {
							c.JSON(http.StatusAlreadyReported, gin.H{
								constant.JSONError: constant.RequestLengthPassword,
							})
							return
						}

					}
				}
			} else {
				c.JSON(http.StatusAlreadyReported, gin.H{
					constant.JSONError: constant.TokenExpiration,
				})
				return
			}
		}
	}

}

func (u *handlerForgotPassword) resetPassword(id int, password string) bool {
	ok := u.useCase.ResetPassword(id, password)
	if !ok {
		log.Print("Error resetPassword")
	}
	return ok
}
