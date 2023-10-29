package routes

import (
	"admin/api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigureRoutes(r *mux.Router) {
	// allowedOrigins := []string{"http://admin.lumonidy.studio", "http://localhost:3000"}

	// c := middleware.CorsMiddleware(allowedOrigins)
	// r.Use(c)

	r.Handle("/api/admin/", http.HandlerFunc(handlers.HomeHandler))

	// Rutas para las all Schemas
	r.Handle("/api/admin/all_aeropuertos", http.HandlerFunc(handlers.GetAllAeropuertos))
	r.Handle("/api/admin/all_ciudades", http.HandlerFunc(handlers.GetAllCiudades))
	r.Handle("/api/admin/all_paises", http.HandlerFunc(handlers.GetAllPaises))
	r.Handle("/api/admin/all_paquetes", http.HandlerFunc(handlers.GetAllPaquetes))

	// Ruta para las imágenes
	r.Handle("/api/admin/imagenes/listar", http.HandlerFunc(handlers.ListarTodasLasImagenes))
	r.Handle("/api/admin/imagenes/listar/bucket", http.HandlerFunc(handlers.ImagenesBucket))
	r.Handle("/api/admin/imagenes/listar/bd", http.HandlerFunc(handlers.ImagenesBd))

	r.Handle("/api/admin/imagenes/subir", http.HandlerFunc(handlers.PostImagen)).Methods("POST")

}
