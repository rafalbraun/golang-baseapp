package routes

import(
	"github.com/gin-gonic/gin"
	"baseapp/models"
	"log"
	"net/http"
)

func (controller Controller) PagedUsers(c *gin.Context) {
	pd := controller.DefaultPageData(c)
    pd.Title = "Users"

    res := controller.db.Model(&models.User{}).
            Preload("Role").
            Where(func(search string) string {if search=="" {return "1=1"} else {return "username like \"%"+search+"%\" "}}(pd.Search)).
            Paged(&pd.Pagination, &pd.Users)
    if res.Error != nil {
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: "Could not list paged users",
		})
		log.Println(res.Error)
        c.HTML(http.StatusInternalServerError, "pagedUsers.html", pd)
        return
    }

	c.HTML(http.StatusOK, "pagedUsers.html", pd)
}
