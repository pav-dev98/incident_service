package incident

import "time"

// Incident representa una incidencia reportada por un ciudadano
type Incident struct {
	ID             string    `json:"id"` // Usar UUID o similar
	Type           string    `json:"type"` // Ej: "Robo", "Incendio", "Extorsion"
	Description    string    `json:"description,omitempty"` // Opcional
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	Timestamp      time.Time `json:"timestamp"` // Fecha y hora del incidente/reporte
	ReporterUserID string    `json:"reporter_user_id"` // ID del usuario que reportó
	Status         string    `json:"status"` // Ej: "Reportado", "En Proceso", "Cerrado"
	// Podrías añadir campos para adjuntos (ej: una lista de IDs de adjuntos)
}

// CreateIncidentRequest representa el payload de entrada al crear una incidencia
// Esto es útil si la estructura de entrada es diferente al modelo interno
type CreateIncidentRequest struct {
    Type         string    `json:"type" binding:"required"` // 'binding:"required"' para validación básica con Gin
    Description  string    `json:"description,omitempty"`
    Latitude     float64   `json:"latitude" binding:"required"`
    Longitude    float64   `json:"longitude" binding:"required"`
    Timestamp    time.Time `json:"timestamp"` // Gin puede bindear esto si el formato es correcto
}