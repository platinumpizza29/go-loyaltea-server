package services

import "errors"

// User service errors
var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserIDRequired    = errors.New("user ID is required")
)

// Shop service errors
var (
	ErrShopNotFound        = errors.New("shop not found")
	ErrShopNameRequired    = errors.New("shop name is required")
	ErrShopAddressRequired = errors.New("shop address is required")
	ErrInvalidLocation     = errors.New("invalid location coordinates")
	ErrInvalidLongitude    = errors.New("longitude must be between -180 and 180")
	ErrInvalidLatitude     = errors.New("latitude must be between -90 and 90")
	ErrShopNotActive       = errors.New("shop is not active")
)

// Shopping plan service errors
var (
	ErrPlanNotFound         = errors.New("shopping plan not found")
	ErrPlanNameRequired     = errors.New("plan name is required")
	ErrPlanAlreadyCompleted = errors.New("plan is already completed")
	ErrShopNotInPlan        = errors.New("shop is not in the plan")
	ErrShopAlreadyInPlan    = errors.New("shop is already in the plan")
	ErrShopAlreadyVisited   = errors.New("shop has already been visited")
	ErrUnauthorizedAccess   = errors.New("unauthorized access to plan")
)

// General service errors
var (
	ErrInternalServer = errors.New("internal server error")
	ErrBadRequest     = errors.New("bad request")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
)
