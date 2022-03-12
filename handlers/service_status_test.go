package handlers

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/large-scale-deployment/coordinator/models"
	"github.com/large-scale-deployment/coordinator/services"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	spacesRegexp = regexp.MustCompile(`[\n\t]`)
	statusJSON   = spacesRegexp.ReplaceAllString(`{
        "name":"My Go Service",
        "version":"1.0",
        "group":"Group 11",
        "node_name":"node name 1",
        "host_ip":"192.168.1.1",
        "pod_ip":"192.168.8.8",
        "pod_name":"pod-name-xxx",
        "pod_namespace":"default"
    }`, "")
)

func TestRegisterServiceStatus(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/registry/statuses", strings.NewReader(statusJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	m := &models.Models{DB: db}
	m.AutoMigrate()
	registry := &services.Registry{DB: m.DB}
	h := &ServiceStatusHandler{Registry: registry}

	// Assertions
	if assert.NoError(t, h.Create(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		// assert.Equal(t, statusJSON, rec.Body.String())
	}
}
