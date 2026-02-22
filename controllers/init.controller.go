package controllers

import (
	"net/http"
	"reflect"
	"strings"

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
	"github.com/jerson2000/jquest/responses"
)

var trans ut.Translator

func InitController(router *gin.Engine) {
	initValidator()

	rateLimit := middlewares.NewRateLimiter("50-M")
	router.Use(rateLimit.Middleware())
	router.Use(cors.Default())
	router.Use(middlewares.CSRFMiddleware(
		config.CSRFKey,
		true,
	))
	router.Use(func(c *gin.Context) {
		c.Header("Content-Security-Policy", "default-src 'self';")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Next()
	})

	router.GET("/api/token", func(c *gin.Context) {
		xtoken, _ := c.Get("csrf_token")
		c.Header("X-CSRF-Token", xtoken.(string))
		res := responses.Success[any](http.StatusOK, gin.H{"token": xtoken})
		c.JSON(http.StatusOK, res)
	})

	public := router.Group("/api")
	newAuthController().registerRoutes(public)

	authorize := router.Group("/api")
	authorize.Use(middlewares.JwtMiddleware())
	authorize.Use(middlewares.RBACMiddleware())
	newUserController().registerRoutes(authorize)
	newRecruiterController().registerRoutes(authorize)

	newCompanyController().registerRoutes(public, authorize)
	newJobController().registerRoutes(public, authorize)
	newApplicationController().registerRoutes(public, authorize)

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
