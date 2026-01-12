package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-gonic/gin"
	"github.com/jerson2000/jquest/config"
	"github.com/jerson2000/jquest/dtos"
	"github.com/jerson2000/jquest/responses"
	"github.com/jerson2000/jquest/services"
	"github.com/jerson2000/jquest/utils"
)

type userController struct {
	service services.UserService
}

func newUserController() *userController {
	userService := services.NewUserService()
	return &userController{service: userService}
}

func (h *userController) registerRoutes(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.GET("", cache.CachePage(config.CacheStore, time.Minute, h.getUsers))
		users.GET("/:id", cache.CachePage(config.CacheStore, time.Minute, h.getUserByID))
		users.POST("", h.createUser)
		users.PUT("/:id", h.updateUser)
		users.DELETE("/:id", h.deleteUser)
	}
}

func (h *userController) getUsers(c *gin.Context) {
	response := h.service.GetUsers()
	c.JSON(utils.ToHTTPStatus(response.Status), response)
}

func (h *userController) getUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	response := h.service.GetUser(id)
	c.JSON(utils.ToHTTPStatus(response.Status), response)
}

func (h *userController) createUser(c *gin.Context) {
	var user dtos.UserCreateRequestDto
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created := h.service.CreateUser(user)
	c.JSON(utils.ToHTTPStatus(created.Status), created)
}

func (h *userController) updateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var user dtos.UserUpdateRequestDto
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated := h.service.UpdateUser(id, user)

	c.JSON(utils.ToHTTPStatus(updated.Status), updated)
}

func (h *userController) deleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Failure[any](http.StatusBadRequest, "invalid id"))
		return
	}

	del := h.service.DeleteUser(id)
	if !del.Success {
		c.JSON(utils.ToHTTPStatus(del.Status), del)
		return
	}

	c.Status(utils.ToHTTPStatus(del.Status))
}
