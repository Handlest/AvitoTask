package handler

import (
	"AvitoTask/pkg/service"
	"github.com/gin-gonic/gin"
	_ "github.com/handlest/avitotask/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		segments := api.Group("/segments")
		{
			segments.POST("/", h.createSegment)   //done
			segments.DELETE("/", h.deleteSegment) //done

			getUserSegment := api.Group("/getUserSegments")
			{
				getUserSegment.POST("/", h.getUserSegments) //done
			}
		}
		users := api.Group("/users")
		{
			users.POST("/", h.createUser)   //done
			users.DELETE("/", h.deleteUser) //done

			userInfo := api.Group("/userInfo") //done
			{
				userInfo.POST("/", h.getUserInfo)
			}
		}
	}
	return router
}
