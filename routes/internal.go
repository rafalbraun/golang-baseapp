package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller Controller) Internal(c *gin.Context) {
    pd := controller.DefaultPageData(c)
	c.HTML(http.StatusOK, "internal.html", pd)
}
