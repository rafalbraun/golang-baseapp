package baseproject

import (
	"baseapp/config"
	"baseapp/middleware"
	"baseapp/models"
	"baseapp/routes"
	"baseapp/system"
	"baseapp/funcmaps"

	"html/template"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func SetupRouter() (*gin.Engine, *config.Config) {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// We load environment variables, these are only read when the application launches
	config := config.LoadEnv()

	// Translations
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	languages := []string{
		"en",
		//"pl",
	}
	for _, l := range languages {
		_, err := bundle.LoadMessageFile(fmt.Sprintf("active.%s.toml", l))
		if err != nil {
			log.Fatalln(err)
		}
	}

	// We create a new cookie store with a key used to secure cookies with HMAC
	store := cookie.NewStore([]byte(config.CookieSecret))

	// We define our session middleware to be used globally on all routes
	router.Use(sessions.Sessions("golang_base_project_session", store))

	// Must be here, before loading templates
	router.SetFuncMap(template.FuncMap(funcmaps.FuncMap))

	// Load templates
	files, err := wrapLoadTemplates("templates")
	if err != nil {
		log.Println(err)
	}
	router.LoadHTMLFiles(files...)

	// We connect to the database using the configuration generated from the environment variables.
	db, err := ConnectToDatabaseSQLite(config, os.Stdout /*, gormLogWriter*/)
	if err != nil {
		log.Fatalf("err connecting to database: %v", err)
	}

    // Create log file for gin output
//     logFile, err := openLogFile("gin.log")
//     if err != nil {
//         log.Fatal(err)
//     }
//     log.SetOutput(logFile)
//     log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
// 	router.Use(gin.LoggerWithWriter(logFile))

	// Set a lower memory limit for multipart forms (default is 32 MiB), 8 << 20 == 8 MiB
	router.MaxMultipartMemory = 8 << 20

	// Hosting static css and js files
	router.Static("/static", "./static")

	// Session middleware is applied to all groups after this point.
	router.Use(middleware.Session(db))

	loggerInterface, err := models.OpenAuditLogger()
	if err != nil {
		log.Fatalf("err creating audit logger: %v", err)
	}

	// Initialize system roles by creating admin, moderator and user roles
	system.InitSystemRoles(db)

	// Connection extends gorm.DB adding new features like paging
	connection := models.New(db, loggerInterface)

	// A new instance of the routes controller is created
	controller := routes.New(connection, config, bundle)

	// We define our 404 handler for when a page can not be found
	router.NoRoute(controller.NoRoute)
	router.GET("/settings", controller.NotImplemented)
	router.GET("/profile/:username", controller.NotImplemented)

	// noAuth is a group for routes which should only be accessed if the user is not authenticated
	noAuth := router.Group("/")
	noAuth.Use(middleware.NoAuth())
	router.GET("/", controller.Index)
	router.GET("/index", controller.Index)
	noAuth.GET("/login", controller.Login)
	noAuth.GET("/register", controller.Register)
	noAuth.GET("/activate/resend", controller.ResendActivation)
	noAuth.GET("/activate/:token", controller.Activate)
	noAuth.GET("/user/password/forgot", controller.ForgotPassword)
	noAuth.GET("/user/password/reset/:token", controller.ResetPassword)

	// we make a separate group for our post requests on the same endpoints so that we can define our throttling middleware on POST requests only.
	noAuthPost := noAuth.Group("/")
	noAuthPost.Use(middleware.Throttle(config.RequestsPerMinute))
	noAuthPost.POST("/login", controller.LoginPost)
	noAuthPost.POST("/register", controller.RegisterPost)
	noAuthPost.POST("/activate/resend", controller.ResendActivationPost)
	noAuthPost.POST("/user/password/forgot", controller.ForgotPasswordPost)
	noAuthPost.POST("/user/password/reset/:token", controller.ResetPasswordPost)

	// the admin group handles routes that should only be accessible to authenticated users
	auth := router.Group("/")
	auth.Use(middleware.Auth())
	auth.GET("/users", controller.PagedUsers)
	auth.GET("/logout", controller.Logout)
	auth.GET("/internal", controller.Internal)

    authPost := router.Group("/")
    authPost.Use(middleware.Throttle(config.RequestsPerMinute))
    authPost.POST("/internal", controller.Internal)

	return router, config
}
