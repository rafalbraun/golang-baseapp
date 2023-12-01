package routes

import (
	"net/http"
	"time"

	"baseapp/middleware"
	"baseapp/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (controller Controller) Login(c *gin.Context) {
	pd := controller.DefaultPageData(c)
	pd.Title = "Login"
	c.HTML(http.StatusOK, "login.html", pd)
}

func (controller Controller) LoginPost(c *gin.Context) {
	pd := controller.DefaultPageData(c)
	loginError := pd.Trans("Could not login, please make sure that you have typed in the correct email and password. If you have forgotten your password, please click the forgot password link below.")
	pd.Title = pd.Trans("Login")
	username := c.PostForm("username")

	user := models.User{}
	res := controller.db.Where("username=?", username).First(&user)
	if res.Error != nil {
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: loginError + "1",
		})
		c.HTML(http.StatusInternalServerError, "login.html", pd)
		return
	}

	if res.RowsAffected == 0 {
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: loginError + "2",
		})
		c.HTML(http.StatusBadRequest, "login.html", pd)
		return
	}

	if user.ActivatedAt == nil {
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: pd.Trans("Account is not activated yet."),
		})
		c.HTML(http.StatusBadRequest, "login.html", pd)
		return
	}

	password := c.PostForm("password")
	err := user.CheckPassword(password)
	if err != nil {
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: loginError + "3",
		})
		c.HTML(http.StatusBadRequest, "login.html", pd)
		return
	}

	// Generate a ulid for the current session
	sessionIdentifier := generateSessionId()
	ses := models.Session{
		Identifier: sessionIdentifier,
	}

	remember := c.PostForm("remember")
	if remember == "on" {
		// Session is valid forever
		ses.ExpiresAt = time.Date(2070, time.January, 1, 0, 0, 0, 0, time.UTC)
	} else {
		// Session is valid for 1 hour
		ses.ExpiresAt = time.Now().Add(time.Hour)
	}
	ses.UserID = user.ID

	ses.UserAgent = c.Request.UserAgent()

	res = controller.db.Save(&ses)
	if res.Error != nil {
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: loginError + "4",
		})
		c.HTML(http.StatusInternalServerError, "login.html", pd)
		return
	}

	session := sessions.Default(c)
	session.Set(middleware.SessionIDKey, sessionIdentifier)

	err = session.Save()
	if err != nil {
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: loginError + "5",
		})
		c.HTML(http.StatusInternalServerError, "login.html", pd)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "/internal")
}
