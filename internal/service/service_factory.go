package service

import "com.rlohmus.checklist/internal/repository"

func CreateChecklistService(checklistRepository repository.IChecklistRepository) IChecklistService {
	return &checklistService{
		repository: checklistRepository,
	}
}

func CreateChecklistItemService(repository repository.IChecklistItemsRepository) IChecklistItemsService {
	return &checklistItemsService{
		repository: repository,
	}
}

func CreateChecklistItemTemplateService(repository repository.IChecklistItemTemplateRepository) IChecklistItemTemplateService {
	return &checklistItemTemplateService{
		repository: repository,
	}
}
