package handlers

import (
	"net/http"
    // "time" // Necesario si usas time.Time en tus structs

	"github.com/gin-gonic/gin"
	"github.com/pav-dev98/incident_service/internal/incident" // Importa tu modelo
	// Importa tu repositorio y/o servicio aquí más adelante
	// "incident_service/internal/repository"
)

// GetIncidents maneja la solicitud GET /incidents
func GetIncidents(c *gin.Context) {
	// TODO: Aquí iría la lógica para consultar la base de datos
	// usando el repositorio. Podrías filtrar por query params (ubicación, tipo, tiempo).

    // Ejemplo: obtener query params para filtrado (ej. /incidents?type=Robo)
    // incidentTypeFilter := c.Query("type")

	// Placeholder: Devuelve una lista vacía o datos de prueba por ahora
	incidents := []incident.Incident{
        // {ID: "1", Type: "Robo", Latitude: -9.933, Longitude: -76.075, Timestamp: time.Now(), Status: "Reportado"},
    }


	c.JSON(http.StatusOK, incidents) // Devuelve la lista de incidencias como JSON
}

// CreateIncident maneja la solicitud POST /incidents
func CreateIncident(c *gin.Context) {
	var req incident.CreateIncidentRequest // Usa la estructura de request si es diferente al modelo

	// 1. Bind (enlaza) el cuerpo de la solicitud JSON a la estructura req
	if err := c.ShouldBindJSON(&req); err != nil {
		// Si el binding falla (ej. JSON mal formado, campos requeridos faltantes si usaste binding:"required")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return // Sale de la función handler
	}

	// TODO: Aquí iría la lógica de negocio para:
	// 1. Obtener el ID del usuario que reporta (del token de autenticación, que manejaría el API Gateway o middleware).
    // 2. Validar más a fondo los datos si es necesario.
	// 3. Crear una instancia del modelo incidente.Incident a partir de req.
	//    newIncident := incident.Incident{ ... campos llenados desde req y otros generados ... }
	// 4. Guardar el nuevo incidente en la base de datos usando el repositorio.
    //    repository.SaveIncident(newIncident)
	// 5. Notificar a otros servicios si es necesario (ej. Servicio de Alertas y Notificaciones).

	// Placeholder: Simula la creación exitosa y devuelve el objeto creado (o su ID)
    // Por ahora, solo mostramos lo que recibimos
	c.JSON(http.StatusCreated, gin.H{
        "message": "Incidencia recibida (TODO: guardar y procesar)",
        "received_data": req,
        // "id":      "id-generado-real", // Devolver el ID real después de guardar
    })
}

// TODO: Añade handlers para GET /incidents/:id, PUT /incidents/:id, DELETE /incidents/:id
// Necesitarás leer parámetros de la URL (ej. id := c.Param("id"))