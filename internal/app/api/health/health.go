package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-react-app/internal/app/generated"
)

type Controller struct{}

func NewHealth() *Controller {
	return &Controller{}
}

func (h *Controller) Get(c *gin.Context) {
	response := generated.HealthResponse{Status: "OK"}
	c.JSON(http.StatusOK, response)
}
