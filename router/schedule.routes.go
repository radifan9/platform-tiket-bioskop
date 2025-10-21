package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/radifan9/platform-tiket-bioskop/handlers"
	"github.com/radifan9/platform-tiket-bioskop/middlewares"
	"github.com/radifan9/platform-tiket-bioskop/repositories"

	"github.com/redis/go-redis/v9"
)

func RegisterSchedulesRoutes(v1 *gin.RouterGroup, db *pgxpool.Pool, rdb *redis.Client) {
	scheduleRepo := repositories.NewScheduleRepository(db)
	scheduleHandler := handlers.NewScheduleHandler(scheduleRepo)
	VerifyTokenWithBlacklist := middlewares.VerifyTokenWithBlacklist(rdb)

	schedules := v1.Group("/schedules")
	schedules.Use(VerifyTokenWithBlacklist, middlewares.Access("admin")).POST("", scheduleHandler.CreateSchedule)

}
