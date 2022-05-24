package server

import (
	"Hybrid-blog/handlers/managerCategory"
	"github.com/gin-gonic/gin"

	"Hybrid-blog/handlers/authen"
	home "Hybrid-blog/handlers/home"
	img "Hybrid-blog/handlers/image"
	managerPost "Hybrid-blog/handlers/managerPost"
	"Hybrid-blog/handlers/managerUser"
	"Hybrid-blog/handlers/middleware"
	"Hybrid-blog/handlers/user"
)

var (
	loginLogout    authen.Authenticator
	ForgotPassword user.ForgotPwInterface
	changePassword user.ChangePwInterface
)

var (
	index      home.Home
	userDetail user.UserDetail
	mgrUser    managerUser.ManagerUser
)

var (
	mgrPost managerPost.ManagerPost
)

var Img img.Images

func router(app *gin.Engine) {
	newHandler()
	app.GET("/login", loginLogout.TemplateLogin)
	app.POST("/login", loginLogout.Login)
	app.GET("/logout", loginLogout.Logout)

	app.GET("/resetpass", ForgotPassword.TemplateResetPass)
	app.POST("/resetpass", ForgotPassword.ResetPassword)
	app.POST("/forgotpass", ForgotPassword.ForgotPassword)
	app.GET("/home", index.TemplateHome)
	app.GET("/", index.TemplateHome)

	categoryHandler := managerCategory.NewCategoryHandler()
	category := app.Group("/category")
	category.Use(middleware.Middleware(), middleware.MiddleAdmin())
	{
		category.GET("/s", categoryHandler.ListCategory)
		category.GET("/add", categoryHandler.AddCategory)
		category.POST("/create", categoryHandler.CreateCategory)
		category.GET("/edit/:ID", categoryHandler.EditCategory)
		category.POST("/update/:ID", categoryHandler.UpdateCategory)
		category.GET("/delete/:ID", categoryHandler.DeleteCategory)
	}
	subCategory := app.Group("/subCategory")
	subCategory.Use(middleware.Middleware())
	{
		subCategory.GET("/s", middleware.MiddleAdmin(), categoryHandler.ListSubCategory)
		subCategory.GET("/add", middleware.MiddleAdmin(), categoryHandler.AddSubCategory)
		subCategory.POST("/create", middleware.MiddleAdmin(), categoryHandler.CreateSubCategory)
		subCategory.GET("/edit/:ID", middleware.MiddleAdmin(), categoryHandler.EditSubCategory)
		subCategory.POST("/update/:ID", middleware.MiddleAdmin(), categoryHandler.UpdateSubCategory)
		subCategory.GET("/delete/:ID", middleware.MiddleAdmin(), categoryHandler.DeleteSubCategory)
		subCategory.POST("/choose", categoryHandler.ChooseSubCategory)
	}

	user := app.Group("/user/")
	user.Use(middleware.Middleware())
	{
		user.GET("/changepass", changePassword.TemplateChangePassword)
		user.POST("/changepass", changePassword.ChangePassword)
		user.GET("detail", userDetail.TemplateUserDetail)
		user.POST("edit", userDetail.EditUserDetail)
		user.GET("create", middleware.MiddleAdmin(), mgrUser.TemplateCreateUser)
		user.POST("create", middleware.MiddleAdmin(), mgrUser.CreateUser)
		user.POST("delete/:id", middleware.MiddleAdmin(), mgrUser.DeleteUser)
		user.POST("edit/:id", middleware.MiddleAdmin(), mgrUser.EditUser)
		user.GET("s/", middleware.MiddleAdmin(), mgrUser.TemplateUsers)
	}

	post := app.Group("/post/")

	hybridPage := post.Group("hybrid")
	{
		hybridPage.GET("/category", mgrPost.TemplatePostsByCategory)
		hybridPage.GET("/subCategory", mgrPost.TemplatePostsBySubCategory)
		hybridPage.GET("/hynews", mgrPost.TemplateHyBirdNew)
		hybridPage.GET("/hybriders", mgrPost.TemplateHyBirders)
		hybridPage.GET("/hylearning", mgrPost.TemplateHyLearning)
		hybridPage.GET("/:id", mgrPost.DetailPost)
	}

	{
		post.Use(middleware.Middleware())
		post.GET("create", mgrPost.TemplateCreatePost)
		post.POST("create", mgrPost.CreatePost)
		post.GET("edit/:id", mgrPost.TemplateEditPost)
		post.POST("edit/:id", mgrPost.EditPost)
		post.POST("delete/:id", mgrPost.DeletePost)
		post.GET("s/", mgrPost.TemplatePosts)
		post.POST("s/", mgrPost.TemplatePosts)
	}

	image := app.Group("/image/")
	image.Use(middleware.Middleware())
	{
		image.POST("/uploader/upload", Img.UploadFile)
	}

}

func newHandler() {
	loginLogout = authen.NewHandlerAuthenticator()
	ForgotPassword = user.NewHandlerForgotPassword()
	changePassword = user.NewHandlerChangePass()
	index = home.NewHandlerHome()
	userDetail = user.NewHandlerUserDetail()
	mgrUser = managerUser.NewHandlerManagerUser()
	mgrPost = managerPost.NewHandlerManagerPost()
	Img = img.NewImages()
}
