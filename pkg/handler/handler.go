package handler

import (
	"AvitoTask/pkg/service"
	"github.com/gin-gonic/gin"
	//"github.com/handlest/avito/pkg/service"
	//
	//"github.com/swaggo/gin-swagger/swaggerFiles"
	//
	//_ "github.com/handlest/avito/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRoutes Инициализирует маршруты и перенаправляет на обработчики в зависимости от пути
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		segments := api.Group("/segments")
		{
			segments.POST("/", h.createSegment)
			segments.DELETE("/:name", h.deleteSegment)
			segments.GET("/", h.getAllSegments)
		}
		users := api.Group("/users")
		{
			users.POST("/", h.createUser)
			users.POST("/:ttl", h.createUserTTL)
			users.DELETE("/:id", h.deleteUser)

			userInfo := api.Group("/userInfo")
			{
				userInfo.GET("/:dateRange", h.getUserInfo)
			}
		}
	}
	return router
}
