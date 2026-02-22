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

type applicationController struct {
	service services.ApplicationService
}

func newApplicationController() *applicationController {
	return &applicationController{service: services.NewApplicationService()}
}

func (c *applicationController) registerRoutes(public *gin.RouterGroup, protected *gin.RouterGroup) {
	prot := protected.Group("/applications")
	{
		prot.POST("", c.applyJob)
		prot.GET("/my", c.getMyApplications)
		prot.GET("/job/:jobId", c.getJobApplications)
		prot.PATCH("/:id/status", c.updateApplicationStatus)
	}
}

func (c *applicationController) applyJob(ctx *gin.Context) {
	userId := ctx.MustGet("id").(int)
	var dto dtos.ApplicationCreateRequestDto
	if !utils.ValidationShouldBind(http.StatusBadRequest, &dto, trans, ctx) {
		return
	}

	res := c.service.ApplyJob(ctx.Request.Context(), userId, dto)
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}

func (c *applicationController) getMyApplications(ctx *gin.Context) {
	userId := ctx.MustGet("id").(int)
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	res := c.service.GetMyApplications(ctx.Request.Context(), userId, page, limit)
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}

func (c *applicationController) getJobApplications(ctx *gin.Context) {
	userId := ctx.MustGet("id").(int)
	jobId, err := strconv.Atoi(ctx.Param("jobId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Failure[any](http.StatusBadRequest, "Invalid Job ID"))
		return
	}

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	res := c.service.GetJobApplications(ctx.Request.Context(), userId, jobId, page, limit)
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}

func (c *applicationController) updateApplicationStatus(ctx *gin.Context) {
	userId := ctx.MustGet("id").(int)
	applicationId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Failure[any](http.StatusBadRequest, "Invalid Application ID"))
		return
	}

	var dto dtos.ApplicationUpdateStatusRequestDto
	if !utils.ValidationShouldBind(http.StatusBadRequest, &dto, trans, ctx) {
		return
	}

	res := c.service.UpdateApplicationStatus(ctx.Request.Context(), userId, applicationId, dto)
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}
