package handlers

import (
	"loyaltea-server/internal/models"
	"loyaltea-server/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShoppingPlanHandler struct {
	planService *services.ShoppingPlanService
}

func NewShoppingPlanHandler(planService *services.ShoppingPlanService) *ShoppingPlanHandler {
	return &ShoppingPlanHandler{planService: planService}
}

// POST /shopping-plans
func (h *ShoppingPlanHandler) CreatePlan(ctx *gin.Context) {
	userID := ctx.GetHeader("user_id")
	if userID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var plan models.ShoppingPlan
	if err := ctx.ShouldBindJSON(&plan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plan.UserID = userID

	if err := h.planService.CreatePlan(ctx, &plan); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, plan)
}

// GET /shopping-plans/:id
func (h *ShoppingPlanHandler) GetPlanByID(ctx *gin.Context) {
	planID := ctx.Param("id")
	if planID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Plan ID is required"})
		return
	}

	plan, err := h.planService.GetPlanByID(ctx, planID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Plan not found"})
		return
	}

	ctx.JSON(http.StatusOK, plan)
}

// GET /shopping-plans/user/:id
func (h *ShoppingPlanHandler) GetUserPlans(ctx *gin.Context) {
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	// Check if requesting user matches the user_id in header (basic auth check)
	requestingUserID := ctx.GetHeader("user_id")
	if requestingUserID != userID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	plans, err := h.planService.GetUserPlans(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"plans": plans,
		"count": len(plans),
	})
}

// GET /shopping-plans/user/:id/active
func (h *ShoppingPlanHandler) GetActivePlans(ctx *gin.Context) {
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	// Check if requesting user matches the user_id in header
	requestingUserID := ctx.GetHeader("user_id")
	if requestingUserID != userID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	plans, err := h.planService.GetActivePlans(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"plans": plans,
		"count": len(plans),
	})
}

// GET /shopping-plans/user/:id/completed
func (h *ShoppingPlanHandler) GetCompletedPlans(ctx *gin.Context) {
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	// Check if requesting user matches the user_id in header
	requestingUserID := ctx.GetHeader("user_id")
	if requestingUserID != userID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	plans, err := h.planService.GetCompletedPlans(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"plans": plans,
		"count": len(plans),
	})
}

// PUT /shopping-plans/:id
func (h *ShoppingPlanHandler) UpdatePlan(ctx *gin.Context) {
	userID := ctx.GetHeader("user_id")
	if userID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	planID := ctx.Param("id")
	if planID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Plan ID is required"})
		return
	}

	var plan models.ShoppingPlan
	if err := ctx.ShouldBindJSON(&plan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plan.ID = planID
	plan.UserID = userID

	if err := h.planService.UpdatePlan(ctx, &plan); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, plan)
}

// DELETE /shopping-plans/:id
func (h *ShoppingPlanHandler) DeletePlan(ctx *gin.Context) {
	userID := ctx.GetHeader("user_id")
	if userID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	planID := ctx.Param("id")
	if planID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Plan ID is required"})
		return
	}

	if err := h.planService.DeletePlan(ctx, planID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Plan deleted successfully"})
}

// PUT /shopping-plans/:id/visit/:shopId
func (h *ShoppingPlanHandler) MarkShopVisited(ctx *gin.Context) {
	userID := ctx.GetHeader("user_id")
	if userID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	planID := ctx.Param("id")
	shopID := ctx.Param("shopId")

	if planID == "" || shopID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Plan ID and Shop ID are required"})
		return
	}

	if err := h.planService.MarkShopVisited(ctx, planID, userID, shopID); err != nil {
		switch err {
		case services.ErrUnauthorizedAccess:
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		case services.ErrShopNotInPlan, services.ErrShopAlreadyVisited:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Shop marked as visited and points awarded"})
}

// POST /shopping-plans/:id/shops
func (h *ShoppingPlanHandler) AddShopToPlan(ctx *gin.Context) {
	userID := ctx.GetHeader("user_id")
	if userID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	planID := ctx.Param("id")
	if planID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Plan ID is required"})
		return
	}

	var request struct {
		ShopID string `json:"shop_id"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.ShopID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Shop ID is required"})
		return
	}

	if err := h.planService.AddShopToPlan(ctx, planID, userID, request.ShopID); err != nil {
		switch err {
		case services.ErrUnauthorizedAccess:
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		case services.ErrShopAlreadyInPlan, services.ErrPlanAlreadyCompleted, services.ErrShopNotActive:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case services.ErrShopNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Shop added to plan successfully"})
}

// DELETE /shopping-plans/:id/shops/:shopId
func (h *ShoppingPlanHandler) RemoveShopFromPlan(ctx *gin.Context) {
	userID := ctx.GetHeader("user_id")
	if userID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	planID := ctx.Param("id")
	shopID := ctx.Param("shopId")

	if planID == "" || shopID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Plan ID and Shop ID are required"})
		return
	}

	if err := h.planService.RemoveShopFromPlan(ctx, planID, userID, shopID); err != nil {
		switch err {
		case services.ErrUnauthorizedAccess:
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		case services.ErrShopNotInPlan, services.ErrPlanAlreadyCompleted:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Shop removed from plan successfully"})
}

// GET /shopping-plans/:id/progress
func (h *ShoppingPlanHandler) GetPlanProgress(ctx *gin.Context) {
	planID := ctx.Param("id")
	if planID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Plan ID is required"})
		return
	}

	progress, err := h.planService.GetPlanProgress(ctx, planID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, progress)
}

// GET /shopping-plans/user/:id/stats
func (h *ShoppingPlanHandler) GetUserStats(ctx *gin.Context) {
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	// Check if requesting user matches the user_id in header
	requestingUserID := ctx.GetHeader("user_id")
	if requestingUserID != userID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	stats, err := h.planService.GetUserStats(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, stats)
}
