package controllers

import (
	"com.rlohmus.checklist/internal/controllers/mapper"
	"com.rlohmus.checklist/internal/service"
)

func CreateChecklistController(service service.IChecklistService, mapper mapper.IChecklistDtoMapper) IChecklistController {
	return &checklistController{
		service: service,
		mapper:  mapper,
	}
}

func CreateChecklistItemsController(service service.IChecklistItemsService) IChecklistItemController {
	return &checklistItemController{
		service: service,
	}
}

func CreateChecklistItemTemplateController(service service.IChecklistItemTemplateService,
	mapper mapper.IChecklistItemTemplateDtoMapper) IChecklistItemTemplateController {
	return &checklistItemTemplateController{
		service: service,
		mapper:  mapper,
	}
}
