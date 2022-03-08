package main

import (
	"lss/coordinator/registry"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ServiceStatus
type ReqServiceStatus struct {
	Name    string `json:"name" form:"name" query:"name"`
	Ip      string `json:"ip" form:"ip" query:"ip"`
	Version string `json:"version" form:"version" query:"version"`
	Cluster string `json:"cluster" form:"cluster" query:"cluster"`
}

type RequestHandler struct {
	Registry *registry.Registry
}

// Handler
func (handler *RequestHandler) registerServiceStatus(c echo.Context) (err error) {
	ss := new(ReqServiceStatus)
	if err = c.Bind(ss); err != nil {
		return err
	}
	handler.Registry.ReportServiceStatus(ss.Name, ss.Ip, ss.Version, ss.Cluster)
	return c.String(http.StatusOK, ss.Name)
}

func startService(handler *RequestHandler) {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes
	e.POST("/registry/statuses", handler.registerServiceStatus)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	reg := &registry.Registry{DB: db}
	// Migrate the schema
	reg.AutoMigrate()

	handler := &RequestHandler{Registry: reg}
	startService(handler)
}
