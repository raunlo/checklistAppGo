package repository

import (
	domain "com.rlohmus.checklist/internal/domain"
	gorm "gorm.io/gorm"
)

type IChecklistItemsRepository interface {
	UpdateChecklistItem(checklistId uint, checklistItem domain.ChecklistItem) (domain.ChecklistItem, domain.Error)
	SaveChecklistItem(checklistId uint, checklistItem domain.ChecklistItem) (domain.ChecklistItem, domain.Error)
	FindChecklistItemById(checklistId uint, id uint) (*domain.ChecklistItem, domain.Error)
	DeleteChecklistItemById(checklistId uint, id uint) domain.Error
	FindAllChecklistItems(checklistId uint) ([]domain.ChecklistItem, domain.Error)
}

type checklistItemRepository struct {
	db *gorm.DB
}

func (repository *checklistItemRepository) UpdateChecklistItem(checklistId uint, checklistItem domain.ChecklistItem) (domain.ChecklistItem, domain.Error) {
	return domain.ChecklistItem{}, nil
}

func (repository *checklistItemRepository) SaveChecklistItem(checklistId uint, checklistItem domain.ChecklistItem) (domain.ChecklistItem, domain.Error) {
	return domain.ChecklistItem{}, nil
}

func (repository *checklistItemRepository) FindChecklistItemById(checklistId uint, id uint) (*domain.ChecklistItem, domain.Error) {
	return &domain.ChecklistItem{}, nil
}

func (repository *checklistItemRepository) DeleteChecklistItemById(checklistId uint, id uint) domain.Error {
	return nil
}

func (repository *checklistItemRepository) FindAllChecklistItems(checklistId uint) ([]domain.ChecklistItem, domain.Error) {
	return []domain.ChecklistItem{}, nil
}
