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
	HostIP       string `json:"host_ip" form:"host_ip" query:"host_ip"`
	PodIP        string `json:"pod_ip" form:"pod_ip" query:"pod_ip"`
	PodName      string `json:"pod_name" form:"pod_name" query:"pod_name"`
	PodNamespace string `json:"pod_namespace" form:"pod_namespace" query:"pod_namespace"`
	// The pods (or containers) for every new deployment blongs to a new group
	// The group name may be set by env vars
}

// It has an implicit auto increment primary key `id`.
type PodStatusData struct {
	Name              string    `json:"name" form:"name" query:"name"`
    PodScheduledAt    time.Time `json:"pod_schedualed_at" form:"pod_schedualed_at" query:"pod_schedualed_at"`
    InitializedAt     time.Time `json:"initialized_at" form:"initialized_at" query:"initialized_at"`
	ContainersReadyAt time.Time `json:"containers_ready_at" form:"containers_ready_at" query:"containers_ready_at"`
	ReadyAt           time.Time `json:"ready_at" form:"ready_at" query:"ready_at"`
}

type ServiceStatus struct {
	gorm.Model
	ServiceStatusData
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PodStatus struct {
	gorm.Model
	PodStatusData
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Models struct {
	DB *gorm.DB
}

func (models *Models) AutoMigrate() {
	models.DB.AutoMigrate(&ServiceStatus{})
	models.DB.AutoMigrate(&PodStatus{})
}
