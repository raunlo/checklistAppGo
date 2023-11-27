package controllers

import (
	"com.rlohmus.checklist/internal/controllers/dto"
	"com.rlohmus.checklist/internal/controllers/mapper"
	"com.rlohmus.checklist/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IChecklistItemController interface {
	SaveChecklistItem(c *gin.Context)
	UpdateChecklistItem(c *gin.Context)
	DeleteChecklistItemById(c *gin.Context)
	FindChecklistItemById(c *gin.Context)
	FindAllChecklistItems(c *gin.Context)
}

type checklistItemController struct {
	service service.IChecklistItemsService
	mapper  mapper.IChecklistItemDtoMapper
}

func (controller *checklistItemController) SaveChecklistItem(c *gin.Context) {
	var checklistItemDto dto.ChecklistItemDto
	if err := c.Bind(&checklistItemDto); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	checklistId := getChecklistId(c)

	checklistItem := controller.mapper.MapDtoToDomain(checklistItemDto)
	if savedChecklistItem, err := controller.service.SaveChecklistItem(checklistId, checklistItem); err != nil {
		c.JSON(err.ResponseCode(), dto.ErrorDto{Message: err.Error()})
	} else {
		savedChecklistItem := controller.mapper.MapDomainToDto(savedChecklistItem)
		c.JSON(http.StatusCreated, savedChecklistItem)
	}
}

func (controller *checklistItemController) UpdateChecklistItem(c *gin.Context) {
	var checklistItemDto dto.ChecklistItemDto
	if err := c.Bind(&checklistItemDto); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	checklistId := getChecklistId(c)

	checklistItem := controller.mapper.MapDtoToDomain(checklistItemDto)
	checklistItem.Id = getChecklistItemId(c)

	if savedChecklistItem, err := controller.service.UpdateChecklistItem(checklistId, checklistItem); err != nil {
		c.JSON(err.ResponseCode(), dto.ErrorDto{Message: err.Error()})
	} else {
		savedChecklistItem := controller.mapper.MapDomainToDto(savedChecklistItem)
		c.JSON(http.StatusCreated, savedChecklistItem)
	}
}

func (controller *checklistItemController) DeleteChecklistItemById(c *gin.Context) {
	checklistId := getChecklistId(c)
	checklistItemId := getChecklistItemId(c)

	if err := controller.service.DeleteChecklistItemById(checklistId, checklistItemId); err != nil {
		c.JSON(err.ResponseCode(), dto.ErrorDto{Message: err.Error()})
	} else {
		c.Status(http.StatusNoContent)
	}
}

func (controller *checklistItemController) FindChecklistItemById(c *gin.Context) {
	checklistId := getChecklistId(c)
	checklistItemId := getChecklistItemId(c)

	if checklistItem, err := controller.service.FindChecklistItemById(checklistId, checklistItemId); err != nil {
		c.JSON(err.ResponseCode(), dto.ErrorDto{Message: err.Error()})
	} else if checklistItem == nil {
		c.Status(404)
	} else {
		c.JSON(http.StatusOK, controller.mapper.MapDomainToDto(*checklistItem))
	}
}

func (controller *checklistItemController) FindAllChecklistItems(c *gin.Context) {
	checklistId := getChecklistId(c)

	if checklistItems, err := controller.service.FindAllChecklistItems(checklistId); err != nil {
		c.JSON(err.ResponseCode(), dto.ErrorDto{Message: err.Error()})
	} else {
		c.JSON(200, controller.mapper.MapDomainListToDtoList(checklistItems))
	}
}
