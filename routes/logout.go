package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"baseapp/middleware"
	"baseapp/lang"
)

func (controller Controller) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(middleware.SessionIDKey)
	session.Delete(lang.LangID)
	err := session.Save()
	if err != nil {
		log.Println(err)
	}
	c.Redirect(http.StatusTemporaryRedirect, "/index")
}
