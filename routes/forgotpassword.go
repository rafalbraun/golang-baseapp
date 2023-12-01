package routes

import(
    "gorm.io/gorm"
    "net/http"
    "net/url"
    "github.com/gin-gonic/gin"
    "log"
    "time"
    "fmt"
	"path"

   	"baseapp/emails"
	"baseapp/models"
)

func (controller Controller) ForgotPassword(c *gin.Context) {
    pd := controller.DefaultPageData(c)
    pd.Title = "Forgot Password"
    c.HTML(http.StatusOK, "forgotpassword.html", pd)
}

func (controller Controller) ForgotPasswordPost(c *gin.Context) {
	pd := controller.DefaultPageData(c)
	pd.Title = "Forgot Password"

	email := c.PostForm("email")
	user := models.User{Email: email}
	res := controller.db.Where(&user).First(&user)
	if res.Error == nil && user.ActivatedAt != nil {
		go controller.forgotPasswordEmailHandler(user.ID, email)
	}

	pd.Messages = append(pd.Messages, models.Message{
		Type:    "success",
		Content: "An email with instructions describing how to reset your password has been sent.",
	})

	// We always return a positive response here to prevent user enumeration
	c.HTML(http.StatusOK, "forgotpassword.html", pd)
}

func (controller Controller) forgotPasswordEmailHandler(userID uint, email string) {
	forgotPasswordToken := models.Token{
		Value: generateToken(),
		Type:  models.TokenPasswordReset,
	}

	res := controller.db.Where(&forgotPasswordToken).First(&forgotPasswordToken)
	if (res.Error != nil && res.Error != gorm.ErrRecordNotFound) || res.RowsAffected > 0 {
		// If the forgot password token already exists we try to generate it again
		controller.forgotPasswordEmailHandler(userID, email)
		return
	}

	forgotPasswordToken.ModelID = int(userID)
	forgotPasswordToken.ModelType = "User"

	// The token will expire 10 minutes after it was created
	forgotPasswordToken.ExpiresAt = time.Now().Add(time.Minute * 10)

	res = controller.db.Save(&forgotPasswordToken)
	if res.Error != nil || res.RowsAffected == 0 {
		log.Println(res.Error)
		return
	}
	controller.sendForgotPasswordEmail(forgotPasswordToken.Value, email)
}

func (controller Controller) sendForgotPasswordEmail(token string, email string) {
	u, err := url.Parse(controller.conf.BaseURL)
	if err != nil {
		log.Println(err)
		return
	}

	u.Path = path.Join(u.Path, "/user/password/reset/", token)

	resetPasswordURL := u.String()

	emailService := emails.New(controller.conf)

	err = emailService.Send(email, "Password Reset",
	    fmt.Sprintf("Use the following link to reset your password. If this was not requested by you, please ignore this email.\n%s", resetPasswordURL))

    if err != nil {
        log.Printf("Failed sending forgotten password email: [%v] \n", err.Error())
    }
}
