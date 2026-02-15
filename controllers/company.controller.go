package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jerson2000/jquest/dtos"
	"github.com/jerson2000/jquest/services"
	"github.com/jerson2000/jquest/utils"
)

type companyController struct {
	service services.CompanyService
}

func newCompanyController() *companyController {
	return &companyController{service: services.NewCompanyService()}
}

func (c *companyController) registerRoutes(public *gin.RouterGroup, protected *gin.RouterGroup) {
	pub := public.Group("/companies")
	{
		pub.GET("", c.getCompanies)
		pub.POST("/apply", c.applyAsRecruiter)
	}

	prot := protected.Group("/companies")
	{
		prot.POST("", c.createCompany)
	}
}

func (c *companyController) createCompany(ctx *gin.Context) {
	var dto dtos.CompanyCreateRequestDto
	if !utils.ValidationShouldBind(http.StatusBadRequest, &dto, trans, ctx) {
		return
	}

	userId := ctx.MustGet("id").(int)
	res := c.service.CreateCompany(ctx.Request.Context(), userId, dto)
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}

func (c *companyController) getCompanies(ctx *gin.Context) {
	res := c.service.GetCompanies(ctx.Request.Context())
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}

func (c *companyController) applyAsRecruiter(ctx *gin.Context) {
	var dto dtos.CompanyApplyRequestDto
	if !utils.ValidationShouldBind(http.StatusBadRequest, &dto, trans, ctx) {
		return
	}
	res := c.service.ApplyAsRecruiter(ctx.Request.Context(), dto)
	ctx.JSON(utils.ToHTTPStatus(res.Status), res)
}
