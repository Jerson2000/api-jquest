package controllers

import (
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/jerson2000/jquest/config"
	"github.com/jerson2000/jquest/enums"
	"github.com/jerson2000/jquest/middlewares"
	"github.com/jerson2000/jquest/repositories"
	"github.com/jerson2000/jquest/responses"
	"github.com/jerson2000/jquest/services"
)

var trans ut.Translator

func InitController(router *gin.Engine) {
	initValidator()

	// Repositories
	userRepo := repositories.NewUserRepository(config.Database)
	recruiterRepo := repositories.NewRecruiterRepository(config.Database)
	companyRepo := repositories.NewCompanyRepository(config.Database)
	jobRepo := repositories.NewJobRepository(config.Database)
	candidateRepo := repositories.NewCandidateRepository(config.Database)
	applicationRepo := repositories.NewApplicationRepository(config.Database)
	savedJobRepo := repositories.NewSavedJobRepository(config.Database)
	skillRepo := repositories.NewSkillRepository(config.Database)

	// Services
	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)
	recruiterService := services.NewRecruiterService(recruiterRepo, companyRepo, userRepo)
	companyService := services.NewCompanyService(companyRepo, recruiterRepo, userRepo)
	jobService := services.NewJobService(jobRepo, recruiterRepo)
	applicationService := services.NewApplicationService(applicationRepo, jobRepo, candidateRepo, recruiterRepo, userRepo)
	savedJobService := services.NewSavedJobService(savedJobRepo, candidateRepo, jobRepo)
	skillService := services.NewSkillService(skillRepo, candidateRepo, jobRepo, recruiterRepo)

	rateLimit := middlewares.NewRateLimiter("50-M")
	router.Use(rateLimit.Middleware())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     config.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(middlewares.CSRFMiddleware(
		config.CSRFKey,
		config.AppEnv == "production",
	))
	router.Use(func(c *gin.Context) {
		c.Header("Content-Security-Policy", "default-src 'self';")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Next()
	})


	router.GET("/api/token", func(c *gin.Context) {
		xtoken, exists := c.Get("csrf_token")
		var tokenStr string
		if exists && xtoken != nil {
			if s, ok := xtoken.(string); ok {
				tokenStr = s
			}
		}
		if tokenStr != "" {
			c.Header("X-CSRF-Token", tokenStr)
		}
		res := responses.Success[any](http.StatusOK, gin.H{"token": tokenStr})
		c.JSON(http.StatusOK, res)
	})

	public := router.Group("/api")
	newAuthController(authService).registerRoutes(public)

	authorize := router.Group("/api")
	authorize.Use(middlewares.JwtMiddleware())
	authorize.Use(middlewares.RBACMiddleware())
	newUserController(userService).registerRoutes(authorize)
	newRecruiterController(recruiterService).registerRoutes(authorize)

	newCompanyController(companyService).registerRoutes(public, authorize)
	newJobController(jobService).registerRoutes(public, authorize)
	newApplicationController(applicationService).registerRoutes(public, authorize)
	newSavedJobController(savedJobService).registerRoutes(authorize)
	newSkillController(skillService).registerRoutes(public, authorize)

	authorize.GET("/current", func(c *gin.Context) {
		id := c.MustGet("id").(int)
		name := c.MustGet("name").(string)
		role := c.MustGet("role").(enums.Role)
		c.JSON(http.StatusOK, gin.H{"id": id, "name": name, "role": role})
	})
}

func initValidator() {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		enLocale := en.New()
		uni := ut.New(enLocale, enLocale)
		trans, _ = uni.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(v, trans)

		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "" || name == "-" {
				return fld.Name
			}
			return name
		})
	}
}
