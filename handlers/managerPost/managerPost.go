package managerpost

import (
	"Hybrid-blog/constant"
	img "Hybrid-blog/handlers/image"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"Hybrid-blog/helpers/utility"
	"Hybrid-blog/models"
	"Hybrid-blog/sessions"
	usecase "Hybrid-blog/usecases"
)

type ManagerPost interface {
	TemplateCreatePost(c *gin.Context)
	TemplateEditPost(c *gin.Context)
	CreatePost(c *gin.Context)
	TemplatePosts(c *gin.Context)
	DeletePost(c *gin.Context)
	DetailPost(c *gin.Context)
	EditPost(c *gin.Context)
	TemplateHyBirdNew(c *gin.Context)
	TemplateHyBirders(c *gin.Context)
	TemplateHyLearning(c *gin.Context)
	TemplatePostsByCategory(c *gin.Context)
	TemplatePostsBySubCategory(c *gin.Context)
}

type handlerManagerPost struct {
	useCase  usecase.UserCasePost
	category usecase.CategoryUseCase
	usImage  usecase.UserCaseImage
}

type PostSearch struct {
	Title       string
	Description string
	Create_By   string
	Create_from string
	Create_to   string
}

type PostbyCategory struct {
	ID   string
	Name string
}

type PagePostInfo struct {
	Search       PostSearch
	PostCategory PostbyCategory
	TotalPage    int
	CurrentPage  int
	CPMI         []int
	CPAJ         []int
	TPM5         int
}

func NewHandlerManagerPost() ManagerPost {
	return &handlerManagerPost{
		useCase:  usecase.NewPostUseCase(),
		usImage:  usecase.NewUserCaseImage(),
		category: usecase.NewCategoryUseCase(),
	}
}

func (h *handlerManagerPost) TemplateEditPost(c *gin.Context) {
	user, _ := sessions.GetInfoUserSession()
	userID, _ := strconv.ParseInt(user.UserID, 10, 64)
	categories, errDelete := usecase.NewCategoryUseCase().GetAllCategoryNoPagination(c)
	if errDelete != nil {
		log.Println(errDelete)
	} else {
		id := c.Param("id")
		edit := fmt.Sprintf("%s", id)
		editU, _ := strconv.Atoi(edit)
		post, err := h.useCase.FindPostByID(editU)
		if err == nil {
			subCategory, _ := h.category.FindCategoryBySubCategory(post.CategoryID)
			category, _ := h.category.FindCategoryBySubCategory(*subCategory.ParentID)
			if user.Role == constant.DefaultRoleAdmin {
				c.HTML(http.StatusOK, "editPost.html", gin.H{
					"user":          user,
					"category":      categories,
					"tempPost":      post,
					"tempCategory":  category,
					"tempSCategory": subCategory,
				})
			} else if user.Role == constant.DefaultRoleUser {
				if *post.UserID == userID {
					c.HTML(http.StatusOK, "editPost.html", gin.H{
						"user":          user,
						"category":      categories,
						"tempPost":      post,
						"tempCategory":  category,
						"tempSCategory": subCategory,
					})
				} else {
					utility.RedirectError(c, constant.NoPermission)
				}
			}
		}
	}
}

func (h *handlerManagerPost) TemplateCreatePost(c *gin.Context) {
	user, _ := sessions.GetInfoUserSession()
	categories, err := usecase.NewCategoryUseCase().GetAllCategoryNoPagination(c)
	if err != nil {
		log.Println(err)
	}
	c.HTML(http.StatusOK, "createPost.html", gin.H{
		"title":    "Create post",
		"user":     user,
		"category": categories,
	})
}

func (h *handlerManagerPost) TemplateHyBirdNew(c *gin.Context) {
	hynews, pageInfo, errors := h.templateByType(c, constant.DefaultTypeHyNew)
	if errors != nil {
		utility.RedirectError(c, errors)
		return
	}
	user, err := sessions.GetInfoUserSession()
	if err != nil {
		c.HTML(http.StatusOK, "hynews.html", gin.H{
			"posts":    hynews,
			"title":    "Hynews",
			"pageInfo": pageInfo,
		})
	} else {
		c.HTML(http.StatusOK, "hynews.html", gin.H{
			"user":     user,
			"posts":    hynews,
			"title":    "Hynews",
			"pageInfo": pageInfo,
		})
	}
}
func (h *handlerManagerPost) TemplateHyLearning(c *gin.Context) {
	hylearning, pageInfo, errors := h.templateByType(c, constant.DefaultTypeHyLearning)
	if errors != nil {
		utility.RedirectError(c, errors)
		return
	}
	user, err := sessions.GetInfoUserSession()
	if err != nil {
		c.HTML(http.StatusOK, "hynews.html", gin.H{
			"posts":    hylearning,
			"title":    "Hynews",
			"pageInfo": pageInfo,
		})
	} else {
		c.HTML(http.StatusOK, "hynews.html", gin.H{
			"user":     user,
			"posts":    hylearning,
			"title":    "Hynews",
			"pageInfo": pageInfo,
		})
	}
}
func (h *handlerManagerPost) TemplateHyBirders(c *gin.Context) {
	hybirders, pageInfo, errors := h.templateByType(c, constant.DefaultTypeHyBirder)
	if errors != nil {
		utility.RedirectError(c, errors)
		return
	}
	user, err := sessions.GetInfoUserSession()
	if err != nil {
		c.HTML(http.StatusOK, "hynews.html", gin.H{
			"posts":    hybirders,
			"title":    "Hynews",
			"pageInfo": pageInfo,
		})
	} else {
		c.HTML(http.StatusOK, "hynews.html", gin.H{
			"user":     user,
			"posts":    hybirders,
			"title":    "Hynews",
			"pageInfo": pageInfo,
		})
	}
}

func (h *handlerManagerPost) templateByType(c *gin.Context, t int) ([]*models.PostHome, *PagePostInfo, []error) {
	var (
		countQuery        int64
		typePost          []*models.PostHome
		errors            []error
		search            PostSearch
		currentPage       int
		totalPage         int
		currentPageMinusI []int
		currentPageAddJ   []int
		totalPageMinus5   int
		error             error
		maxPage           int
	)
	offset := c.DefaultQuery("page", "1")
	pageNumber, setErr := strconv.Atoi(offset)
	if setErr != nil {
		errors = append(errors, setErr)
	}
	_, countQuery, _ = h.useCase.GetPostByType(t, 0, 0)
	sumPage := int(math.Ceil(float64(countQuery) / float64(constant.LimitPages)))
	if pageNumber >= sumPage {
		maxPage = sumPage
	} else if pageNumber <= 1 {
		maxPage = 1
	} else {
		maxPage = pageNumber
	}
	typePost, _, error = h.useCase.GetPostByType(t, int(maxPage-1)*constant.LimitPages, constant.LimitPages)
	if error != nil {
		errors = append(errors, error)
	}
	currentPage, totalPage,
		currentPageMinusI, currentPageAddJ,
		totalPageMinus5 = utility.CalculatePage(currentPage, maxPage, countQuery, constant.LimitPages)
	pageInfo := &PagePostInfo{
		Search:      search,
		TotalPage:   totalPage,
		CurrentPage: currentPage,
		CPMI:        currentPageMinusI,
		CPAJ:        currentPageAddJ,
		TPM5:        totalPageMinus5,
	}
	return typePost, pageInfo, errors
}

func (h *handlerManagerPost) CreatePost(c *gin.Context) {
	user, _ := sessions.GetInfoUserSession()
	var post models.Post
	subCategory := c.PostForm("categoryID")
	title := c.PostForm("title")
	description := c.PostForm("description")
	content := c.PostForm("context")
	subCategoryID, errParsed := strconv.ParseInt(subCategory, 0, 64)
	userid := fmt.Sprintf("%s", user.UserID)
	UserI1, _ := strconv.ParseInt(userid, 0, 64)
	post.Title = title
	post.Description = description
	post.CategoryID = subCategoryID
	post.Content = content
	post.UserID = UserI1
	if errParsed != nil {
		log.Println("err", errParsed)
	}
	Img := img.NewImages()
	image, errFile := Img.ReturnImages(c, constant.ImageTypeSmallImage, "imageU")
	if errFile == nil {
		h.usImage.CreateImage(image)
		id := strconv.FormatUint(uint64(image.ID), 10)
		ImageID, _ := strconv.ParseInt(id, 10, 64)
		post.ImageID = &ImageID
	} else {
		c.JSON(http.StatusAlreadyReported, gin.H{
			constant.JSONError: constant.ImageNotFormat,
		})
		return
	}
	if err := h.useCase.CreatePost(&post); err != nil {
		c.JSON(http.StatusAlreadyReported, gin.H{
			constant.JSONError: constant.ErrorCreatePost,
		})
		return
	} else {
		url := fmt.Sprintf("/post/hybrid/%v", post.ID)
		c.JSON(http.StatusOK, gin.H{
			constant.JSONSuccess: constant.NotificationCreatePost,
			constant.JSONUrl:     url,
		})
		return
	}
}

func (h *handlerManagerPost) DetailPost(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	ID := int(i)
	if err != nil {
		log.Println(err)
	}
	post, _ := h.useCase.FindPostByID(ID)
	subCategorys, _ := h.category.FindCategoryByID(post.CategoryID)
	subCategory := subCategorys[0]
	nameParent := subCategory.ParentName

	user, err := sessions.GetInfoUserSession()
	if err != nil {
		c.HTML(http.StatusOK, "detailPost.html", gin.H{
			"tempPost":      post,
			"subCategoryID": subCategory.ID,
			"CategoryID":    subCategory.ParentID,
			"categoryName":  *nameParent,
			"title":         post.Title,
		})
	} else {
		c.HTML(http.StatusOK, "detailPost.html", gin.H{
			"tempPost":      post,
			"user":          user,
			"subCategoryID": subCategory.ID,
			"CategoryID":    subCategory.ParentID,
			"categoryName":  *nameParent,
			"title":         post.Title,
		})
	}
}

func (h *handlerManagerPost) TemplatePosts(c *gin.Context) {
	userInfo, _ := sessions.GetInfoUserSession()
	var (
		countQuery        int64
		post              []*models.PostHome
		err               error
		search            PostSearch
		currentPage       int
		totalPage         int
		currentPageMinusI []int
		currentPageAddJ   []int
		totalPageMinus5   int
		maxpage           int
		userID            string
	)
	if userInfo.Role == constant.DefaultRoleAdmin {
		userID = ""
	} else if userInfo.Role == constant.DefaultRoleUser {
		userID = userInfo.UserID
	}
	Title := c.DefaultQuery("Title", "")
	Description := c.DefaultQuery("Description", "")
	Create_By := c.DefaultQuery("Create_By", "")
	Create_to := c.DefaultQuery("Create-to", "")
	Create_from := c.DefaultQuery("Create-from", "")
	offset := c.DefaultQuery("page", "1")
	pageNumber, setErr := strconv.Atoi(offset)
	if setErr != nil {
		log.Println("Err", setErr.Error())
	}
	_, countQuery, _ = h.useCase.GetAllPostByID(userID, Title, Description,
		Create_By, Create_from, Create_to, int(pageNumber-1)*constant.LimitPages,
		constant.LimitPages)
	sumPage := int(math.Ceil(float64(countQuery) / float64(constant.LimitPages)))
	if pageNumber >= sumPage {
		maxpage = sumPage
	} else if pageNumber <= 1 {
		maxpage = 1
	} else {
		maxpage = pageNumber
	}
	if Title != "" || Description != "" || Create_By != "" || Create_from != "" || Create_to != "" {
		post, _, err = h.useCase.GetAllPostByID(userID, Title, Description,
			Create_By, Create_from, Create_to, int(maxpage-1)*constant.LimitPages,
			constant.LimitPages)
		search.Title = Title
		search.Description = Description
		search.Create_By = Create_By
		search.Create_from = Create_from
		search.Create_to = Create_to
	} else {
		post, _, err = h.useCase.GetAllPostByID(userID, Title, Description,
			Create_By, Create_from, Create_to, int(maxpage-1)*constant.LimitPages,
			constant.LimitPages)
	}
	if err == nil {
		currentPage, totalPage,
			currentPageMinusI, currentPageAddJ,
			totalPageMinus5 = utility.CalculatePage(currentPage, maxpage, countQuery, constant.LimitPages)
		pageInfo := &PagePostInfo{
			Search:      search,
			TotalPage:   totalPage,
			CurrentPage: currentPage,
			CPMI:        currentPageMinusI,
			CPAJ:        currentPageAddJ,
			TPM5:        totalPageMinus5,
		}
		c.HTML(http.StatusOK, "post-list.html", gin.H{
			"title":    "List posts",
			"user":     userInfo,
			"posts":    post,
			"pageInfo": pageInfo,
		})
	}
}

func (h *handlerManagerPost) DeletePost(c *gin.Context) {
	id := c.Param("id")
	delete := fmt.Sprintf("%s", id)
	deleteU, _ := strconv.Atoi(delete)
	post, err := h.useCase.FindPostByID(deleteU)

	if err != nil {
		c.JSON(http.StatusAlreadyReported, gin.H{
			constant.JSONError: constant.IDNotFound,
		})
		return
	} else {
		userInfo, _ := sessions.GetInfoUserSession()
		userID, _ := strconv.ParseInt(userInfo.UserID, 10, 64)
		var errDelete error
		if userInfo.Role == constant.DefaultRoleAdmin {
			errDelete = h.useCase.DeletePostByID(deleteU)
		} else if userInfo.Role == constant.DefaultRoleUser {
			log.Print(*post.UserID, userID)
			if *post.UserID == userID {
				errDelete = h.useCase.DeletePostByID(deleteU)
			} else {
				c.JSON(http.StatusOK, gin.H{
					constant.JSONError: constant.NoPermission,
				})
				return
			}
		}
		if errDelete == nil {
			c.JSON(http.StatusOK, gin.H{
				constant.JSONSuccess: constant.NotificationDeletePost,
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				constant.JSONError: constant.ErrorDeletePost,
			})
			return
		}
	}
}

func (h *handlerManagerPost) EditPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	category := c.PostForm("categoryID")
	title := c.PostForm("title")
	description := c.PostForm("description")
	context := c.PostForm("context")
	if id != "" && (title != "" || description != "" || context != "" || category != "") {
		categoryID, err := strconv.Atoi(category)
		ID, _ := strconv.ParseUint(id, 0, 64)
		post.CategoryID = int64(categoryID)
		log.Println(category, err, "Category ID", int64(categoryID))
		post.Title = title
		post.Description = description
		post.Content = context
		post.ID = uint(ID)
		postCurrent, _ := h.useCase.FindPostByID(int(ID))
		post.UserID = *postCurrent.UserID
		Img := img.NewImages()
		image, errFile := Img.ReturnImages(c, constant.ImageTypeSmallImage, "imageU")
		if errFile == nil {
			h.usImage.CreateImage(image)
			id := strconv.FormatUint(uint64(image.ID), 10)
			ImageID, _ := strconv.ParseInt(id, 10, 64)
			post.ImageID = &ImageID
		} else {
			post.ImageID = postCurrent.ImageID
		}
		errEdit := h.useCase.EditPost(&post)
		if errEdit != nil {
			c.JSON(http.StatusAlreadyReported, gin.H{
				constant.JSONError: constant.ErrorEditPost,
			})
			return
		} else {
			url := fmt.Sprintf("/post/hybrid/%v", post.ID)
			c.JSON(http.StatusOK, gin.H{
				constant.JSONSuccess: constant.NotificationEditPost,
				constant.JSONUrl:     url,
			})
			return
		}
	}
}

func (h *handlerManagerPost) TemplatePostsByCategory(c *gin.Context) {
	var (
		countQuery        int64
		hybirders         []*models.PostHome
		err               error
		currentPage       int
		totalPage         int
		currentPageMinusI []int
		currentPageAddJ   []int
		totalPageMinus5   int
		maxpage           int
		category          PostbyCategory
	)
	categoryName := c.DefaultQuery("category", "")
	categoryID := c.DefaultQuery("id", "")
	offset := c.DefaultQuery("page", "1")
	pageNumber, setErr := strconv.Atoi(offset)
	if setErr != nil {
		log.Print(setErr)
	}
	_, countQuery, _ = h.useCase.FindPostByCategoryID(0, categoryName, 0, 0)
	sumPage := int(math.Ceil(float64(countQuery) / float64(constant.LimitPages)))
	if pageNumber >= sumPage {
		maxpage = sumPage
	} else if pageNumber <= 1 {
		maxpage = 1
	} else {
		maxpage = pageNumber
	}
	category.ID = categoryID
	category.Name = categoryName
	hybirders, _, err = h.useCase.FindPostByCategoryID(0, categoryName, int64(maxpage-1)*constant.LimitPages, constant.LimitPages)
	currentPage, totalPage,
		currentPageMinusI, currentPageAddJ,
		totalPageMinus5 = utility.CalculatePage(currentPage, maxpage, countQuery, constant.LimitPages)
	pageInfo := &PagePostInfo{
		PostCategory: category,
		TotalPage:    totalPage,
		CurrentPage:  currentPage,
		CPMI:         currentPageMinusI,
		CPAJ:         currentPageAddJ,
		TPM5:         totalPageMinus5,
	}
	user, err := sessions.GetInfoUserSession()
	if err != nil {
		c.HTML(http.StatusOK, "postbycategory.html", gin.H{
			"posts":    hybirders,
			"title":    categoryName,
			"pageInfo": pageInfo,
		})
	} else {
		c.HTML(http.StatusOK, "postbycategory.html", gin.H{
			"user":     user,
			"posts":    hybirders,
			"title":    categoryName,
			"pageInfo": pageInfo,
		})
	}
}

func (h *handlerManagerPost) TemplatePostsBySubCategory(c *gin.Context) {
	var (
		countQuery        int64
		hybirders         []*models.PostHome
		err               error
		currentPage       int
		totalPage         int
		currentPageMinusI []int
		currentPageAddJ   []int
		totalPageMinus5   int
		maxpage           int
		category          PostbyCategory
	)
	subcategorygoryName := c.DefaultQuery("subcategory", "")
	subcategory := c.Query("id")
	offset := c.DefaultQuery("page", "1")
	pageNumber, setErr := strconv.Atoi(offset)
	subcategoryID, setErr := strconv.Atoi(subcategory)
	if setErr != nil {
		log.Print(setErr)
	}
	_, countQuery, _ = h.useCase.FindPostByCategoryID(subcategoryID, "", 0, 0)
	sumPage := int(math.Ceil(float64(countQuery) / float64(constant.LimitPages)))
	if pageNumber >= sumPage {
		maxpage = sumPage
	} else if pageNumber <= 1 {
		maxpage = 1
	} else {
		maxpage = pageNumber
	}
	category.ID = subcategory
	category.Name = subcategorygoryName
	hybirders, _, err = h.useCase.FindPostByCategoryID(subcategoryID, "", int64(maxpage-1)*constant.LimitPages, constant.LimitPages)
	currentPage, totalPage,
		currentPageMinusI, currentPageAddJ,
		totalPageMinus5 = utility.CalculatePage(currentPage, maxpage, countQuery, constant.LimitPages)
	pageInfo := &PagePostInfo{
		PostCategory: category,
		TotalPage:    totalPage,
		CurrentPage:  currentPage,
		CPMI:         currentPageMinusI,
		CPAJ:         currentPageAddJ,
		TPM5:         totalPageMinus5,
	}
	user, err := sessions.GetInfoUserSession()
	if err != nil {
		c.HTML(http.StatusOK, "postbySubcategory.html", gin.H{
			"posts":    hybirders,
			"title":    subcategorygoryName,
			"pageInfo": pageInfo,
		})
	} else {
		c.HTML(http.StatusOK, "postbySubcategory.html", gin.H{
			"posts":    hybirders,
			"title":    subcategorygoryName,
			"user":     user,
			"pageInfo": pageInfo,
		})
	}
}
