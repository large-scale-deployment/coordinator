package handlers

import (
	"large-scale-deployment/coordinator/models"
	"large-scale-deployment/coordinator/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// ServiceStatus
type ReqServiceStatus struct {
	models.ServiceStatusData
}

type ServiceStatusHandler struct {
	Registry *services.Registry
}

// Handler
func (handler *ServiceStatusHandler) Create(c echo.Context) (err error) {
	ss := new(ReqServiceStatus)
	if err = c.Bind(ss); err != nil {
		return err
	}
	ssData := &models.ServiceStatusData{
		Name:         ss.Name,
		Version:      ss.Version,
		Group:        ss.Group,
		NodeName:     ss.NodeName,
		PodIP:        ss.PodIP,
		PodName:      ss.PodName,
		PodNamespace: ss.PodNamespace,
	}
	ID := handler.Registry.ReportServiceStatus(ssData)
	return c.String(http.StatusCreated, strconv.FormatUint(uint64(ID), 10))
}
