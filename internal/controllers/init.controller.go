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
	"github.com/jerson2000/jquest/internal/config"
	"github.com/jerson2000/jquest/internal/enums"
	"github.com/jerson2000/jquest/internal/middlewares"
)

var trans ut.Translator

func InitController(router *gin.Engine) {
	config.InitConfig()
	initValidator()

	router.Use(cors.Default())

	public := router.Group("/api")
	newAuthController().registerRoutes(public)

	authorize := router.Group("/api")
	authorize.Use(middlewares.JwtMiddleware())
	newUserController().registerRoutes(authorize)

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
