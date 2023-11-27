package mapper

import (
	"com.rlohmus.checklist/internal/controllers/dto"
	"com.rlohmus.checklist/internal/domain"
	"github.com/rendis/structsconv"
)

type IChecklistItemDtoMapper interface {
	MapDomainToDto(checklistItem domain.ChecklistItem) dto.ChecklistItemDto
	MapDtoToDomain(checklistItemDto dto.ChecklistItemDto) domain.ChecklistItem
	MapDomainListToDtoList(checklistItems []domain.ChecklistItem) []dto.ChecklistItemDto
}

type checklistItemMapper struct {
}

func NewChecklistItemDtoMapper() IChecklistItemDtoMapper {
	return &checklistItemMapper{}
}

func (mapper *checklistItemMapper) MapDomainListToDtoList(checklistItems []domain.ChecklistItem) []dto.ChecklistItemDto {
	checklistItemsDtoList := make([]dto.ChecklistItemDto, len(checklistItems))
	for index, item := range checklistItems {
		checklistItemsDtoList[index] = mapper.MapDomainToDto(item)
	}

	return checklistItemsDtoList
}

func (mapper *checklistItemMapper) MapDomainToDto(checklistItem domain.ChecklistItem) dto.ChecklistItemDto {
	checklistItemDto := dto.ChecklistItemDto{}
	structsconv.Map(&checklistItem, &checklistItemDto)
	return checklistItemDto
}

func (mapper *checklistItemMapper) MapDtoToDomain(checklistItemDto dto.ChecklistItemDto) domain.ChecklistItem {
	checklistItem := domain.ChecklistItem{}
	structsconv.Map(&checklistItemDto, &checklistItem)
	return checklistItem
}
