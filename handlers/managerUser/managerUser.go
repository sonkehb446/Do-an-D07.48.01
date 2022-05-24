package managerUser

import (
	"Hybrid-blog/constant"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"Hybrid-blog/helpers/utility"
	"Hybrid-blog/models"
	"Hybrid-blog/sessions"
	usecase "Hybrid-blog/usecases"
)

type ManagerUser interface {
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	EditUser(c *gin.Context)
	TemplateUsers(c *gin.Context)
	TemplateCreateUser(c *gin.Context)
}

type SearchUser struct {
	Email  string
	Name   string
	AtFrom string
	AtTo   string
}

type PageUserInfo struct {
	Search      SearchUser
	TotalPage   int
	CurrentPage int
	CPMI        []int
	CPAJ        []int
	TPM5        int
}

type handlerManagerUser struct {
	useCase usecase.UserCaseUser
}

func NewHandlerManagerUser() ManagerUser {
	return &handlerManagerUser{
		useCase: usecase.NewuserCase(),
	}
}

func (m *handlerManagerUser) TemplateCreateUser(c *gin.Context) {
	userInfo, _ := sessions.GetInfoUserSession()
	if userInfo.Role == constant.DefaultRoleAdmin {
		c.HTML(http.StatusOK, "createUser.html", gin.H{
			"title": "Create user",
			"user":  userInfo,
		})
	}
}

func (h *handlerManagerUser) CreateUser(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	fullName := c.PostForm("fullName")
	role := c.PostForm("role")
	if email == "" || password == "" || fullName == "" {
		c.JSON(http.StatusAlreadyReported, gin.H{
			constant.JSONError: constant.RequestCreateUserNotEmpty,
		})
		return
	} else {
		lengthEmail := utility.CheckLength(email, 255)
		lengthPassword := utility.CheckLength(password, 255)
		lengthName := utility.CheckLength(fullName, 150)
		if !lengthEmail {
			c.JSON(http.StatusAlreadyReported, gin.H{
				constant.JSONError: constant.RequestLengthEmail,
			})
			return
		}
		if !lengthPassword {
			c.JSON(http.StatusAlreadyReported, gin.H{
				constant.JSONError: constant.RequestLengthPassword,
			})
			return
		}
		if !lengthName {
			c.JSON(http.StatusAlreadyReported, gin.H{
				constant.JSONError: constant.RequestLengthName,
			})
			return
		}
	}
	ps, HashPass := utility.HashPassword(password)
	if HashPass != nil {
		log.Println("Error HashPassword", HashPass)
		return
	}
	userInfo, _ := sessions.GetInfoUserSession()
	user, _ := h.useCase.FindUserByEmail(email)
	if user == nil && userInfo.Role == constant.DefaultRoleAdmin {
		userid := fmt.Sprintf("%s", userInfo.UserID)
		createby, errcreateby := strconv.ParseInt(userid, 0, 64)
		if errcreateby != nil {
			log.Println(errcreateby)
		}
		var imageID = 1
		var roleUser bool
		if role == constant.DefaultRoleAdmin {
			roleUser = true
		}
		var user = models.User{
			Email:    strings.ToLower(email),
			Password: ps,
			ImageID:  &imageID,
			FullName: fullName,
			CreateBy: &createby,
			RoleUser: &roleUser,
		}
		ok, err := h.useCase.CreateUser(user)
		if err != nil {
			c.JSON(http.StatusAlreadyReported, gin.H{
				constant.JSONError: constant.ErrorCreateUser,
			})
			log.Println(err)
			return
		}
		if ok {
			c.JSON(http.StatusOK, gin.H{
				constant.JSONSuccess: constant.NotificationCreateUser,
				constant.JSONUrl:     constant.URLUsers,
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			constant.JSONError: constant.DuplicateEmail,
		})
		return
	}
}

func (h *handlerManagerUser) EditUser(c *gin.Context) {
	userInfo, _ := sessions.GetInfoUserSession()
	id := c.Param("id")
	useName := c.PostForm("NameEdit")
	email := c.PostForm("EmailEdit")
	role := c.PostForm("genderS")
	log.Println("role", role)
	if useName != "" && id != "" && email != "" {
		ID, _ := strconv.ParseInt(id, 0, 64)
		user, _ := h.useCase.FindUserByEmail(email)
		if user == nil {
			editUser(h, role, ID, email, useName, userInfo, id, c)
		} else {
			if user.ID == uint(ID) {
				editUser(h, role, ID, email, useName, userInfo, id, c)
			} else {
				c.JSON(http.StatusOK, gin.H{
					constant.JSONError: constant.DuplicateEmail,
				})
				return
			}
		}
	} else {
		c.JSON(http.StatusAlreadyReported, gin.H{
			constant.JSONError: constant.InformationNotEmpty,
		})
		return
	}
}

func editUser(h *handlerManagerUser, role string, ID int64, email string, useName string, userInfo *sessions.InfoUserSession, id string, c *gin.Context) {
	errEdit := h.useCase.EditUserByID(ID, strings.ToLower(email), useName, role)
	if errEdit == nil {
		var url = ""
		if userInfo.UserID == id {
			if userInfo.Email != email {
				url = "/logout"
			}
		}
		c.JSON(http.StatusOK, gin.H{
			constant.JSONSuccess: constant.NotificationEditUser,
			constant.JSONUrl:     url,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			constant.JSONError: constant.ErrorEditUser,
		})

	}
}

func (h *handlerManagerUser) DeleteUser(c *gin.Context) {
	userInfo, _ := sessions.GetInfoUserSession()
	id := c.Param("id")
	userid := fmt.Sprintf("%s", userInfo.UserID)
	delete := fmt.Sprintf("%s", id)
	deleteU, _ := strconv.Atoi(delete)
	if id != "" {
		if userid == delete {
			c.JSON(http.StatusAlreadyReported, gin.H{
				constant.JSONError: "Can't delete myself",
			})
			return
		} else {
			errDelete := h.useCase.DeletebyID(deleteU)
			if errDelete == nil {
				c.JSON(http.StatusOK, gin.H{
					constant.JSONSuccess: constant.NotificationDeleteUser,
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					constant.JSONError: constant.ErrorDeleteUser,
				})
				return
			}
		}
	} else {
		c.JSON(http.StatusAlreadyReported, gin.H{
			constant.JSONError: constant.IDNotFound,
		})
		return
	}
}

func (h *handlerManagerUser) TemplateUsers(c *gin.Context) {
	var (
		countQuery        int64
		users             []*models.UserDetail
		err               error
		search            SearchUser
		currentPage       int
		totalPage         int
		currentPageMinusI []int
		currentPageAddJ   []int
		totalPageMinus5   int
		maxPage           int
	)
	userName := c.DefaultQuery("UserName", "")
	email := c.DefaultQuery("Email", "")
	atFrom := c.DefaultQuery("AtFrom", "")
	atTo := c.DefaultQuery("AtTo", "")
	offset := c.DefaultQuery("page", "1")
	pageNumber, setErr := strconv.Atoi(offset)
	if setErr != nil {
		log.Println("Error", setErr.Error())
	}
	_, countQuery, _ = h.useCase.GetAllUser(email, userName, atFrom, atTo, int64(pageNumber-1)*constant.LimitPages, constant.LimitPages)
	sumPage := int(math.Ceil(float64(countQuery) / float64(constant.LimitPages)))
	if pageNumber >= sumPage {
		maxPage = sumPage
	} else if pageNumber <= 1 {
		maxPage = 1
	} else {
		maxPage = pageNumber
	}
	if userName != "" || email != "" || atFrom != "" || atTo != "" {
		users, _, err = h.useCase.GetAllUser(email, userName, atFrom, atTo, int64(maxPage-1)*constant.LimitPages, constant.LimitPages)
		search.Name = userName
		search.Email = email
		search.AtTo = atTo
		search.AtFrom = atFrom
	} else {
		users, _, err = h.useCase.GetAllUser(email, userName, atFrom, atTo, int64(maxPage-1)*constant.LimitPages, constant.LimitPages)
		log.Println(countQuery)
	}
	if err == nil {
		currentPage, totalPage,
			currentPageMinusI, currentPageAddJ,
			totalPageMinus5 = utility.CalculatePage(currentPage, maxPage, countQuery, constant.LimitPages)
		pageInfo := &PageUserInfo{
			Search:      search,
			TotalPage:   totalPage,
			CurrentPage: currentPage,
			CPMI:        currentPageMinusI,
			CPAJ:        currentPageAddJ,
			TPM5:        totalPageMinus5,
		}
		user, _ := sessions.GetInfoUserSession()
		c.HTML(http.StatusOK, "user-list.html", gin.H{
			"title":    "List users",
			"user":     user,
			"Users":    users,
			"pageInfo": pageInfo,
		})
	}
}
