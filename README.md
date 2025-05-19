# Incident Service

Servicio REST para manejo de incidentes utilizando Go y Gin.(Secure app)

## Características

- API REST para manejo de incidentes
- Gestión de incidentes (creación, consulta)
- Documentación de endpoints
- Manejo de errores

## Requisitos

- Go 1.20 o superior
- Docker (opcional)
- Base de datos (configurable)

## Instalación

1. Clona el repositorio:
```bash
git clone https://github.com/pav-dev98/incident_service.git
cd incident_service
```

2. Instala las dependencias:
```bash
go mod tidy
```

3. Ejecuta el servicio:
```bash
go run main.go
```

## Uso

El servicio estará disponible en `http://localhost:8080` por defecto.

### Endpoints disponibles

- `GET /`: Mensaje de bienvenida
- `GET /incidents`: Lista todos los incidentes
- `POST /incidents`: Crea un nuevo incidente

## Estructura del Proyecto

```
incident_service/
├── internal/
│   ├── handlers/        # Handlers de la API
│   └── repository/      # Repositorio de datos
├── main.go              # Punto de entrada
├── go.mod              # Dependencias de Go
└── docker/             # Configuración de Docker
```

## Desarrollo

Para contribuir al proyecto:

1. Crea una rama:
```bash
git checkout -b feature/tu-feature
```

2. Commit tus cambios:
```bash
git commit -m "feat: tu cambio"
```

3. Push a la rama:
```bash
git push origin feature/tu-feature
```

## Licencia

MIT

## Autor

Pavel Mansilla (pav-dev98)