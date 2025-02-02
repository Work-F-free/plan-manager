package plan

import "seatPlanner/internal/service"

type Handler struct {
	planService service.PlannerService
}

func NewPlanHandler(plannerService service.PlannerService) *Handler {
	return &Handler{
		planService: plannerService,
	}
}
