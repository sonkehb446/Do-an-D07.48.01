package server

import (
	"Hybrid-blog/constant"

	"github.com/gin-gonic/gin"

	"Hybrid-blog/helpers/view"
)

func Route() error {
	app := setup()
	return app.Run(":" + constant.PortWeb)
}

func setup() *gin.Engine {
	app := gin.Default()
	// sess := utility.ConnectAws()
	// app.Use(func(c *gin.Context) {
	// 	c.Set("sess", sess)
	// 	c.Next()
	// })
	app.Static("/assets", "./assets")
	app.SetFuncMap(view.Funcs())
	app.LoadHTMLGlob("templates/**/*.html")
	router(app)
	return app
}
