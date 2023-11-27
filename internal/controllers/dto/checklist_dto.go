package dto

type ChecklistDTO struct {
	Id             uint               `json:id`
	Name           string             `json:name`
	ChecklistItems []ChecklistItemDto `json:"checklist_items"`
}
