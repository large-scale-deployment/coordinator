package handlers

import (
	"lss/coordinator/models"
	"lss/coordinator/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ServiceStatus
type ReqServiceStatus struct {
	Name         string `json:"name" form:"name" query:"name"`          // Service name
	Version      string `json:"version" form:"version" query:"version"` // Service version
	Group        string `json:"group" form:"group" query:"group"`       // Deployment group
	NodeName     string `json:"node_name" form:"node_name" query:"node_name"`
	PodIp        string `json:"pod_id" form:"pod_id" query:"pod_id"`
	PodName      string `json:"pod_name" form:"pod_name" query:"pod_name"`
	PodNamespace string `json:"pod_namespace" form:"pod_namespace" query:"pod_namespace"`
}

type RequestHandler struct {
	Registry *services.Registry
}

// Handler
func (handler *RequestHandler) RegisterServiceStatus(c echo.Context) (err error) {
	ss := new(ReqServiceStatus)
	if err = c.Bind(ss); err != nil {
		return err
	}
	serviceStatus := &models.ServiceStatus{
		Name:         ss.Name,
		Version:      ss.Version,
		Group:        ss.Group,
		NodeName:     ss.NodeName,
		PodIp:        ss.PodIp,
		PodName:      ss.PodName,
		PodNamespace: ss.PodNamespace,
	}
	handler.Registry.ReportServiceStatus(serviceStatus)
	return c.String(http.StatusOK, ss.Name)
}
