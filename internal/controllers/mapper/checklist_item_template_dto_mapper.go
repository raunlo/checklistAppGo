package mapper

import (
	"com.rlohmus.checklist/internal/controllers/dto"
	"com.rlohmus.checklist/internal/domain"
	"github.com/rendis/structsconv"
)

type IChecklistItemTemplateDtoMapper interface {
	ToDomain(checklistTemplateDto dto.ChecklistItemTemplateDto) domain.ChecklistItemTemplate
	ToDTO(checklistTemplate domain.ChecklistItemTemplate) dto.ChecklistItemTemplateDto
}

type checklistItemTemplateDtoMapper struct {
}

func (c checklistItemTemplateDtoMapper) ToDomain(checklistTemplateDto dto.ChecklistItemTemplateDto) domain.ChecklistItemTemplate {
	var checklistItemTemplate domain.ChecklistItemTemplate
	structsconv.Map(&checklistTemplateDto, &checklistItemTemplate)
	return checklistItemTemplate
}

func (c checklistItemTemplateDtoMapper) ToDTO(checklistTemplate domain.ChecklistItemTemplate) dto.ChecklistItemTemplateDto {
	var checklistItemTemplateDto dto.ChecklistItemTemplateDto
	structsconv.Map(&checklistTemplate, &checklistItemTemplateDto)
	return checklistItemTemplateDto
}

func NewChecklistItemTemplateDtoMapper() IChecklistItemTemplateDtoMapper {
	return &checklistItemTemplateDtoMapper{}
}
