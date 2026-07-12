package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jerson2000/jquest/dtos"
	"github.com/jerson2000/jquest/services"
	"github.com/jerson2000/jquest/utils"
)

type skillController struct {
	service services.SkillService
}

func newSkillController(service services.SkillService) *skillController {
	return &skillController{service: service}
}

func (c *skillController) registerRoutes(public *gin.RouterGroup, protected *gin.RouterGroup) {
	publicGroup := public.Group("/skills")
	{
		publicGroup.GET("", c.getSkills)
	}

	protectedGroup := protected.Group("/skills")
	{
		protectedGroup.POST("/candidate", c.addSkillToCandidate)
		protectedGroup.DELETE("/candidate", c.removeSkillFromCandidate)
		protectedGroup.POST("/job/:jobId", c.addSkillToJob)
		protectedGroup.DELETE("/job/:jobId", c.removeSkillFromJob)
	}
}

func (c *skillController) addSkillToCandidate(ctx *gin.Context) {
	userId := ctx.MustGet("id").(int)
	var dto dtos.SkillCreateRequestDto
	if !utils.ValidationShouldBind(http.StatusBadRequest, &dto, trans, ctx) {
		return
	}

	res := c.service.AddSkillToCandidate(ctx.Request.Context(), userId, dto.Name)
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}

func (c *skillController) removeSkillFromCandidate(ctx *gin.Context) {
	userId := ctx.MustGet("id").(int)
	var dto dtos.SkillCreateRequestDto
	if !utils.ValidationShouldBind(http.StatusBadRequest, &dto, trans, ctx) {
		return
	}

	res := c.service.RemoveSkillFromCandidate(ctx.Request.Context(), userId, dto.Name)
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}

func (c *skillController) addSkillToJob(ctx *gin.Context) {
	recruiterUserId := ctx.MustGet("id").(int)
	jobIdStr := ctx.Param("jobId")
	jobId, err := strconv.Atoi(jobIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid job id"})
		return
	}

	var dto dtos.SkillCreateRequestDto
	if !utils.ValidationShouldBind(http.StatusBadRequest, &dto, trans, ctx) {
		return
	}

	res := c.service.AddSkillToJob(ctx.Request.Context(), recruiterUserId, jobId, dto.Name)
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}

func (c *skillController) removeSkillFromJob(ctx *gin.Context) {
	recruiterUserId := ctx.MustGet("id").(int)
	jobIdStr := ctx.Param("jobId")
	jobId, err := strconv.Atoi(jobIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid job id"})
		return
	}

	var dto dtos.SkillCreateRequestDto
	if !utils.ValidationShouldBind(http.StatusBadRequest, &dto, trans, ctx) {
		return
	}

	res := c.service.RemoveSkillFromJob(ctx.Request.Context(), recruiterUserId, jobId, dto.Name)
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}

func (c *skillController) getSkills(ctx *gin.Context) {
	res := c.service.GetSkills(ctx.Request.Context())
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}
