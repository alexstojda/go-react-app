package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-react-app/internal/app/generated"
)

type Controller struct {
}

func NewHello() *Controller {
	return &Controller{}
}

func (h *Controller) Get(c *gin.Context) {
	response := generated.HelloResponse{Message: "Hello, World!"}
	c.JSON(http.StatusOK, response)
}
