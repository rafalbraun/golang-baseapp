package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller Controller) Index(c *gin.Context) {
	pd := controller.DefaultPageData(c)
	c.HTML(http.StatusOK, "index.html", pd)
}
