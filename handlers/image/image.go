package image

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"Hybrid-blog/constant"
	"Hybrid-blog/helpers/utility"
	"Hybrid-blog/models"
	usecase "Hybrid-blog/usecases"
)

type Images interface {
	UploadFile(c *gin.Context)
	ReturnImages(c *gin.Context, typeImage int, KeyImage string) (*models.Image, error)
}

type image struct {
	usImage usecase.UserCaseImage
}

func NewImages() Images {
	return &image{
		usImage: usecase.NewUserCaseImage(),
	}
}

// func (i *image) UploadFile(c *gin.Context) {
// 	file, header, errFile := c.Request.FormFile("upload")
// 	if errFile != nil {
// 		log.Fatal(errFile, file)
// 	}
// 	ok := utility.CheckTypeImage(header)
// 	if ok {
// 		NameFile := utility.NameRandomImage(header)
// 		s3 := utility.NewS3(file, 1, NameFile)
// 		s3.UploadImage(c)
// 		url := s3.LinkPath
// 		if url != "" {
// 			var image = models.Image{
// 				Url: url,
// 			}
// 			err := i.usImage.CreateImage(&image)
// 			if err == nil {
// 				uri := &image.Url
// 				log.Println(uri)
// 				a := c.Request.FormValue("CKEditorFuncNum")
// 				// https://hylife-dev.s3-ap-southeast-1.amazonaws.com/1639036474.jpg
// 				link := fmt.Sprintf("<script>window.parent.CKEDITOR.tools.callFunction(%s,\""+*uri+"\");</script>", a)
// 				log.Println("Link ->", link)
// 				fmt.Fprintln(c.Writer, link)
// 			}
// 		}

// 	}
// }

func (i *image) UploadFile(c *gin.Context) {
	file, err := c.FormFile("upload")
	if err != nil {
		log.Println("Error ->", err)
	}
	ok := utility.CheckTypeImage(file)
	if ok {
		uri, err := utility.SaveFile(file, constant.ImageTypePost)
		if err != nil {
			log.Println("Error ->", err)
			return
		}

		if uri != "" {
			var image = models.Image{
				Url: uri,
			}
			err := i.usImage.CreateImage(&image)
			if err == nil {
				uri := &image.Url
				log.Println(uri)
				a := c.Request.FormValue("CKEditorFuncNum")
				link := fmt.Sprintf("<script>window.parent.CKEDITOR.tools.callFunction(%s,\""+*uri+"\");</script>", a)
				log.Println("Link ->", link)
				fmt.Fprintln(c.Writer, link)
			}
		}
	}
}

func (i *image) ReturnImages(c *gin.Context, typeImage int, KeyImage string) (*models.Image, error) {
	file, err := c.FormFile(KeyImage)
	if err != nil {
		log.Println("Error ->", err)
	}
	ok := utility.CheckTypeImage(file)
	if ok {
		uri, err := utility.SaveFile(file, typeImage)
		if err != nil {
			log.Println("Error ->", err)
			return nil, err
		}

		if uri != "" {
			var image = models.Image{
				Url: uri,
			}
			err := i.usImage.CreateImage(&image)
			if err == nil {
				return &image, nil
			}
		}
	}
	return nil, nil
}
