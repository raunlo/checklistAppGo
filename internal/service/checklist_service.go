package service

import (
	domain "com.rlohmus.checklist/internal/domain"
	repository "com.rlohmus.checklist/internal/repository"
)

type IChecklistService interface {
	UpdateChecklist(checklist domain.Checklist) (domain.Checklist, domain.Error)
	SaveChecklist(checklist domain.Checklist) (domain.Checklist, domain.Error)
	FindChecklistById(id uint) (*domain.Checklist, domain.Error)
	DeleteChecklistById(id uint) domain.Error
	FindAllChecklists() ([]domain.Checklist, domain.Error)
}

type checklistService struct {
	repository repository.IChecklistRepository
}

func (service *checklistService) UpdateChecklist(checklist domain.Checklist) (domain.Checklist, domain.Error) {
	return service.repository.UpdateChecklist(checklist)
}

func (service *checklistService) SaveChecklist(checklist domain.Checklist) (domain.Checklist, domain.Error) {
	return service.repository.SaveChecklist(checklist)
}

func (service *checklistService) FindChecklistById(id uint) (*domain.Checklist, domain.Error) {
	return service.repository.FindChecklistById(id)
}

func (service *checklistService) DeleteChecklistById(id uint) domain.Error {
	return service.repository.DeleteChecklistById(id)
}

func (service *checklistService) FindAllChecklists() ([]domain.Checklist, domain.Error) {
	return service.repository.FindAllChecklists()
}
