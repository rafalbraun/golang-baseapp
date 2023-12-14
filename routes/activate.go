package routes

import(
	"github.com/gin-gonic/gin"
	"baseapp/models"
	"log"
	"net/http"
	"time"
)

// Activate handles requests used to activate a users account
func (controller Controller) Activate(c *gin.Context) {
	pd := controller.DefaultPageData(c)
	activationError := pd.Trans("Please provide a valid activation token")
	activationSuccess := pd.Trans("Account activated. You may now proceed to login to your account.")
	pd.Title = pd.Trans("Activate")
	token := c.Param("token")
	activationToken := models.Token{
		Value: token,
		Type:  models.TokenUserActivation,
	}

	res := controller.db.Where(&activationToken).First(&activationToken)
	if res.Error != nil {
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: activationError,
		})
		c.HTML(http.StatusBadRequest, "activate.html", pd)
		return
	}

	if activationToken.HasExpired() {
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: activationError,
		})
		c.HTML(http.StatusBadRequest, "activate.html", pd)
		return
	}

	user := models.User{}
	user.ID = uint(activationToken.ModelID)

	res = controller.db.Where(&user).First(&user)
	if res.Error != nil {
		log.Println(res.Error)
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: activationError,
		})
		c.HTML(http.StatusBadRequest, "activate.html", pd)
		return
	}

	now := time.Now()
	user.ActivatedAt = &now

	res = controller.db.Save(&user)
	if res.Error != nil {
		log.Println(res.Error)
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: activationError,
		})
		c.HTML(http.StatusBadRequest, "activate.html", pd)
		return
	}

	// We don't need to check for an error here, even if it's not deleted it will not really affect application logic
	controller.db.Delete(&activationToken)

	pd.Messages = append(pd.Messages, models.Message{
		Type:    "success",
		Content: activationSuccess,
	})
	c.HTML(http.StatusOK, "activate.html", pd)
}