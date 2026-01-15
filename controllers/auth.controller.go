package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jerson2000/jquest/dtos"
	"github.com/jerson2000/jquest/middlewares"
	"github.com/jerson2000/jquest/services"
	"github.com/jerson2000/jquest/utils"
)

type authController struct {
	service services.AuthService
}

func newAuthController() *authController {
	service := services.NewAuthService()
	return &authController{service: service}
}

func (auth *authController) registerRoutes(r *gin.RouterGroup) {
	routes := r.Group("/auth")
	{
		routes.POST("login", auth.login)
		routes.POST("signup", auth.signup)
		routes.POST("refresh", middlewares.JwtMiddleware(), auth.refresh)
	}
}

func (auth *authController) login(c *gin.Context) {
	var dto dtos.AuthLoginRequestDto
	if !utils.ValidationhouldBind(http.StatusBadRequest, &dto, trans, c) {
		return
	}

	res := auth.service.Login(dto)
	c.JSON(utils.ToHTTPStatus(res.Status), res)
}

func (auth *authController) signup(c *gin.Context) {
	var dto dtos.AuthSignupRequestDto

	if !utils.ValidationhouldBind(http.StatusBadRequest, &dto, trans, c) {
		return
	}

	res := auth.service.Signup(dto)
	c.JSON(utils.ToHTTPStatus(res.Status), res)
}

func (auth *authController) refresh(c *gin.Context) {
	var dto dtos.AuthRefreshRequestDto
	if !utils.ValidationhouldBind(http.StatusBadRequest, &dto, trans, c) {
		return
	}

	res := auth.service.Refresh(dto)
	c.JSON(utils.ToHTTPStatus(res.Status), res)
}
