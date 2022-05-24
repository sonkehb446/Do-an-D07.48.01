package Home

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"Hybrid-blog/sessions"
	usecase "Hybrid-blog/usecases"
)

type Home interface {
	TemplateHome(c *gin.Context)
}

type handlerHome struct {
	useCase usecase.UserCasePost
}

func NewHandlerHome() Home {
	return &handlerHome{
		useCase: usecase.NewPostUseCase(),
	}
}

func (h *handlerHome) TemplateHome(c *gin.Context) {
	user, err := sessions.GetInfoUserSession()
	news, _, _ := h.useCase.GetPostByType(1, 0, 3)
	birders, _, _ := h.useCase.GetPostByType(2, 0, 3)
	learning, _, _ := h.useCase.GetPostByType(3, 0, 2)
	if err != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":    "Home",
			"news":     news,
			"birders":  birders,
			"learning": learning,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":    "Home",
			"user":     user,
			"news":     news,
			"birders":  birders,
			"learning": learning,
		})
	}
}
