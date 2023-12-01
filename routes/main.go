package routes

import (
	"baseapp/config"
	"baseapp/lang"
	"baseapp/middleware"
	"baseapp/models"
	"baseapp/system"
	"baseapp/text"

    "time"
    "strconv"

	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// Controller holds all the variables needed for routes to perform their logic
type Controller struct {
	db      *models.Conn
	conf    *config.Config
	bundle  *i18n.Bundle
}

// New creates a new instance of the routes.Controller
func New(db *models.Conn, conf *config.Config, bundle *i18n.Bundle) Controller {
	return Controller{
		db:     db,
		conf:   conf,
		bundle: bundle,
	}
}

// isAuthenticated checks if the current user is authenticated or not
func isAuthenticated(c *gin.Context) bool {
	_, exists := c.Get(middleware.UserIDKey)
	return exists
}

func (controller Controller) DefaultPageData(c *gin.Context) models.PageData {
	var user models.User
	var loggedInUser *models.User = nil

	userId, exists := c.Get(middleware.UserIDKey)
	if exists {
		controller.db.Where("id=?", userId).Preload("Role").First(&user)
		loggedInUser = &user
	}

	langService, language := lang.New(c, controller.bundle)

    search := c.Query("search")
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("records", "10"))
	pagination := models.Pagination{
	    DefaultLimit:   controller.conf.DefaultPageSize,
	    Limit:          limit,
	    Page:           page,
	    PreviousPage:   page-1,
	    NextPage:       page+1,
	}

	return models.PageData{
		IsAuthenticated:    isAuthenticated(c),
		IsAdmin:            system.IsAdmin(&user),
		Messages:           nil,
		LoggedIn:           loggedInUser,
		Pagination:         pagination,
		Search:             search,
		Trans:              langService.Trans,
		Language:           language,
		BaseURL:            controller.conf.BaseURL,
	}
}

func generateSessionId() string {
	charset := append(text.LettersUppercase(), text.LettersLowercase()...)
	return text.RandomString(charset, 20)
}

func generateToken() string {
    charset := append(text.LettersUppercase(), text.LettersLowercase()...)
    return text.RandomString(charset, 20)
}

func generatePassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hashedPassword), err
}

func toTimePtr(t time.Time) *time.Time {
    return &t
}
