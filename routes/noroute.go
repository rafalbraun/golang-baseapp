package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller Controller) NoRoute(c *gin.Context) {
	pd := controller.DefaultPageData(c)
	pd.Title = "No route"
	c.HTML(http.StatusOK, "noroute.html", pd)
}
