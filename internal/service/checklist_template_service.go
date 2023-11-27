package service

import (
	"com.rlohmus.checklist/internal/domain"
	"com.rlohmus.checklist/internal/repository"
)

type IChecklistItemTemplateService interface {
	SaveChecklistTemplate(checklistTemplate domain.ChecklistItemTemplate) (domain.ChecklistItemTemplate, domain.Error)
	GetAllChecklistTemplates() ([]domain.ChecklistItemTemplate, domain.Error)
	UpdateChecklistTemplate(checklistTemplate domain.ChecklistItemTemplate) (domain.ChecklistItemTemplate, domain.Error)
	DeleteChecklistTemplateById(id uint) domain.Error
	FindChecklistTemplateById(id uint) (*domain.ChecklistItemTemplate, domain.Error)
}

type checklistItemTemplateService struct {
	repository repository.IChecklistItemTemplateRepository
}

func (c *checklistItemTemplateService) SaveChecklistTemplate(checklistTemplate domain.ChecklistItemTemplate) (domain.ChecklistItemTemplate, domain.Error) {
	//TODO implement me
	panic("implement me")
}

func (c *checklistItemTemplateService) GetAllChecklistTemplates() ([]domain.ChecklistItemTemplate, domain.Error) {
	//TODO implement me
	panic("implement me")
}

func (c *checklistItemTemplateService) UpdateChecklistTemplate(checklistTemplate domain.ChecklistItemTemplate) (domain.ChecklistItemTemplate, domain.Error) {
	//TODO implement me
	panic("implement me")
}

func (c *checklistItemTemplateService) DeleteChecklistTemplateById(id uint) domain.Error {
	//TODO implement me
	panic("implement me")
}

func (c *checklistItemTemplateService) FindChecklistTemplateById(id uint) (*domain.ChecklistItemTemplate, domain.Error) {
	//TODO implement me
	panic("implement me")
}
