package registry

import (
	"time"

	"gorm.io/gorm"
)

type Registry struct {
	DB *gorm.DB
}

// It has an implicit auto increment primary key `id`.
type ServiceStatus struct {
	gorm.Model
	Name    string
	Ip      string
	Version string
	// The pods (or containers) for every new deployment blongs to a new cluster
	// The cluster name may be set by env vars
	Cluster   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (reg *Registry) AutoMigrate() {
	reg.DB.AutoMigrate(&ServiceStatus{})
}

// ReverseRunes returns its argument string reversed rune-wise left to right.
func (reg *Registry) ReportServiceStatus(name, ip, version, cluster string) {
	ss := &ServiceStatus{Name: name, Ip: ip, Version: version, Cluster: cluster}
	reg.DB.Create(ss)
}
