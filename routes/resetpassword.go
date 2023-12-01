package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"

	"baseapp/models"
	"baseapp/validation"
)

// ResetPasswordPageData defines additional data needed to render the reset password page
type ResetPasswordPageData struct {
	models.PageData
	Token string
}

// ResetPassword renders the HTML page for resetting the users password
func (controller Controller) ResetPassword(c *gin.Context) {
	token := c.Param("token")
	pdPre := controller.DefaultPageData(c)
	pdPre.Title = pdPre.Trans("Reset Password")
	pd := ResetPasswordPageData{
		PageData: pdPre,
		Token:    token,
	}
	c.HTML(http.StatusOK, "resetpassword.html", pd)
}

// ResetPasswordPost handles post request used to reset users passwords
func (controller Controller) ResetPasswordPost(c *gin.Context) {
	pdPre := controller.DefaultPageData(c)
	resetError := pdPre.Trans("Could not reset password, please try again")

	token := c.Param("token")
	pdPre.Title = pdPre.Trans("Reset Password")
	pd := ResetPasswordPageData{
		PageData: pdPre,
		Token:    token,
	}
	password := c.PostForm("password")
	passwordValidation := validation.NewPasswordValidation(password, controller.db, pd.Trans)
	if !passwordValidation.IsValid() {
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: passwordValidation.Error,
		})
		c.HTML(http.StatusBadRequest, "register.html", pd)
		return
	}

	forgotPasswordToken := models.Token{
		Value: token,
		Type:  models.TokenPasswordReset,
	}

	res := controller.db.Where(&forgotPasswordToken).First(&forgotPasswordToken)
	if res.Error != nil {
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: resetError,
		})
		c.HTML(http.StatusBadRequest, "resetpassword.html", pd)
		return
	}

	if forgotPasswordToken.HasExpired() {
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: resetError,
		})
		c.HTML(http.StatusBadRequest, "resetpassword.html", pd)
		return
	}

	user := models.User{}
	user.ID = uint(forgotPasswordToken.ModelID)
	res = controller.db.Where(&user).First(&user)
	if res.Error != nil {
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: resetError,
		})
		c.HTML(http.StatusBadRequest, "resetpassword.html", pd)
		return
	}

	hashedPassword, err := generatePassword(password)

	if err != nil {
		log.Println(err)
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: resetError,
		})
		c.HTML(http.StatusBadRequest, "resetpassword.html", pd)
		return
	}

	user.Password = string(hashedPassword)

	res = controller.db.Save(&user)
	if res.Error != nil {
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: resetError,
		})
		c.HTML(http.StatusBadRequest, "resetpassword.html", pd)
		return
	}

	res = controller.db.Delete(&forgotPasswordToken)
	if res.Error != nil {
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: resetError,
		})
		c.HTML(http.StatusBadRequest, "resetpassword.html", pd)
		return
	}

	pd.Messages = append(pd.Messages, models.Message{
		Type:    "success",
		Content: pdPre.Trans("Your password has successfully been reset."),
	})

	c.HTML(http.StatusOK, "resetpassword.html", pd)
}