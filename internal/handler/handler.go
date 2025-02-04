package handler

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"seatPlanner/internal/handler/plan"
	"seatPlanner/internal/service"
	"time"
)

type Handler struct {
	planHandler *plan.Handler
}

func New(planServ service.PlannerService) *Handler {
	return &Handler{
		planHandler: plan.NewPlanHandler(planServ),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n")
	}))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "refreshToken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := router.Group("/api")
	{
		mongo := api.Group("/plan")
		{
			mongo.GET("/", h.planHandler.GetAllPlans)
			mongo.GET("/:planId", h.planHandler.GetPlan)
			mongo.POST("/", h.planHandler.CreatePlan)
			mongo.PUT("/:planId", h.planHandler.UpdatePlan)
			mongo.DELETE("/:planId", h.planHandler.DeletePlan)

			mongo.GET("/seat/", h.planHandler.GetAllSeats)
			mongo.GET("/seat/:seatId", h.planHandler.GetSeat)
			mongo.POST("/seat/", h.planHandler.CreateSeat)
			mongo.PUT("/seat/:seatId", h.planHandler.UpdateSeat)
			mongo.DELETE("/seat/:seatId", h.planHandler.DeleteSeat)
		}
	}

	return router
}
