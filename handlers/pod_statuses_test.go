package handlers

import (
	"net/http"
	"net/http/httptest"
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
	podStatusJSON = `{
        "name":"Pod Status 1",
        "pod_schedualed_at":"2022-03-13T08:07:47Z",
        "initialized_at":"2022-03-14T09:07:47Z",
        "containers_ready_at":"2022-03-14T10:07:47Z",
        "ready_at":"2022-03-14T11:07:47Z"
    }`
)

func TestRegisterPodStatus(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/registry/pod_statuses", strings.NewReader(podStatusJSON))
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
	h := &PodStatusHandler{Registry: registry}

	// Assertions
	if assert.NoError(t, h.Create(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		// assert.Equal(t, statusJSON, rec.Body.String())
	}
}
