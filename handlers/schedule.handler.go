package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/radifan9/platform-tiket-bioskop/models"
	"github.com/radifan9/platform-tiket-bioskop/repositories"
	"github.com/radifan9/platform-tiket-bioskop/utils"
)

type ScheduleHandler struct {
	sr *repositories.ScheduleRepository
}

func NewScheduleHandler(sr *repositories.ScheduleRepository) *ScheduleHandler {
	return &ScheduleHandler{sr: sr}
}

func (h *ScheduleHandler) CreateSchedule(ctx *gin.Context) {
	var req models.Schedule

	if err := ctx.ShouldBind(&req); err != nil {
		utils.HandleError(ctx, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	schedule, err := h.sr.CreateSchedule(context.Background(), req)
	if err != nil {
		utils.HandleError(ctx, http.StatusInternalServerError, "failed to create schedule", err.Error())
		return
	}

	utils.HandleResponse(ctx, http.StatusCreated, schedule)
}
