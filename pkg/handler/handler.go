package handler

import (
	"github.com/StanislavDimitrenco/restapi/pkg/service"
	"github.com/gin-gonic/gin"
)

//Handler struct
type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

//InitRouts - initializing routs
func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sing-up", h.singUp)
		auth.POST("sing-in", h.singIn)
	}

	api := router.Group("/api")
	{
		list := api.Group("/list")
		{
			list.POST("/", h.createList)
			list.GET("/", h.getAllList)
			list.GET("/:id", h.getListByID)
			list.PUT("/:id", h.updateList)
			list.DELETE("/:id", h.deleteList)

			items := list.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItem)
				items.GET("/:item_id", h.getItemByID)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}

	return router

}
