package user

import (
	"Hybrid-blog/constant"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"Hybrid-blog/helpers/utility"
	"Hybrid-blog/models"
	"Hybrid-blog/sessions"
	usecase "Hybrid-blog/usecases"
)

type UserDetail interface {
	TemplateUserDetail(c *gin.Context)
	EditUserDetail(c *gin.Context)
}

type handlerUserDetail struct {
	useCase    usecase.UserCaseUser
	partBranch usecase.UserCaseBranchDepart
	image      usecase.UserCaseImage
}

func NewHandlerUserDetail() UserDetail {
	return &handlerUserDetail{
		useCase:    usecase.NewuserCase(),
		partBranch: usecase.NewUCaseBranchDepart(),
		image:      usecase.NewUserCaseImage(),
	}
}

func (h *handlerUserDetail) TemplateUserDetail(c *gin.Context) {
	userInfo, _ := sessions.GetInfoUserSession()
	userDetail, _ := h.useCase.GetInfoUserDetails(userInfo.Email)
	departments, branch, position := h.getbranchPartPosi()
	c.HTML(http.StatusOK, "user-detail.html", gin.H{
		"title":       "User detail",
		"user":        userInfo,
		"infoUser":    userDetail,
		"departments": departments,
		"branch":      branch,
		"position":    position,
	})
}

func (h *handlerUserDetail) EditUserDetail(c *gin.Context) {
	c.Request.ParseForm()
	userInfo, _ := sessions.GetInfoUserSession()
	userid := c.PostForm("userID")
	name := c.PostForm("name")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	fb := c.PostForm("facebook")
	birthday := c.PostForm("birthday")
	description := c.PostForm("description")
	position := c.PostForm("position")
	branch := c.PostForm("branch")
	department := c.PostForm("department")
	// file, header, errFile := c.Request.FormFile("imageU")
	file, errFile := c.FormFile("imageU")
	var userEdit models.UserEdit
	if errFile != nil {
		if userInfo.Email == string(email) {
			userCurrent, _ := h.useCase.GetInfoUserDetails(string(email))
			if userCurrent != nil {
				userEdit.ImageID = int64(userCurrent.ImageID)
			}
		} else {
			userCurrent, _ := h.useCase.GetInfoUserDetails(string(userInfo.Email))
			if userCurrent != nil {
				userEdit.ImageID = int64(userCurrent.ImageID)
			}
		}
	} else {
		if utility.CheckTypeImage(file) {
			url, _ := utility.SaveFile(file, 0)
			// NameFile := utility.NameRandomImage(header)
			// s3 := utility.NewS3(file, 0, NameFile)
			// s3.UploadImage(c)
			// url := s3.LinkPath
			if url != "" {
				var image = models.Image{
					Url: url,
				}
				h.image.CreateImage(&image)
				id := strconv.FormatUint(uint64(image.ID), 10)
				ImageID, _ := strconv.ParseInt(id, 10, 64)
				userEdit.ImageID = ImageID
			}
		} else {
			c.JSON(http.StatusAlreadyReported, gin.H{
				constant.JSONError: constant.NoImages,
			})
		}
	}
	if name == "" || email == "" {
		c.JSON(http.StatusAlreadyReported, gin.H{
			constant.JSONError: constant.ErrorEmailPasswordEmpty,
		})
		return
	}
	userID, erruserID := strconv.ParseInt(userid, 0, 64)
	positionID, errpositionID := strconv.ParseInt(position, 0, 64)
	branchID, errbranchID := strconv.ParseInt(branch, 0, 64)
	departmentID, errdepartmentID := strconv.ParseInt(department, 0, 64)
	if erruserID != nil ||
		errpositionID != nil ||
		errbranchID != nil ||
		errdepartmentID != nil {
	}
	userEdit.FullName = name
	userEdit.Email = strings.ToLower(email)
	userEdit.BranchID = branchID
	userEdit.DepartmentID = departmentID
	userEdit.PositionID = positionID
	userEdit.Phone = phone
	userEdit.Birthday = birthday
	userEdit.Description = description
	userEdit.Facebook = fb
	user, _ := h.useCase.FindUserByEmail(email)
	if user == nil {
		editUser(h, userID, userEdit, c, userInfo)
	} else {
		if user.ID == uint(userID) {
			editUser(h, userID, userEdit, c, userInfo)
		} else {
			c.JSON(http.StatusOK, gin.H{
				constant.JSONError: constant.DuplicateEmail,
			})
			return
		}
	}
}

func editUser(h *handlerUserDetail, userID int64, userEdit models.UserEdit, c *gin.Context, userInfo *sessions.InfoUserSession) {
	ok, _ := h.editUser(userID, userEdit)
	if !ok {
		c.JSON(http.StatusAlreadyReported, gin.H{
			constant.JSONError: constant.ErrorEditUser,
		})
	} else {
		// session, _ := sessions.GetSession().Get(c.Request, sessions.IDClient)
		user, _ := h.useCase.GetInfoUserDetails(userEdit.Email)
		// session.Values[constant.SessionAvatar] = user.ImageUrl
		// session.Values[constant.SessionName] = "sssss"
		// session.Save(c.Request, c.Writer)
		// userInfo, _ := sessions.GetInfoUserSession()
		// fmt.Print(userInfo)
		var role = ""
		if user != nil {
			userID := fmt.Sprintf("%d", user.UserID)
			if *user.RoleUser == true {
				role = constant.DefaultRoleAdmin
			} else {
				role = constant.DefaultRoleUser
			}
			sessions.UpdateInfoUserSession(user.ImageUrl, user.Email, user.FullName, role, userID, c)
		}
		if userInfo.Email != userEdit.Email {
			c.JSON(http.StatusAlreadyReported, gin.H{
				constant.JSONSuccess: constant.NotificationEditUser,
				constant.JSONUrl:     "/logout",
			})
		} else {
			c.JSON(http.StatusAlreadyReported, gin.H{
				constant.JSONSuccess: constant.NotificationEditUser,
				constant.JSONUrl:     "/user/detail",
			})
		}
	}
}

func (h *handlerUserDetail) getbranchPartPosi() ([]*models.Department, []*models.Branch, []*models.Position) {
	departments, errdepartments := h.partBranch.GetDepartments()
	position, errposition := h.partBranch.GetPositions()
	branch, errbranch := h.partBranch.Getbranch()

	if errdepartments != nil {
		log.Print(errdepartments)
		return nil, branch, position
	} else if errposition != nil {
		log.Print(errposition)
		return departments, branch, nil
	} else if errbranch != nil {
		log.Print(errbranch)
		return departments, nil, position
	}

	return departments, branch, position
}

func (h *handlerUserDetail) editUser(userId int64, user models.UserEdit) (bool, error) {
	err := h.useCase.EditUserDetail(userId, user)
	if err != nil {
		return false, err
	}
	return true, nil
}
