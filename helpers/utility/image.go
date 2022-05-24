package utility

import (
	"Hybrid-blog/constant"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/google/uuid"
)

var magicTable = map[string]string{
	"\xff\xd8\xff":      "image/jpeg",
	"\x89PNG\r\n\x1a\n": "image/png",
	"GIF87a":            "image/gif",
	"GIF89a":            "image/gif",
}

var TypeImage = map[int]string{
	constant.ImageTypeProfile:    "profile",
	constant.ImageTypePost:       "post",
	constant.ImageTypeSmallImage: "smallImage",
}

func CheckTypeImage(file *multipart.FileHeader) bool {
	filetype := file.Filename
	check := false
	if validateType(filetype) {
		check = true
	}
	return check
}

func validateType(url string) bool {
	regex := "([^\\s]+(\\.(?i)(jpe?g|png))$)"
	var check bool
	re := regexp.MustCompile(regex)
	check = re.MatchString(url)
	return check
}

func SaveFile(fh *multipart.FileHeader, Type int) (string, error) {
	fileExt := filepath.Ext(fh.Filename)
	newName := fmt.Sprint(time.Now().Unix()) + fileExt
	var fileNew *os.File
	url := ""
	// if Type == 1 {
	// 	url = "/assets/img/upload/" + newName
	// 	uri := "./assets/img/upload/"
	// 	fileNew, _ = os.Create(filepath.Join(uri, filepath.Base(newName)))
	// } else if Type == 2 {
	// 	url = "/assets/img/post/" + newName
	// 	uri := "./assets/img/post/"
	// 	fileNew, _ = os.Create(filepath.Join(uri, filepath.Base(newName)))
	// }
	var folder = TypeImage[Type]
	url = "/assets/img/" + folder + "/" + newName
	uri := "./assets/img/" + folder + "/"
	fileNew, _ = os.Create(filepath.Join(uri, filepath.Base(newName)))

	if fileNew != nil {
		src, errOpen := fh.Open()
		if errOpen != nil {
			return "", errOpen
		}
		defer src.Close()
		if _, errCopy := io.Copy(fileNew, src); errCopy != nil {
			return "", errCopy
		}
	}
	defer fileNew.Close()
	return url, nil
}

func NameRandomImage(fh *multipart.FileHeader) string {
	uuid, _ := uuid.NewUUID()
	fileExt := filepath.Ext(fh.Filename)
	newName := fmt.Sprint(uuid) + fileExt
	return newName
}
