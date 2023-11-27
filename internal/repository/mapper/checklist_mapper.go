package mapper

import (
	"com.rlohmus.checklist/internal/domain"
	"com.rlohmus.checklist/internal/repository/dbo"
	"github.com/rendis/structsconv"
)

type IChecklistDboMapper interface {
	ToDomain(checklistDbo dbo.ChecklistDbo) domain.Checklist
	ToDbo(checklist domain.Checklist) dbo.ChecklistDbo
}

type checklistMapper struct {
}

func (_ checklistMapper) ToDomain(checklistDbo dbo.ChecklistDbo) domain.Checklist {
	var checklist domain.Checklist
	structsconv.Map(&checklistDbo, &checklist)
	return checklist
}

func (_ checklistMapper) ToDbo(checklist domain.Checklist) dbo.ChecklistDbo {
	var checklistDbo dbo.ChecklistDbo
	structsconv.Map(&checklist, &checklistDbo)
	return checklistDbo
}

func NewChecklistDboMapper() IChecklistDboMapper {
	return &checklistMapper{}
}
