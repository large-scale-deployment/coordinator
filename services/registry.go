package services

import (
	"github.com/large-scale-deployment/coordinator/models"

	"gorm.io/gorm"
)

type Registry struct {
	DB *gorm.DB
}

func (reg *Registry) ReportServiceStatus(ssData *models.ServiceStatusData) uint {
	serviceStatusObject := &models.ServiceStatus{
		ServiceStatusData: *ssData,
	}
	reg.DB.Create(serviceStatusObject)
	return serviceStatusObject.ID
}

func (reg *Registry) ReportPodStatus(psData *models.PodStatusData) uint {
    podStatusObject := &models.PodStatus{
        PodStatusData: *psData,
    }
    reg.DB.Create(podStatusObject)
    return podStatusObject.ID
}
