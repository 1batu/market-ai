package handlers

import (
	"context"

	"github.com/1batu/market-ai/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MetricsHandler struct {
	db *pgxpool.Pool
}

func NewMetricsHandler(db *pgxpool.Pool) *MetricsHandler {
	return &MetricsHandler{db: db}
}

type DataSourceMetrics struct {
	SourceType        string `json:"source_type"`
	SourceName        string `json:"source_name"`
	IsActive          bool   `json:"is_active"`
	TotalFetches      int    `json:"total_fetches"`
	SuccessCount      int    `json:"success_count"`
	ErrorCount        int    `json:"error_count"`
	AvgResponseTimeMs int    `json:"avg_response_time_ms"`
	Status            string `json:"status"`
	LastError         string `json:"last_error,omitempty"`
	LastFetchAt       string `json:"last_fetch_at,omitempty"`
	UpdatedAt         string `json:"updated_at,omitempty"`
}

// Get /api/v1/metrics
func (h *MetricsHandler) Get(c *fiber.Ctx) error {
	rows, err := h.db.Query(context.Background(), `
		SELECT source_type, source_name, COALESCE(is_active,true),
		       COALESCE(total_fetches,0), COALESCE(success_count,0), COALESCE(error_count,0),
		       COALESCE(avg_response_time_ms,0), COALESCE(status,'active'),
		       COALESCE(last_error,''), COALESCE(last_fetch_at::text,''), COALESCE(updated_at::text,'')
		FROM data_sources
		ORDER BY source_type, source_name
	`)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{Success: false, Message: "metrics query error"})
	}
	defer rows.Close()
	var metrics []DataSourceMetrics
	for rows.Next() {
		var m DataSourceMetrics
		if err := rows.Scan(&m.SourceType, &m.SourceName, &m.IsActive, &m.TotalFetches, &m.SuccessCount, &m.ErrorCount, &m.AvgResponseTimeMs, &m.Status, &m.LastError, &m.LastFetchAt, &m.UpdatedAt); err != nil {
			continue
		}
		metrics = append(metrics, m)
	}
	return c.JSON(models.Response{Success: true, Data: fiber.Map{
		"data_sources": metrics,
	}})
}
