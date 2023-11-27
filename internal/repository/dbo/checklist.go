package dbo

type ChecklistDbo struct {
	Id   uint   `gorm:"primaryKey;column:checklist_id" sql:"nextval('checklist_id_sequence')"`
	Name string `gorm:"column:checklist_name"`
}

func (ChecklistDbo) TableName() string {
	// sets table name for ChecklistDbo model
	return "checklist"
}
