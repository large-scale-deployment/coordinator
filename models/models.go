package models

import (
	"time"

	"gorm.io/gorm"
)

// It has an implicit auto increment primary key `id`.
type ServiceStatus struct {
	gorm.Model
	Name         string
	Version      string
	Group        string
	NodeName     string
	PodIp        string
	PodName      string
	PodNamespace string
	// The pods (or containers) for every new deployment blongs to a new cluster
	// The cluster name may be set by env vars
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Models struct {
	DB *gorm.DB
}

func (models *Models) AutoMigrate() {
	models.DB.AutoMigrate(&ServiceStatus{})
}
