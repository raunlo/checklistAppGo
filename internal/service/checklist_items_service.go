package service

import (
	domain "com.rlohmus.checklist/internal/domain"
	repository "com.rlohmus.checklist/internal/repository"
)

type IChecklistItemsService interface {
	SaveChecklistItem(checklistId uint, checklistItem domain.ChecklistItem) (domain.ChecklistItem, domain.Error)
	UpdateChecklistItem(checklistId uint, checklistItem domain.ChecklistItem) (domain.ChecklistItem, domain.Error)
	FindChecklistItemById(checklistId uint, id uint) (*domain.ChecklistItem, domain.Error)
	DeleteChecklistItemById(checklistId uint, id uint) domain.Error
	FindAllChecklistItems(checklistId uint) ([]domain.ChecklistItem, domain.Error)
}

type checklistItemsService struct {
	repository repository.IChecklistItemsRepository
}

func (service *checklistItemsService) UpdateChecklistItem(checklistId uint, checklistItem domain.ChecklistItem) (domain.ChecklistItem, domain.Error) {
	return service.repository.UpdateChecklistItem(checklistId, checklistItem)
}

func (service *checklistItemsService) SaveChecklistItem(checklistId uint, checklistItem domain.ChecklistItem) (domain.ChecklistItem, domain.Error) {
	return service.repository.SaveChecklistItem(checklistId, checklistItem)
}

func (service *checklistItemsService) FindChecklistItemById(checklistId uint, id uint) (*domain.ChecklistItem, domain.Error) {
	return service.repository.FindChecklistItemById(checklistId, id)
}

func (service *checklistItemsService) DeleteChecklistItemById(checklistId uint, id uint) domain.Error {
	return service.repository.DeleteChecklistItemById(checklistId, id)
}

func (service *checklistItemsService) FindAllChecklistItems(checklistId uint) ([]domain.ChecklistItem, domain.Error) {
	return service.repository.FindAllChecklistItems(checklistId)
}
