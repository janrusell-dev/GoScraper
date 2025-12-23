package dto

type HealthCheckResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Version   string `json:"version"`
}

type IndexResponse struct {
	Name      string            `json:"name"`
	Version   string            `json:"version"`
	Endpoints map[string]string `json:"endpoints"`
}
