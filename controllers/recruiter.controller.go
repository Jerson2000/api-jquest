package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jerson2000/jquest/dtos"
	"github.com/jerson2000/jquest/responses"
	"github.com/jerson2000/jquest/services"
	"github.com/jerson2000/jquest/utils"
)

type recruiterController struct {
	service services.RecruiterService
}

func newRecruiterController() *recruiterController {
	return &recruiterController{service: services.NewRecruiterService()}
}

func (c *recruiterController) registerRoutes(r *gin.RouterGroup) {
	group := r.Group("/recruiters")
	{
		group.GET("", c.getRecruiters)
		group.GET("/company/:id", c.getByCompanyID)
		group.GET("/user/:userId", c.getRecruiterInfoByUserID)
		group.POST("", c.createRecruiter)
	}
}

func (c *recruiterController) createRecruiter(ctx *gin.Context) {
	var dto dtos.RecruiterCreateRequestDto
	if !utils.ValidationShouldBind(http.StatusBadRequest, &dto, trans, ctx) {
		return
	}

	userId := ctx.MustGet("id").(int)
	res := c.service.CreateRecruiter(ctx.Request.Context(), userId, dto)
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}

func (c *recruiterController) getRecruiters(ctx *gin.Context) {
	res := c.service.GetRecruiters(ctx.Request.Context())
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}

func (c *recruiterController) getRecruiterInfoByUserID(ctx *gin.Context) {
	idString := ctx.Param("userId")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Failure[any](http.StatusBadRequest, "invalid id"))
		return
	}
	res := c.service.GetByUserID(ctx.Request.Context(), id)
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}

func (c *recruiterController) getByCompanyID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Failure[any](http.StatusBadRequest, "invalid id"))
		return
	}
	res := c.service.GetByCompanyID(ctx.Request.Context(), id)
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}
