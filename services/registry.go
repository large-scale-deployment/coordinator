package services

import (
	"lss/coordinator/models"

	"gorm.io/gorm"
)

type Registry struct {
	DB *gorm.DB
}

// ReverseRunes returns its argument string reversed rune-wise left to right.
func (reg *Registry) ReportServiceStatus(serviceStatus *models.ServiceStatus) uint {
	reg.DB.Create(serviceStatus)
	return serviceStatus.ID
}
