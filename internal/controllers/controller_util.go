package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func getChecklistId(c *gin.Context) uint {
	checklistIdParam := c.Param("checklist-id")
	if checklistId, err := strconv.Atoi(checklistIdParam); err != nil {
		panic(errors.New(fmt.Sprintf("Invalid checklist id: %s ", err.Error())))
	} else {
		return uint(checklistId)
	}
}

func getChecklistItemId(c *gin.Context) uint {
	checklistItemIdParam := c.Param("checklist-item-id")
	if checklistItemId, err := strconv.Atoi(checklistItemIdParam); err != nil {
		panic(errors.New(fmt.Sprintf("Invalid checklist item id: %s ", err.Error())))
	} else {
		return uint(checklistItemId)
	}
}

func getChecklistItemTemplateId(c *gin.Context) uint {
	checklistItemTemplateIdParam := c.Param("checklist-item-template-id")
	if checklistItemTemplateId, err := strconv.Atoi(checklistItemTemplateIdParam); err != nil {
		panic(errors.New(fmt.Sprintf("Invalid checklist item template id: %s ", err.Error())))
	} else {
		return uint(checklistItemTemplateId)
	}
}
