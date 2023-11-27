package controllers

import (
	"com.rlohmus.checklist/internal/controllers/dto"
	"com.rlohmus.checklist/internal/controllers/mapper"
	"com.rlohmus.checklist/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IChecklistController interface {
	SaveChecklist(c *gin.Context)
	GetAllChecklists(c *gin.Context)
	UpdateChecklist(c *gin.Context)
	DeleteChecklistById(c *gin.Context)
	FindChecklistById(c *gin.Context)
}

type checklistController struct {
	service service.IChecklistService
	mapper  mapper.IChecklistDtoMapper
}

func (controller *checklistController) SaveChecklist(c *gin.Context) {
	var checklistDTO dto.ChecklistDTO
	if err := c.Bind(&checklistDTO); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	checklist := controller.mapper.ToDomain(checklistDTO)
	if savedChecklist, err := controller.service.SaveChecklist(checklist); err != nil {
		c.JSON(err.ResponseCode(), dto.ErrorDto{Message: err.Error()})
	} else {
		c.JSON(http.StatusCreated, controller.mapper.ToDTO(savedChecklist))
	}
}

func (controller *checklistController) GetAllChecklists(c *gin.Context) {
	if checklists, err := controller.service.FindAllChecklists(); err != nil {
		c.JSON(err.ResponseCode(), dto.ErrorDto{Message: err.Error()})
	} else {
		c.JSON(200, controller.mapper.ToDtoArray(checklists))
	}
}

func (controller *checklistController) UpdateChecklist(c *gin.Context) {
	var checklistDTO dto.ChecklistDTO
	if err := c.Bind(&checklistDTO); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	checklist := controller.mapper.ToDomain(checklistDTO)

	checklist.Id = getChecklistId(c)

	if updatedChecklist, err := controller.service.UpdateChecklist(checklist); err != nil {
		c.JSON(err.ResponseCode(), dto.ErrorDto{Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, controller.mapper.ToDTO(updatedChecklist))
	}
}

func (controller *checklistController) DeleteChecklistById(c *gin.Context) {
	checklistId := getChecklistId(c)

	if err := controller.service.DeleteChecklistById(checklistId); err != nil {
		c.JSON(err.ResponseCode(), dto.ErrorDto{Message: err.Error()})
	}

	c.Status(http.StatusNoContent)
}

func (controller *checklistController) FindChecklistById(c *gin.Context) {
	checklistId := getChecklistId(c)

	if checklist, err := controller.service.FindChecklistById(checklistId); err != nil {
		c.JSON(err.ResponseCode(), dto.ErrorDto{Message: err.Error()})
	} else if checklist == nil {
		c.JSON(404, dto.ErrorDto{Message: "Checklist not found"})
	} else {
		c.JSON(http.StatusOK, controller.mapper.ToDTO(*checklist))
	}
}
