package main

import (
	"log"
	"os"
	// "time" // Importa time si necesitas time.Now() o similar en main (probablemente no)

	"github.com/gin-gonic/gin"
	"github.com/pav-dev98/incident_service/internal/handlers" // Importa tus handlers
    // Importa inicialización de base de datos aquí
    // "incident_service/internal/repository"
	"net/http"
)

func main() {
    // TODO: Inicializar la conexión a la base de datos aquí antes de configurar el router
    // db, err := repository.NewDatabaseConnection(...)
    // if err != nil {
    //     log.Fatalf("Failed to connect to database: %v", err)
    // }
    // defer db.Close() // Asegúrate de cerrar la conexión al salir

    // Configura el modo de Gin (release en producción, debug en desarrollo)
    // gin.SetMode(gin.ReleaseMode) // Descomentar para producción

	router := gin.Default() // Crea un router Gin con middleware por defecto (Logger, Recovery)

    // Define las rutas de la API y asócialas a tus handlers
	router.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Bienvenido al Incident Service",
		})
	})
    router.GET("/incidents", handlers.GetIncidents) // GET /incidents -> handlers.GetIncidents
    router.POST("/incidents", handlers.CreateIncident) // POST /incidents -> handlers.CreateIncident

	// TODO: Añade rutas para /incidents/:id (GET, PUT, DELETE)

    // Define el puerto para el servidor
    // Lee la variable de entorno PORT, si no existe usa 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

    // Configura el servidor HTTP
    // Esto es más flexible que router.Run() para opciones avanzadas,
    // pero router.Run() es suficiente para empezar.
    // srv := &http.Server{
    //     Addr:    ":" + port,
    //     Handler: router,
    //     ReadTimeout: 10 * time.Second, // Tiempos de espera
    //     WriteTimeout: 10 * time.Second,
    //     MaxHeaderBytes: 1 << 20, // 1MB
    // }

	log.Printf(" Incident Service corriendo en el puerto %s", port)

    // Inicia el servidor HTTP. router.Run() es un shorthand para iniciar el servidor.
    // Si necesitas configurar SSL/TLS, tiempos de espera, etc., usa http.Server{}
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

    // Si usaras http.Server{}:
    // if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
    //    log.Fatalf("Error al iniciar el servidor: %v", err)
    // }
}