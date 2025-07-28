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

// CreatePlan godoc
// @Summary Create a shopping plan
// @Description Create a new shopping plan for a user
// @Tags Shopping Plans
// @Accept json
// @Produce json
// @Param user_id header string true "User ID"
// @Param plan body models.ShoppingPlan true "Shopping Plan"
// @Success 201 {object} models.ShoppingPlan
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /shopping-plans [post]
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

// GetPlanByID godoc
// @Summary Get shopping plan by ID
// @Description Retrieve a shopping plan using its ID
// @Tags Shopping Plans
// @Produce json
// @Param id path string true "Plan ID"
// @Success 200 {object} models.ShoppingPlan
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /shopping-plans/{id} [get]
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

// GetUserPlans godoc
// @Summary Get all plans of a user
// @Description Retrieve all shopping plans associated with a user
// @Tags Shopping Plans
// @Produce json
// @Param id path string true "User ID"
// @Param user_id header string true "Requesting User ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 403 {object} gin.H
// @Router /shopping-plans/user/{id} [get]
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

// GetActivePlans godoc
// @Summary Get active plans
// @Description Get all active shopping plans for a user
// @Tags Shopping Plans
// @Produce json
// @Param id path string true "User ID"
// @Param user_id header string true "Requesting User ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 403 {object} gin.H
// @Router /shopping-plans/user/{id}/active [get]
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

// GetCompletedPlans godoc
// @Summary Get completed plans
// @Description Get all completed shopping plans for a user
// @Tags Shopping Plans
// @Produce json
// @Param id path string true "User ID"
// @Param user_id header string true "Requesting User ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 403 {object} gin.H
// @Router /shopping-plans/user/{id}/completed [get]
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

// UpdatePlan godoc
// @Summary Update a shopping plan
// @Description Update a shopping plan details by ID
// @Tags Shopping Plans
// @Accept json
// @Produce json
// @Param user_id header string true "User ID"
// @Param id path string true "Plan ID"
// @Param plan body models.ShoppingPlan true "Shopping Plan"
// @Success 200 {object} models.ShoppingPlan
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /shopping-plans/{id} [put]
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

// DeletePlan godoc
// @Summary Delete a shopping plan
// @Description Delete a shopping plan by ID
// @Tags Shopping Plans
// @Produce json
// @Param user_id header string true "User ID"
// @Param id path string true "Plan ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /shopping-plans/{id} [delete]
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

// MarkShopVisited godoc
// @Summary Mark a shop as visited
// @Description Mark a shop as visited in a shopping plan
// @Tags Shopping Plans
// @Produce json
// @Param user_id header string true "User ID"
// @Param id path string true "Plan ID"
// @Param shopId path string true "Shop ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 403 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /shopping-plans/{id}/visit/{shopId} [put]
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

// AddShopToPlan godoc
// @Summary Add a shop to a plan
// @Description Add a shop to an existing shopping plan
// @Tags Shopping Plans
// @Accept json
// @Produce json
// @Param user_id header string true "User ID"
// @Param id path string true "Plan ID"
// @Param shop body map[string]string true "Shop ID payload"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 403 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /shopping-plans/{id}/shops [post]
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

// RemoveShopFromPlan godoc
// @Summary Remove a shop from a plan
// @Description Remove a shop from a shopping plan
// @Tags Shopping Plans
// @Produce json
// @Param user_id header string true "User ID"
// @Param id path string true "Plan ID"
// @Param shopId path string true "Shop ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 403 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /shopping-plans/{id}/shops/{shopId} [delete]
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

// GetPlanProgress godoc
// @Summary Get plan progress
// @Description Get progress status of a shopping plan
// @Tags Shopping Plans
// @Produce json
// @Param id path string true "Plan ID"
// @Success 200 {object} interface{}
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /shopping-plans/{id}/progress [get]
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

// GetUserStats godoc
// @Summary Get user statistics
// @Description Get aggregated shopping plan stats for a user
// @Tags Shopping Plans
// @Produce json
// @Param id path string true "User ID"
// @Param user_id header string true "Requesting User ID"
// @Success 200 {object} interface{}
// @Failure 400 {object} gin.H
// @Failure 403 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /shopping-plans/user/{id}/stats [get]
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
