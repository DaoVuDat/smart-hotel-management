package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"user-service/internal/domain"
)

type HealthCheckHandler struct{}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func SetupRoutes(e *echo.Echo) {
	h := NewHealthCheckHandler()

	e.GET("/health/live", h.GetLive)
	e.GET("/health/ready", h.GetReady)
}

// GetLive handles the liveness probe endpoint.
// Returns 200 OK if the service is running and can accept requests.
// @Summary Liveness probe endpoint. Purpose: Is the service running (not stuck or crashed?)
func (h *HealthCheckHandler) GetLive(c echo.Context) error {
	return c.JSON(http.StatusOK, domain.HealthStatus{
		Status: "alive",
	})
}

// GetReady handles the readiness probe endpoint.
// Returns 200 OK if the service is ready to handle traffic.
// @Summary Readiness probe endpoint. Purpose: Can the service serve traffic?
func (h *HealthCheckHandler) GetReady(c echo.Context) error {
	//TODO: check DB here
	checkDb := true
	if !checkDb {
		return c.JSON(http.StatusInternalServerError, domain.HealthStatus{
			Status: "db error",
		})
	}

	return c.JSON(http.StatusOK, domain.HealthStatus{
		Status: "ready",
	})
}
