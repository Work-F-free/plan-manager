package plan

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"seatPlanner/internal/common/dto"
	"seatPlanner/internal/common/errors"
	"seatPlanner/internal/common/responses"
	"seatPlanner/internal/service"
)

type Handler struct {
	planService service.PlannerService
}

func NewPlanHandler(plannerService service.PlannerService) *Handler {
	return &Handler{
		planService: plannerService,
	}
}

// GetAllPlans @Summary Get all coworking plans
// @Description Get data about all coworkings plans in applicationi
// @ID get-plans
// @Tags Plans
// @Accept  json
// @Produce  json
// @Success 200 {object} []dto.Plan
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Failure default {object} errors.ErrorResponse
// @Router /api/plan [get]
func (h *Handler) GetAllPlans(c *gin.Context) {
	plans, statusCode, err := h.planService.GetAllPlans(c)
	if err != nil {
		c.JSON(statusCode, errors.ErrorResponse{
			Error:   "error while getting plans",
			Status:  statusCode,
			Details: err.Error(),
		})
		return
	}
	c.JSON(statusCode, plans)
}

// GetPlan @Summary Get coworking plan by id
// @Description Get data about one coworking
// @ID get-plan
// @Tags Plans
// @Accept  json
// @Produce  json
// @Param planId path string true "Plan Id"
// @Success 200 {object} dto.Plan
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Failure default {object} errors.ErrorResponse
// @Router /api/plan/{planId} [get]
func (h *Handler) GetPlan(c *gin.Context) {
	id, err := uuid.Parse(c.Param("planId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{
			Error:   "system can't parse this query to UUID",
			Status:  http.StatusBadRequest,
			Details: err.Error(),
		})
		return
	}

	plan, statusCode, err := h.planService.GetPlan(c, id)
	if err != nil {
		c.JSON(statusCode, errors.ErrorResponse{
			Error:   "error while getting plan",
			Status:  statusCode,
			Details: err.Error(),
		})
		return
	}

	c.JSON(statusCode, plan)
}

// GetAllSeats @Summary Get all seats plans
// @Description Get all seats from coworking
// @ID get-seats
// @Tags Seats
// @Accept  json
// @Produce  json
// @Param planId path string true "Plan Id"
// @Success 200 {object} []dto.Seat
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Failure default {object} errors.ErrorResponse
// @Router /api/plan/seat/{planId} [get]
func (h *Handler) GetAllSeats(c *gin.Context) {
	plans, statusCode, err := h.planService.GetAllSeats(c)
	if err != nil {
		c.JSON(statusCode, errors.ErrorResponse{
			Error:   "error while getting seats",
			Status:  statusCode,
			Details: err.Error(),
		})
		return
	}
	c.JSON(statusCode, plans)
}

// GetSeat @Summary Get seat by id
// @Description Get data about one coworking
// @ID get-seat
// @Tags Seats
// @Accept  json
// @Produce  json
// @Param seatId path string true "Seat Id"
// @Success 200 {object} dto.Seat
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Failure default {object} errors.ErrorResponse
// @Router /api/plan/seat/{seatId} [get]
func (h *Handler) GetSeat(c *gin.Context) {
	id, err := uuid.Parse(c.Param("seatId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{
			Error:   "system can't parse this query to UUID",
			Status:  http.StatusBadRequest,
			Details: err.Error(),
		})
		return
	}

	plan, statusCode, err := h.planService.GetSeat(c, id)
	if err != nil {
		c.JSON(statusCode, errors.ErrorResponse{
			Error:   "error while getting plan",
			Status:  statusCode,
			Details: err.Error(),
		})
		return
	}

	c.JSON(statusCode, plan)
}

// CreatePlan @Summary Create coworking plan
// @Description Creates plan
// @ID create-plan
// @Tags Plans
// @Accept  json
// @Produce  json
// @Param input body dto.Plan true "Coworking plan"
// @Success 201 {object} responses.SuccessResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Failure default {object} errors.ErrorResponse
// @Router /api/plan/ [post]
func (h *Handler) CreatePlan(c *gin.Context) {
	var input dto.Plan

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{
			Error:   "can't parse JSON",
			Status:  http.StatusBadRequest,
			Details: err.Error(),
		})

		return
	}

	statusCode, err := h.planService.CreatePlan(c, &input)

	if err != nil {
		c.JSON(statusCode, errors.ErrorResponse{
			Error:   "error while creating plan",
			Status:  statusCode,
			Details: err.Error(),
		})

		return
	}

	c.JSON(statusCode, responses.SuccessResponse{
		Status:  statusCode,
		Message: "create plan successfully",
	})
}

// CreateSeat @Summary Create seat to coworking plan
// @Description Creates seat
// @ID create-seat
// @Tags Seats
// @Accept  json
// @Produce  json
// @Param planId path string true "Plan Id"
// @Param input body dto.Seat true "Seat in coworking plan"
// @Success 201 {object} responses.SuccessResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Failure default {object} errors.ErrorResponse
// @Router /api/plan/seat/{planId} [post]
func (h *Handler) CreateSeat(c *gin.Context) {
	var input dto.Seat

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{
			Error:   "can't parse JSON",
			Status:  http.StatusBadRequest,
			Details: err.Error(),
		})

		return
	}

	statusCode, err := h.planService.CreateSeat(c, &input)

	if err != nil {
		c.JSON(statusCode, errors.ErrorResponse{
			Error:   "error while creating seat",
			Status:  statusCode,
			Details: err.Error(),
		})

		return
	}

	c.JSON(statusCode, responses.SuccessResponse{
		Status:  statusCode,
		Message: "create seat successfully",
	})

}

// UpdatePlan @Summary Update plan
// @Description Update plan from coworking plans
// @ID update-plan
// @Tags Plans
// @Accept  json
// @Produce  json
// @Param planId path string true "Plan Id"
// @Param input body dto.Plan true "Coworking plan"
// @Success 204 {object} responses.SuccessResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Failure default {object} errors.ErrorResponse
// @Router /api/plan/{planId} [put]
func (h *Handler) UpdatePlan(c *gin.Context) {
	id, err := uuid.Parse(c.Param("planId"))
	var input dto.Plan

	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{
			Error:   "system can't parse this query to UUID",
			Status:  http.StatusBadRequest,
			Details: err.Error(),
		})
		return
	}

	if err = c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{
			Error:   "can't parse JSON",
			Status:  http.StatusBadRequest,
			Details: err.Error(),
		})

		return
	}

	statusCode, err := h.planService.UpdatePlan(c, &input, id)

	if err != nil {
		c.JSON(statusCode, errors.ErrorResponse{
			Error:   "error while updating plan",
			Status:  statusCode,
			Details: err.Error(),
		})

		return
	}

	c.JSON(statusCode, responses.SuccessResponse{
		Status:  statusCode,
		Message: "update plan successfully",
	})
}

// UpdateSeat @Summary Update seat from coworking plan
// @Description Updates seat
// @ID update-seat
// @Tags Seats
// @Accept  json
// @Produce  json
// @Param planId path string true "Plan Id"
// @Param input body dto.Seat true "Seat in coworking plan"
// @Success 204 {object} responses.SuccessResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Failure default {object} errors.ErrorResponse
// @Router /api/plan/seat/{planId} [put]
func (h *Handler) UpdateSeat(c *gin.Context) {
	id, err := uuid.Parse(c.Param("seatId"))
	var input dto.Seat

	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{
			Error:   "system can't parse this query to UUID",
			Status:  http.StatusBadRequest,
			Details: err.Error(),
		})
		return
	}

	if err = c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{
			Error:   "can't parse JSON",
			Status:  http.StatusBadRequest,
			Details: err.Error(),
		})

		return
	}

	statusCode, err := h.planService.UpdateSeat(c, &input, id)

	if err != nil {
		c.JSON(statusCode, errors.ErrorResponse{
			Error:   "error while updating seat",
			Status:  statusCode,
			Details: err.Error(),
		})

		return
	}

	c.JSON(statusCode, responses.SuccessResponse{
		Status:  statusCode,
		Message: "update seat successfully",
	})
}

// DeletePlan @Summary Delete plan
// @Description Deletes coworking plan
// @ID delete-plan
// @Tags Plans
// @Accept  json
// @Produce  json
// @Param planId path string true "Plan Id"
// @Success 204 {object} responses.SuccessResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Failure default {object} errors.ErrorResponse
// @Router /api/plan/{planId} [delete]
func (h *Handler) DeletePlan(c *gin.Context) {
	id, err := uuid.Parse(c.Param("planId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{
			Error:   "system can't parse this query to UUID",
			Status:  http.StatusBadRequest,
			Details: err.Error(),
		})
		return
	}

	statusCode, err := h.planService.DeletePlan(c, id)

	if err != nil {
		c.JSON(statusCode, errors.ErrorResponse{
			Error:   "error while deleting plan",
			Status:  statusCode,
			Details: err.Error(),
		})

		return
	}

	c.JSON(statusCode, responses.SuccessResponse{
		Status:  statusCode,
		Message: "delete plan successfully",
	})
}

// DeleteSeat @Summary Delete seat from coworking plan
// @Description Deletes seat
// @ID delete-seat
// @Tags Seats
// @Accept  json
// @Produce  json
// @Param seatId path string true "Seat Id"
// @Success 204 {object} responses.SuccessResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Failure default {object} errors.ErrorResponse
// @Router /api/plan/seat/{seatId} [delete]
func (h *Handler) DeleteSeat(c *gin.Context) {
	id, err := uuid.Parse(c.Param("seatId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{
			Error:   "system can't parse this query to UUID",
			Status:  http.StatusBadRequest,
			Details: err.Error(),
		})
		return
	}

	statusCode, err := h.planService.DeleteSeat(c, id)

	if err != nil {
		c.JSON(statusCode, errors.ErrorResponse{
			Error:   "error while deleting plan",
			Status:  statusCode,
			Details: err.Error(),
		})

		return
	}

	c.JSON(statusCode, responses.SuccessResponse{
		Status:  statusCode,
		Message: "delete seat successfully",
	})
}
