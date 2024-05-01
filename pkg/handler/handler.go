package handler

import (
	_ "github.com/TheKruasan/CRM-server/docs"
	"github.com/TheKruasan/CRM-server/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run("localhost:8000")
	user := router.Group("/order")
	{
		user.POST("/add", h.postOrder)
		getUser := user.Group("/get")
		{
			getUser.GET("/all", h.getAllOrderById)
			getUser.GET("/:id", h.getOrderById)
		}

	}

	return nil
}
func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}
