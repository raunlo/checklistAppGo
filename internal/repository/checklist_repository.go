package repository

import (
	"com.rlohmus.checklist/internal/domain"
	"com.rlohmus.checklist/internal/repository/dbo"
	"com.rlohmus.checklist/internal/repository/mapper"
	"database/sql"
	"gorm.io/gorm"
)

type IChecklistRepository interface {
	UpdateChecklist(checklist domain.Checklist) (domain.Checklist, domain.Error)
	SaveChecklist(checklist domain.Checklist) (domain.Checklist, domain.Error)
	FindChecklistById(id uint) (*domain.Checklist, domain.Error)
	DeleteChecklistById(id uint) domain.Error
	FindAllChecklists() ([]domain.Checklist, domain.Error)
}

type checklistRepository struct {
	db     *gorm.DB
	mapper mapper.IChecklistDboMapper
}

func (repository *checklistRepository) UpdateChecklist(checklist domain.Checklist) (domain.Checklist, domain.Error) {
	updateChecklistChannel := make(chan dbo.QueryResult[dbo.ChecklistDbo], 1)
	go func(checklistDbo dbo.ChecklistDbo, channel chan<- dbo.QueryResult[dbo.ChecklistDbo]) {
		queryResult := repository.db.Save(&checklistDbo)
		channel <- dbo.QueryResult[dbo.ChecklistDbo]{
			Error:  queryResult.Error,
			Result: checklistDbo,
		}

	}(repository.mapper.ToDbo(checklist), updateChecklistChannel)

	if result := <-updateChecklistChannel; result.Error != nil {
		return domain.Checklist{}, domain.NewError(result.Error.Error(), 500)
	} else {
		return repository.mapper.ToDomain(result.Result), nil
	}
}

func (repository *checklistRepository) SaveChecklist(checklist domain.Checklist) (domain.Checklist, domain.Error) {
	createChecklistChannel := make(chan dbo.QueryResult[dbo.ChecklistDbo], 1)
	go func(checklistDbo dbo.ChecklistDbo, channel chan<- dbo.QueryResult[dbo.ChecklistDbo]) {
		insertQuery := "INSERT INTO CHECKLIST(CHECKLIST_ID, CHECKLIST_NAME) VALUES(nextval('checklist_id_sequence'), @checklistName) RETURNING CHECKLIST_ID "
		queryResult := repository.db.Raw(insertQuery, sql.Named("checklistName", checklistDbo.Name)).Scan(&checklistDbo.Id)

		channel <- dbo.QueryResult[dbo.ChecklistDbo]{
			Error:  queryResult.Error,
			Result: checklistDbo,
		}

	}(repository.mapper.ToDbo(checklist), createChecklistChannel)

	if result := <-createChecklistChannel; result.Error != nil {
		return domain.Checklist{}, domain.NewError(result.Error.Error(), 500)
	} else {
		return repository.mapper.ToDomain(result.Result), nil
	}
}

func (repository *checklistRepository) FindChecklistById(id uint) (*domain.Checklist, domain.Error) {
	checklistChannel := make(chan dbo.QueryResult[*dbo.ChecklistDbo], 1)
	go func(id uint, checklistChannel chan<- dbo.QueryResult[*dbo.ChecklistDbo]) {
		var checklist dbo.ChecklistDbo
		queryResult := repository.db.Find(&checklist, id)
		if queryResult.RowsAffected == 0 {
			checklistChannel <- dbo.QueryResult[*dbo.ChecklistDbo]{
				Error:  queryResult.Error,
				Result: nil,
			}
		} else {
			checklistChannel <- dbo.QueryResult[*dbo.ChecklistDbo]{
				Error:  queryResult.Error,
				Result: &checklist,
			}
		}
	}(id, checklistChannel)

	if result := <-checklistChannel; result.Error != nil {
		return nil, domain.NewError(result.Error.Error(), 500)
	} else {
		if result.Result == nil {
			return nil, nil
		} else {
			checklist := repository.mapper.ToDomain(*result.Result)
			return &checklist, nil
		}
	}
}

func (repository *checklistRepository) DeleteChecklistById(id uint) domain.Error {
	var checklistDeletedChannel = make(chan dbo.QueryResult[int], 1)
	go func(id uint, checklistDeletedChannel chan<- dbo.QueryResult[int]) {
		queryResult := repository.db.Delete(&dbo.ChecklistDbo{}, id)
		checklistDeletedChannel <- dbo.QueryResult[int]{
			Error:  queryResult.Error,
			Result: int(queryResult.RowsAffected),
		}
	}(id, checklistDeletedChannel)

	if result := <-checklistDeletedChannel; result.Error != nil {
		return domain.NewError(result.Error.Error(), 500)
	} else {
		if result.Result == 0 {
			return domain.NewError("Checklist not found", 404)
		} else {
			return nil
		}
	}
}

func (repository *checklistRepository) FindAllChecklists() ([]domain.Checklist, domain.Error) {
	var checklists []domain.Checklist
	checklistChannel := make(chan dbo.QueryResult[[]dbo.ChecklistDbo], 1)
	go func(checklistChannel chan<- dbo.QueryResult[[]dbo.ChecklistDbo]) {
		var checklists []dbo.ChecklistDbo
		dbError := repository.db.Find(&checklists).Error
		checklistChannel <- dbo.QueryResult[[]dbo.ChecklistDbo]{
			Error:  dbError,
			Result: checklists,
		}

	}(checklistChannel)

	if result := <-checklistChannel; result.Error != nil {
		return nil, domain.NewError(result.Error.Error(), 500)
	} else {
		for _, checklistDbo := range result.Result {
			checklists = append(checklists, repository.mapper.ToDomain(checklistDbo))
		}
	}

	return checklists, nil
}
