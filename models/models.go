package models

import (
	"time"

	"gorm.io/gorm"
)

// It has an implicit auto increment primary key `id`.
type ServiceStatusData struct {
	Name         string `json:"name" form:"name" query:"name"`          // Service name
	Version      string `json:"version" form:"version" query:"version"` // Service version
	Group        string `json:"group" form:"group" query:"group"`       // Deployment group
	NodeName     string `json:"node_name" form:"node_name" query:"node_name"`
	PodIP        string `json:"pod_ip" form:"pod_ip" query:"pod_ip"`
	PodName      string `json:"pod_name" form:"pod_name" query:"pod_name"`
	PodNamespace string `json:"pod_namespace" form:"pod_namespace" query:"pod_namespace"`
	// The pods (or containers) for every new deployment blongs to a new group
	// The group name may be set by env vars
}
type ServiceStatus struct {
	gorm.Model
	ServiceStatusData
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Models struct {
	DB *gorm.DB
}

func (models *Models) AutoMigrate() {
	models.DB.AutoMigrate(&ServiceStatus{})
}
