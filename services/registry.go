package services

import (
	"lss/coordinator/models"

	"gorm.io/gorm"
)

type Registry struct {
	DB *gorm.DB
}

// ReverseRunes returns its argument string reversed rune-wise left to right.
func (reg *Registry) ReportServiceStatus(ssData *models.ServiceStatusData) uint {
	serviceStatusObject := &models.ServiceStatus{
		ServiceStatusData: *ssData,
	}
	reg.DB.Create(serviceStatusObject)
	return serviceStatusObject.ID
}
