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

type jobController struct {
	service services.JobService
}

func newJobController() *jobController {
	return &jobController{service: services.NewJobService()}
}

func (c *jobController) registerRoutes(public *gin.RouterGroup, protected *gin.RouterGroup) {
	pub := public.Group("/jobs")
	{
		pub.GET("", c.getJobs)
		pub.GET("/:id", c.getJobById)
	}

	prot := protected.Group("/jobs")
	{
		prot.POST("", c.createJob)
		prot.PUT("/:id", c.updateJob)
		prot.DELETE("/:id", c.deleteJob)

	}
}

func (c *jobController) createJob(ctx *gin.Context) {
	var dto dtos.JobCreateJobRequestDto
	if !utils.ValidationShouldBind(http.StatusBadRequest, &dto, trans, ctx) {
		return
	}

	userId := ctx.MustGet("id").(int)
	res := c.service.CreateJobPost(ctx.Request.Context(), userId, dto)
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}

func (c *jobController) getJobs(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	res := c.service.GetJobPost(ctx.Request.Context(), page, limit)
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}

func (c *jobController) getJobById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Failure[any](http.StatusBadRequest, "invalid id"))
		return
	}

	res := c.service.GetJobPostById(ctx.Request.Context(), id)
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}

func (c *jobController) updateJob(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Failure[any](http.StatusBadRequest, "invalid id"))
		return
	}

	var dto dtos.JobCreateJobRequestDto
	if !utils.ValidationShouldBind(http.StatusBadRequest, &dto, trans, ctx) {
		return
	}
	userId := ctx.MustGet("id").(int)
	res := c.service.UpdateJobPost(ctx.Request.Context(), userId, id, dto)
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}

func (c *jobController) deleteJob(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Failure[any](http.StatusBadRequest, "invalid id"))
		return
	}
	userId := ctx.MustGet("id").(int)
	res := c.service.DeleteJobPost(ctx.Request.Context(), userId, id)
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}
