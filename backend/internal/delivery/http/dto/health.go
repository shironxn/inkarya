package dto

import "time"

type HealthStatus struct {
	Status      string    `json:"status"`
	Environment string    `json:"environment"`
	Version     string    `json:"version"`
	Timestamp   time.Time `json:"timestamp"`
	Database    string    `json:"database"`
}

type HealthResponse struct {
	Status       HealthStatus `json:"status"`
	ResponseTime string       `json:"response_time"`
}
