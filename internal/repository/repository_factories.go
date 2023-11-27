package repository

import (
	"com.rlohmus.checklist/internal/repository/mapper"
	"gorm.io/gorm"
)

func CreateChecklistRepository(db *gorm.DB, mapper mapper.IChecklistDboMapper) IChecklistRepository {
	return &checklistRepository{
		db:     db,
		mapper: mapper,
	}
}

func CreateChecklistItemRepository() IChecklistItemsRepository {
	return &checklistItemRepository{}
}

func CreateChecklistItemTemplateRepository() IChecklistItemTemplateRepository {
	return &checklistItemTemplateRepository{}
}
