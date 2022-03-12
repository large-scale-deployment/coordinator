package handlers

import (
	"net/http"
	"strconv"

	"github.com/large-scale-deployment/coordinator/models"
	"github.com/large-scale-deployment/coordinator/services"

	"github.com/labstack/echo/v4"
)

// PodStatus
type ReqPodStatus struct {
	models.PodStatusData
}

type PodStatusHandler struct {
	Registry *services.Registry
}

// Handler
func (handler *PodStatusHandler) Create(c echo.Context) (err error) {
	ps := new(ReqPodStatus)
	if err = c.Bind(ps); err != nil {
		return err
	}
	psData := &models.PodStatusData{
		Name:              ps.Name,
		PodScheduledAt:    ps.PodScheduledAt,
		InitializedAt:     ps.InitializedAt,
		ContainersReadyAt: ps.ContainersReadyAt,
		ReadyAt:           ps.ReadyAt,
	}
	ID := handler.Registry.ReportPodStatus(psData)
	return c.String(http.StatusCreated, strconv.FormatUint(uint64(ID), 10))
}
