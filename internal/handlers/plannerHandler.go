package handlers

import (
	"loyaltea-server/internal/models"
	"loyaltea-server/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlannerHandler struct {
	plannerService *services.PlannerService
}

func NewPlannerHandler(plannerService *services.PlannerService) *PlannerHandler {
	return &PlannerHandler{plannerService: plannerService}
}

func (h *PlannerHandler) CreateStop(ctx *gin.Context) {
	userID := ctx.GetHeader("user_id")
	if userID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	stop := &models.Stop{
		ID:     primitive.NewObjectID().Hex(),
		UserID: userID,
	}

	if err := ctx.ShouldBindJSON(stop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.plannerService.CreateStop(ctx, stop); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, stop)
}

// get all geo points
func (h *PlannerHandler) GetAllGeoPoints(ctx *gin.Context) {
	geoPoints, err := h.plannerService.GetAllGeoPoints(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, geoPoints)
}

func (h *PlannerHandler) GetStopsByUserID(ctx *gin.Context) {
	userID := ctx.GetHeader("user_id")
	if userID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	stops, err := h.plannerService.GetStopsByUserID(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, stops)
}

func (h *PlannerHandler) UpdateStop(ctx *gin.Context) {
	userID := ctx.GetHeader("user_id")
	if userID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	stopID := ctx.Param("id")
	if stopID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Stop ID is required"})
		return
	}

	stop := &models.Stop{
		ID: stopID,
	}

	if err := ctx.ShouldBindJSON(stop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (h *PlannerHandler) DeleteStop(ctx *gin.Context) {
	userID := ctx.GetHeader("user_id")
	if userID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	stopID := ctx.Param("id")
	if stopID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Stop ID is required"})
		return
	}

	stop := &models.Stop{
		ID: stopID,
	}

	if err := h.plannerService.DeleteStop(ctx, stop); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
