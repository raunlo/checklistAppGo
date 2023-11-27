package infra

import (
	"com.rlohmus.checklist/internal/controllers"
	"github.com/gin-gonic/gin"
)

type HttpRouterV1 struct {
	checklistItemController         controllers.IChecklistItemController
	checklistController             controllers.IChecklistController
	checklistItemTemplateController controllers.IChecklistItemTemplateController
}

func CreateHttpV1Router(
	checklistItemController controllers.IChecklistItemController,
	checklistController controllers.IChecklistController,
	checklistItemTemplateController controllers.IChecklistItemTemplateController) HttpRouterV1 {
	return HttpRouterV1{
		checklistController:             checklistController,
		checklistItemController:         checklistItemController,
		checklistItemTemplateController: checklistItemTemplateController,
	}
}

func (route *HttpRouterV1) CreateRoutes(routeGroup *gin.RouterGroup) {
	routeV1 := routeGroup.Group("/v1")
	route.createChecklistRouteGroup(routeV1.Group("/checklist"))
	route.createChecklistItemRouteGroup(routeV1.Group("/checklist/:checklist-id/item"))
	route.createChecklistItemTemplateRouteGroup(routeV1.Group("/checklist-item-template"))
}

func (route *HttpRouterV1) createChecklistRouteGroup(routeGroup *gin.RouterGroup) {
	routeGroup.POST("", route.checklistController.SaveChecklist)
	routeGroup.GET("", route.checklistController.GetAllChecklists)
	routeGroup.GET("/:checklist-id", route.checklistController.FindChecklistById)
	routeGroup.DELETE("/:checklist-id", route.checklistController.DeleteChecklistById)
	routeGroup.PUT("/:checklist-id", route.checklistController.UpdateChecklist)
}

func (route *HttpRouterV1) createChecklistItemRouteGroup(routeGroup *gin.RouterGroup) {
	routeGroup.POST("", route.checklistItemController.SaveChecklistItem)
	routeGroup.GET("", route.checklistItemController.FindAllChecklistItems)
	routeGroup.GET("/:checklist-item-id", route.checklistItemController.FindChecklistItemById)
	routeGroup.DELETE("/:checklist-item-id", route.checklistItemController.DeleteChecklistItemById)
	routeGroup.PUT("/:checklist-item-id", route.checklistItemController.UpdateChecklistItem)
}

func (route *HttpRouterV1) createChecklistItemTemplateRouteGroup(routeGroup *gin.RouterGroup) {
	routeGroup.POST("", route.checklistItemTemplateController.SaveChecklistTemplate)
	routeGroup.GET("", route.checklistItemTemplateController.GetAllChecklistTemplates)
	routeGroup.GET("/:checklist-item-template-id", route.checklistItemTemplateController.FindChecklistTemplateById)
	routeGroup.DELETE("/:checklist-item-template-id", route.checklistItemTemplateController.DeleteChecklistTemplateById)
	routeGroup.PUT("/:checklist-item-template-id", route.checklistItemTemplateController.UpdateChecklistTemplate)
}
