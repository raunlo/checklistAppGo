package controllers

import (
	"com.rlohmus.checklist/internal/controllers/dto"
	"com.rlohmus.checklist/internal/controllers/mapper"
	"com.rlohmus.checklist/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IChecklistItemTemplateController interface {
	SaveChecklistTemplate(c *gin.Context)
	GetAllChecklistTemplates(c *gin.Context)
	UpdateChecklistTemplate(c *gin.Context)
	DeleteChecklistTemplateById(c *gin.Context)
	FindChecklistTemplateById(c *gin.Context)
}

type checklistItemTemplateController struct {
	service service.IChecklistItemTemplateService
	mapper  mapper.IChecklistItemTemplateDtoMapper
}

func (controller *checklistItemTemplateController) SaveChecklistTemplate(c *gin.Context) {
	var checklistItemTemplateDTO dto.ChecklistItemTemplateDto
	if err := c.Bind(&checklistItemTemplateDTO); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	checklistItemTemplate := controller.mapper.ToDomain(checklistItemTemplateDTO)

	if saveChecklistTemplate, err := controller.service.SaveChecklistTemplate(checklistItemTemplate); err != nil {
		c.JSON(err.ResponseCode(), dto.ErrorDto{Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, controller.mapper.ToDTO(saveChecklistTemplate))
	}
}

func (controller *checklistItemTemplateController) GetAllChecklistTemplates(c *gin.Context) {
	if checklistItems, err := controller.service.GetAllChecklistTemplates(); err != nil {
		c.JSON(err.ResponseCode(), dto.ErrorDto{Message: err.Error()})
	} else {
		var checklistItemTemplateDTOs []dto.ChecklistItemTemplateDto
		for _, checklistITemTemplate := range checklistItems {
			checklistItemTemplateDTOs = append(checklistItemTemplateDTOs, controller.mapper.ToDTO(checklistITemTemplate))
		}
		c.JSON(200, checklistItemTemplateDTOs)
	}
}

func (controller *checklistItemTemplateController) UpdateChecklistTemplate(c *gin.Context) {
	var checklistItemTemplateDTO dto.ChecklistItemTemplateDto
	if err := c.Bind(&checklistItemTemplateDTO); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	checklistItemTemplate := controller.mapper.ToDomain(checklistItemTemplateDTO)

	checklistItemTemplate.Id = getChecklistItemTemplateId(c)

	if updateChecklistTemplate, err := controller.service.UpdateChecklistTemplate(checklistItemTemplate); err != nil {
		c.JSON(err.ResponseCode(), dto.ErrorDto{Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, controller.mapper.ToDTO(updateChecklistTemplate))
	}
}

func (controller *checklistItemTemplateController) DeleteChecklistTemplateById(c *gin.Context) {
	checklistItemTemplateId := getChecklistItemTemplateId(c)

	if err := controller.service.DeleteChecklistTemplateById(checklistItemTemplateId); err != nil {
		c.JSON(err.ResponseCode(), dto.ErrorDto{Message: err.Error()})
	} else {
		c.Status(http.StatusNoContent)
	}
}

func (controller *checklistItemTemplateController) FindChecklistTemplateById(c *gin.Context) {
	checklistItemTemplateId := getChecklistItemTemplateId(c)

	if checklist, err := controller.service.FindChecklistTemplateById(checklistItemTemplateId); err != nil {
		c.JSON(err.ResponseCode(), dto.ErrorDto{Message: err.Error()})
	} else if checklist == nil {
		c.JSON(404, dto.ErrorDto{Message: "Checklist not found"})
	} else {
		c.JSON(http.StatusOK, controller.mapper.ToDTO(*checklist))
	}
}
