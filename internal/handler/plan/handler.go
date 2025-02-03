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

func (h *Handler) GetAllPlans(c *gin.Context) {
	plans, statusCode, err := h.planService.GetAllPlans()
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

	plan, statusCode, err := h.planService.GetPlan(id)
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

func (h *Handler) GetAllSeats(c *gin.Context) {
	plans, statusCode, err := h.planService.GetAllSeats()
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

	plan, statusCode, err := h.planService.GetSeat(id)
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

	statusCode, err := h.planService.CreatePlan(input)

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

	statusCode, err := h.planService.CreateSeat(input)

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

	statusCode, err := h.planService.UpdatePlan(input, id)

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

	statusCode, err := h.planService.UpdateSeat(input, id)

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

	statusCode, err := h.planService.DeletePlan(id)

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

	statusCode, err := h.planService.DeleteSeat(id)

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
