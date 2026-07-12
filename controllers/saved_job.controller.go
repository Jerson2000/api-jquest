package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jerson2000/jquest/dtos"
	"github.com/jerson2000/jquest/services"
	"github.com/jerson2000/jquest/utils"
)

type savedJobController struct {
	service services.SavedJobService
}

func newSavedJobController(service services.SavedJobService) *savedJobController {
	return &savedJobController{service: service}
}

func (c *savedJobController) registerRoutes(protected *gin.RouterGroup) {
	group := protected.Group("/saved-jobs")
	{
		group.POST("", c.saveJob)
		group.DELETE("/:jobId", c.unsaveJob)
		group.GET("", c.getSavedJobs)
	}
}

func (c *savedJobController) saveJob(ctx *gin.Context) {
	userId := ctx.MustGet("id").(int)
	var dto dtos.SavedJobCreateRequestDto
	if !utils.ValidationShouldBind(http.StatusBadRequest, &dto, trans, ctx) {
		return
	}

	res := c.service.SaveJob(ctx.Request.Context(), userId, dto)
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}

func (c *savedJobController) unsaveJob(ctx *gin.Context) {
	userId := ctx.MustGet("id").(int)
	jobIdStr := ctx.Param("jobId")
	jobId, err := strconv.Atoi(jobIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid job id"})
		return
	}

	res := c.service.UnsaveJob(ctx.Request.Context(), userId, jobId)
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}

func (c *savedJobController) getSavedJobs(ctx *gin.Context) {
	userId := ctx.MustGet("id").(int)
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	res := c.service.GetSavedJobs(ctx.Request.Context(), userId, page, limit)
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}
