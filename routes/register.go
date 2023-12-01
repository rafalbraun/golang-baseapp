package routes

import (
	"baseapp/emails"
	"baseapp/models"
	"baseapp/system"
	"baseapp/validation"

	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Register renders the HTML content of the register page
func (controller Controller) Register(c *gin.Context) {
	pd := controller.DefaultPageData(c)
	pd.Title = "Register"
	c.HTML(http.StatusOK, "register.html", pd)
}

// RegisterPost handles requests to register users and returns appropriate messages as HTML content
func (controller Controller) RegisterPost(c *gin.Context) {
	pd := controller.DefaultPageData(c)
	registerError := pd.Trans("Could not register, please make sure the details you have provided are correct and that you do not already have an existing account.")
	registerSuccess := pd.Trans("Thank you for registering. An activation email has been sent with steps describing how to activate your account.")
	pd.Title = pd.Trans("Register")

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

	hashedPassword, err := generatePassword(password)
	if err != nil {
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: registerError,
		})
		c.HTML(http.StatusBadRequest, "register.html", pd)
		return
	}

	email := c.PostForm("email")
	emailValidation := validation.NewEmailValidation(email, controller.db, pd.Trans)
	if !emailValidation.IsValid() {
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: emailValidation.Error,
		})
		c.HTML(http.StatusBadRequest, "register.html", pd)
		return
	}

	username := c.PostForm("username")
	usernameValidation := validation.NewUsernameValidation(username, controller.db, pd.Trans)
	if !usernameValidation.IsValid() {
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: usernameValidation.Error,
		})
		c.HTML(http.StatusBadRequest, "register.html", pd)
		return
	}

    // if username of newly registered account is "Administrator" we give it such role
	role := system.SystemUser
	if username == system.SystemAdmin.RoleName {
		role = system.SystemAdmin
	}

    // if username of newly registered account is other than "Administrator" we set activation date as nil
	activatedAt := toTimePtr(time.Now())
    if username != system.SystemAdmin.RoleName {
        activatedAt = nil
    }

	user := models.User{
		Email:    email,
		Username: username,
		Password:   string(hashedPassword),
		UserRoleId: role.RoleId,
		ActivatedAt: activatedAt,
	}

	res := controller.db.WriteAuditLog().Save(&user)
	if res.Error != nil || res.RowsAffected == 0 {
		pd.Messages = append(pd.Messages, models.Message{
			Type:    "error",
			Content: registerError,
		})
		log.Println(res.Error)
		c.HTML(http.StatusInternalServerError, "register.html", pd)
		return
	}

	// Generate activation token and send activation email
	go controller.activationEmailHandler(user.ID, email, pd.Trans)

	pd.Messages = append(pd.Messages, models.Message{
		Type:    "success",
		Content: registerSuccess,
	})

	c.HTML(http.StatusOK, "register.html", pd)
}

func (controller Controller) activationEmailHandler(userID uint, email string, trans func(string) string) {
	activationToken := models.Token{
		Value: generateToken(),
		Type:  models.TokenUserActivation,
	}

	res := controller.db.Where(&activationToken).First(&activationToken)
	if (res.Error != nil && res.Error != gorm.ErrRecordNotFound) || res.RowsAffected > 0 {
		// If the activation token already exists we try to generate it again
		controller.activationEmailHandler(userID, email, trans)
		return
	}

	activationToken.ModelID = int(userID)
	activationToken.ModelType = "User"
	activationToken.ExpiresAt = time.Now().Add(time.Minute * 10)

	res = controller.db.Save(&activationToken)
	if res.Error != nil || res.RowsAffected == 0 {
		log.Println(res.Error)
		return
	}
	controller.sendActivationEmail(activationToken.Value, email, trans)
}

func (controller Controller) sendActivationEmail(token string, email string, trans func(string) string) {
	u, err := url.Parse(controller.conf.BaseURL)
	if err != nil {
		log.Println(err)
		return
	}

	u.Path = path.Join(u.Path, "/activate/", token)

	activationURL := u.String()

	emailService := emails.New(controller.conf)

	err = emailService.Send(email, trans("User Activation"), fmt.Sprintf(trans("Use the following link to activate your account. If this was not requested by you, please ignore this email.\n%s"), activationURL))

    if err != nil {
        log.Printf("Failed sending registration email: [%v] \n", err.Error())
    }
}