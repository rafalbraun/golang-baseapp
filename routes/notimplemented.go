package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller Controller) NotImplemented(c *gin.Context) {
	pd := controller.DefaultPageData(c)
	c.HTML(http.StatusOK, "notimplemented.html", pd)
}
