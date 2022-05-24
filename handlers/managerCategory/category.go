package managerCategory

import (
	"Hybrid-blog/constant"
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

var _ CategoryInterface = (*categoryHandler)(nil)

type CategoryInterface interface {
	ListCategory(c *gin.Context)
	AddCategory(c *gin.Context)
	CreateCategory(c *gin.Context)
	EditCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	ListSubCategory(c *gin.Context)
	AddSubCategory(c *gin.Context)
	CreateSubCategory(c *gin.Context)
	EditSubCategory(c *gin.Context)
	UpdateSubCategory(c *gin.Context)
	DeleteSubCategory(c *gin.Context)
	ChooseSubCategory(c *gin.Context)
}

type categoryHandler struct {
	useCase usecase.CategoryUseCase
}

func NewCategoryHandler() CategoryInterface {
	return &categoryHandler{
		useCase: usecase.NewCategoryUseCase(),
	}
}

type SearchCategory struct {
	Name        string
	Description string
	CreatedFrom string
	CreatedTo   string
	NameParent  string
}

type PageCategoryInfo struct {
	Search      SearchCategory
	TotalPage   int
	CurrentPage int
	CPMI        []int
	CPAJ        []int
	TPM5        int
}

func (h *categoryHandler) ListCategory(c *gin.Context) {
	var (
		countQuery        int64
		categories        []*models.Category
		err               error
		search            SearchCategory
		currentPage       int
		totalPage         int
		currentPageMinusI []int
		currentPageAddJ   []int
		totalPageMinus5   int
		maxPage           int
	)
	Name := c.DefaultQuery("NameCategory", "")
	Description := c.DefaultQuery("Description", "")
	CreatedTo := c.DefaultQuery("CreatedTo", "")
	CreatedFrom := c.DefaultQuery("CreatedFrom", "")
	offset := c.DefaultQuery("page", "1")
	pageNumber, setErr := strconv.Atoi(offset)
	if setErr != nil {
		log.Println("Err", setErr.Error())
	}
	_, countQuery, _ = usecase.NewCategoryUseCase().GetAllCategory(Name, Description, CreatedFrom, CreatedTo, int64(pageNumber-1)*constant.LimitPages, constant.LimitPages)
	sumPage := int(math.Ceil(float64(countQuery) / float64(constant.LimitPages)))
	if pageNumber >= sumPage {
		maxPage = sumPage
	} else if pageNumber <= 1 {
		maxPage = 1
	} else {
		maxPage = pageNumber
	}
	categories, _, err = usecase.NewCategoryUseCase().GetAllCategory(Name, Description, CreatedFrom, CreatedTo, int64(maxPage-1)*constant.LimitPages, constant.LimitPages)
	if err != nil {
		log.Println("Err", err.Error())
	}
	search.Name = Name
	search.Description = Description
	search.CreatedFrom = CreatedFrom
	search.CreatedTo = CreatedTo
	currentPage, totalPage,
		currentPageMinusI, currentPageAddJ,
		totalPageMinus5 = utility.CalculatePage(currentPage, maxPage, countQuery, constant.LimitPages)
	pageInfo := &PageCategoryInfo{
		Search:      search,
		TotalPage:   totalPage,
		CurrentPage: currentPage,
		CPMI:        currentPageMinusI,
		CPAJ:        currentPageAddJ,
		TPM5:        totalPageMinus5,
	}
	user, _ := sessions.GetInfoUserSession()
	c.HTML(http.StatusOK, "list_category.html", gin.H{
		"title":    "List Categories",
		"category": categories,
		"user":     user,
		"pageInfo": pageInfo,
	})
}

func (h *categoryHandler) AddCategory(c *gin.Context) {
	user, _ := sessions.GetInfoUserSession()
	allMenu, err := usecase.NewMenuUseCase(c).GetAllMenu(c)
	if err != nil {
		log.Println("Err", err.Error())
	}
	c.HTML(http.StatusOK, "create_category.html", gin.H{
		"title": "Create category",
		"menus": allMenu,
		"user":  user,
	})
}
func (h *categoryHandler) CreateCategory(c *gin.Context) {
	name := c.PostForm("name")
	description := c.PostForm("description")
	MenuID := c.PostForm("menuID")
	menuID, err := strconv.Atoi(MenuID)
	if err != nil {
		log.Println(err)
	}
	category := models.Category{MenuID: menuID, CategoryName: name, Description: description}
	err1 := h.useCase.CreateCategory(category)
	if err1 == nil {
		c.JSON(http.StatusOK, gin.H{
			constant.JSONSuccess: constant.NotificationCreateSuccessCategory,
			constant.JSONUrl:     constant.URLCategories,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			constant.JSONError: constant.ErrorCreateCategory,
		})
		return
	}
}
func (h *categoryHandler) EditCategory(c *gin.Context) {
	user, _ := sessions.GetInfoUserSession()
	ID := c.Param("ID")
	id, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		log.Println(err)
	}
	category, err := usecase.NewCategoryUseCase().FindCategoryByID(id)
	MenuID := category[0].MenuID
	menu, err := usecase.NewMenuUseCase(c).FindMenuByID(MenuID)
	menuName := menu[0].Name
	allMenu, err := usecase.NewMenuUseCase(c).GetAllMenu(c)
	if err != nil {
		log.Println(err)
	}
	c.HTML(http.StatusOK, "edit_category.html", gin.H{
		"title":      "Edit category",
		"CategoryID": id,
		"menus":      allMenu,
		"user":       user,
		"category":   category,
		"menuName":   menuName,
		"MenuID":     MenuID,
	})
}

func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	ID := c.Param("ID")
	id, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		log.Println(err)
	}
	name := c.PostForm("name")
	description := c.PostForm("description")
	MenuID := c.PostForm("menuID")
	menuID, err1 := strconv.Atoi(MenuID)
	if err1 != nil {
		log.Println(err1)
	}
	category := models.Category{MenuID: menuID, CategoryName: name, Description: description}
	err2 := h.useCase.UpdateCategory(&category, id)
	if err2 == nil {
		c.JSON(http.StatusOK, gin.H{
			constant.JSONSuccess: constant.NotificationEditSuccessCategory,
			constant.JSONUrl:     constant.URLCategories,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			constant.JSONError: constant.ErrorEditCategory,
		})
		return
	}
}
func (h *categoryHandler) DeleteCategory(c *gin.Context) {
	ID := c.Param("ID")
	id, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		log.Println(err)
	}
	err1 := h.useCase.DeleteCategoryByID(id)
	if err1 != nil {
		log.Println(err1)
	}
	if err1 == nil {
		c.JSON(http.StatusOK, gin.H{
			constant.JSONSuccess: constant.NotificationDeleteSuccessCategory,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			constant.JSONError: constant.ErrorDeleteCategory,
		})
		return
	}
}

//***********************************************************************************************
func (h *categoryHandler) ListSubCategory(c *gin.Context) {
	var (
		countQuery        int64
		subCategories     []*models.Category
		err               error
		search            SearchCategory
		currentPage       int
		totalPage         int
		currentPageMinusI []int
		currentPageAddJ   []int
		totalPageMinus5   int
		maxPage           int
	)
	Name := c.DefaultQuery("NameSubCategory", "")
	Description := c.DefaultQuery("Description", "")
	CreatedTo := c.DefaultQuery("CreatedTo", "")
	CreatedFrom := c.DefaultQuery("CreatedFrom", "")
	offset := c.DefaultQuery("page", "1")
	NameParent := c.DefaultQuery("selectCategory", "")
	pageNumber, setErr := strconv.Atoi(offset)
	if setErr != nil {
		log.Println("Err", setErr.Error())
	}
	_, countQuery, _ = usecase.NewCategoryUseCase().GetAllSubCategory(Name, Description, CreatedFrom, CreatedTo, NameParent, int64(pageNumber-1)*constant.LimitPages, constant.LimitPages)
	sumPage := int(math.Ceil(float64(countQuery) / float64(constant.LimitPages)))
	if pageNumber >= sumPage {
		maxPage = sumPage
	} else if pageNumber <= 1 {
		maxPage = 1
	} else {
		maxPage = pageNumber
	}
	var selectAll = ""
	if NameParent == "" {
		selectAll = "Select all"
	}
	subCategories, countQuery, err = usecase.NewCategoryUseCase().GetAllSubCategory(Name, Description, CreatedFrom, CreatedTo, NameParent, int64(maxPage-1)*constant.LimitPages, constant.LimitPages)
	if err != nil {
		log.Println(err)
	}

	search.Name = Name
	search.Description = Description
	search.CreatedFrom = CreatedFrom
	search.CreatedTo = CreatedTo
	search.NameParent = NameParent
	currentPage, totalPage,
		currentPageMinusI, currentPageAddJ,
		totalPageMinus5 = utility.CalculatePage(currentPage, maxPage, countQuery, constant.LimitPages)
	pageInfo := &PageCategoryInfo{
		Search:      search,
		TotalPage:   totalPage,
		CurrentPage: currentPage,
		CPMI:        currentPageMinusI,
		CPAJ:        currentPageAddJ,
		TPM5:        totalPageMinus5,
	}
	categories, err1 := usecase.NewCategoryUseCase().GetAllCategoryNoPagination(c)
	if err1 != nil {
		log.Println(err1)
	}
	user, _ := sessions.GetInfoUserSession()
	c.HTML(http.StatusOK, "list_subCategory.html", gin.H{
		"title":        "List sub categories",
		"sub_category": subCategories,
		"category":     categories,
		"user":         user,
		"pageInfo":     pageInfo,
		"selectAll":    selectAll,
	})
}

func (h *categoryHandler) AddSubCategory(c *gin.Context) {
	user, _ := sessions.GetInfoUserSession()
	categories, err1 := usecase.NewCategoryUseCase().GetAllCategoryNoPagination(c)
	if err1 != nil {
		panic(err1)
	}
	c.HTML(http.StatusOK, "create_subCategory.html", gin.H{
		"title":    "Create sub category",
		"category": categories,
		"user":     user,
	})
}

func (h *categoryHandler) ChooseSubCategory(c *gin.Context) {
	id := c.PostForm("id")
	ID, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		log.Println(err)
	} else {
		sub, err := h.useCase.FindSubCategoryByID(ID)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"sub": "",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"sub": sub,
			})
		}
	}
}

func (h *categoryHandler) CreateSubCategory(c *gin.Context) {
	name := c.PostForm("name")
	description := c.PostForm("description")
	CategoryID := c.PostForm("categoryID")
	categoryID, err := strconv.ParseInt(CategoryID, 10, 64)
	if err != nil {
		log.Println(err)
	}
	log.Println("categoryID", categoryID, "name", name, "description", description, "******************")
	category, err2 := usecase.NewCategoryUseCase().FindCategoryByID(categoryID)
	if err2 != nil {
		log.Println(err2)
	}
	menuID := category[0].MenuID
	parentName := category[0].CategoryName
	subCategory := models.Category{MenuID: menuID, CategoryName: name, ParentName: &parentName, Description: description, ParentID: &categoryID}
	err1 := h.useCase.CreateSubCategory(subCategory)
	if err1 != nil {
		log.Println(err1)
	}
	if err1 == nil {
		c.JSON(http.StatusOK, gin.H{
			constant.JSONSuccess: constant.NotificationCreateSuccessSubCategory,
			constant.JSONUrl:     constant.URLSubCategories,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			constant.JSONError: constant.ErrorCreateSubCategory,
		})
		return
	}
}
func (h *categoryHandler) EditSubCategory(c *gin.Context) {
	ID := c.Param("ID")
	id, err := strconv.ParseInt(ID, 10, 64)
	subCategory, err := usecase.NewCategoryUseCase().FindCategoryByID(id)
	if err != nil {
		log.Println(err)
	}
	parentID := subCategory[0].ParentID
	nameParent := subCategory[0].ParentName
	categories, err1 := usecase.NewCategoryUseCase().GetAllCategoryNoPagination(c)
	if err1 != nil {
		log.Println(err1)
	}
	user, _ := sessions.GetInfoUserSession()
	c.HTML(http.StatusOK, "edit_subCategory.html", gin.H{
		"title":       "Edit sub category",
		"categoryID":  id,
		"category":    categories,
		"user":        user,
		"subCategory": subCategory,
		"parentID":    parentID,
		"nameParent":  nameParent,
	})
}

func (h *categoryHandler) UpdateSubCategory(c *gin.Context) {
	ID := c.Param("ID")
	id, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		log.Println(err)
	}
	name := c.PostForm("name")
	description := c.PostForm("description")
	CategoryID := c.PostForm("categoryID")
	categoryID, err := strconv.ParseInt(CategoryID, 0, 64)
	if err != nil {
		log.Println(err)
	}
	category, err1 := usecase.NewCategoryUseCase().FindCategoryByID(categoryID)
	if err1 != nil {
		log.Println(err1)
	}
	menuID := category[0].MenuID
	parentName := category[0].CategoryName
	subcategory := models.Category{MenuID: menuID, CategoryName: name, ParentName: &parentName, Description: description, ParentID: &categoryID}
	err2 := h.useCase.UpdateSubCategory(&subcategory, id)
	if err2 == nil {
		c.JSON(http.StatusOK, gin.H{
			constant.JSONSuccess: constant.NotificationEditSuccessSubCategory,
			constant.JSONUrl:     constant.URLSubCategories,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			constant.JSONError: constant.ErrorEditSubCategory,
		})
		return
	}
}
func (h *categoryHandler) DeleteSubCategory(c *gin.Context) {
	ID := c.Param("ID")
	id, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		log.Println(err)
	}
	err1 := h.useCase.DeleteCategoryByID(id)
	if err1 != nil {
		log.Println(err1)
	}
	if err1 == nil {
		c.JSON(http.StatusOK, gin.H{
			constant.JSONSuccess: constant.NotificationDeleteSuccessSubCategory,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			constant.JSONError: constant.ErrorDeleteSubCategory,
		})
		return
	}
}
