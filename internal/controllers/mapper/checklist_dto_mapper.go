package mapper

import (
	"com.rlohmus.checklist/internal/controllers/dto"
	"com.rlohmus.checklist/internal/domain"
	"github.com/rendis/structsconv"
)

type IChecklistDtoMapper interface {
	ToDomain(source dto.ChecklistDTO) domain.Checklist
	ToDTO(source domain.Checklist) dto.ChecklistDTO
	ToDtoArray(checklists []domain.Checklist) []dto.ChecklistDTO
}

type checklistDtoMapper struct {
}

func NewChecklistDtoMapper() IChecklistDtoMapper {
	structsconv.RegisterRulesDefinitions(getDtoToDomainRules())

	return &checklistDtoMapper{}
}

func getDtoToDomainRules() structsconv.RulesDefinition {
	rules := structsconv.RulesSet{}
	rules["ChecklistItems"] = nil
	rules["Id"] = nil

	return structsconv.RulesDefinition{
		Rules:  rules,
		Source: dto.ChecklistDTO{},
		Target: domain.Checklist{},
	}
}

func (_ *checklistDtoMapper) ToDomain(source dto.ChecklistDTO) domain.Checklist {
	target := domain.Checklist{}
	structsconv.Map(&source, &target)
	return target
}

func (_ *checklistDtoMapper) ToDTO(source domain.Checklist) dto.ChecklistDTO {
	structsconv.RegisterRulesDefinitions()
	target := dto.ChecklistDTO{}
	structsconv.Map(&source, &target)
	return target
}

func (mapper *checklistDtoMapper) ToDtoArray(checklists []domain.Checklist) []dto.ChecklistDTO {
	var checklistDtoArray []dto.ChecklistDTO

	for _, checklist := range checklists {
		checklistDtoArray = append(checklistDtoArray, mapper.ToDTO(checklist))
	}
	return checklistDtoArray
}
