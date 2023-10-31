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

	// Rutas para los Schemas

	// Ruta para las imágenes
	r.Handle("/api/admin/imagenes/listar", http.HandlerFunc(handlers.ListarTodasLasImagenes))
	// Ruta para subir imágenes
	r.Handle("/api/admin/imagenes/subir", http.HandlerFunc(handlers.PostImagen)).Methods("POST")
	// Ruta para ver las imágenes en la bd
	r.Handle("/api/admin/imagenes/bd", http.HandlerFunc(handlers.ImagenesBd)).Methods("GET")
	// Ruta para ver las imágenes en el bucket
	r.Handle("/api/admin/imagenes/bucket", http.HandlerFunc(handlers.ImagenesBucket)).Methods("GET")

	// Facturación
	r.Handle("/api/admin/facturacion", http.HandlerFunc(handlers.GetAllFacturas))
	r.Handle("/api/admin/facturacion/usuario", http.HandlerFunc(handlers.GetFacturasByUsuarios))
	r.Handle("/api/admin/facturacion/crear", http.HandlerFunc(handlers.CreateFactura))
	r.Handle("/api/admin/facturacion/actualizar", http.HandlerFunc(handlers.UpdateFactura))
	r.Handle("/api/admin/facturacion/eliminar", http.HandlerFunc(handlers.DeleteFactura))

	// Paises
	r.Handle("/api/admin/paises", http.HandlerFunc(handlers.GetAllPaises))
	r.Handle("/api/admin/paises/crear", http.HandlerFunc(handlers.CreatePais))
	r.Handle("/api/admin/paises/actualizar", http.HandlerFunc(handlers.UpdatePais))

	//Paquetes
	r.Handle("/api/admin/paquetes", http.HandlerFunc(handlers.GetAllPaquetes))
	r.Handle("/api/admin/paquetes/crear", http.HandlerFunc(handlers.CreatePaquete))
	r.Handle("/api/admin/paquetes/eliminar", http.HandlerFunc(handlers.DeletePaquete))
	r.Handle("/api/admin/paquetes/actualizar", http.HandlerFunc(handlers.UpdatePaquete))

	//Ciudades
	r.Handle("/api/admin/ciudades", http.HandlerFunc(handlers.GetAllCiudades))
	r.Handle("/api/admin/ciudades/crear", http.HandlerFunc(handlers.CreateCiudad))
	r.Handle("/api/admin/ciudades/eliminar", http.HandlerFunc(handlers.DeleteCiudad))
	r.Handle("/api/admin/ciudades/actualizar", http.HandlerFunc(handlers.UpdateCiudad))

	//Aeropuertos
	r.Handle("/api/admin/aeropuertos", http.HandlerFunc(handlers.GetAllAeropuertos))
	r.Handle("/api/admin/aeropuertos/crear", http.HandlerFunc(handlers.CreateAeropuerto))
	r.Handle("/api/admin/aeropuertos/eliminar", http.HandlerFunc(handlers.DeleteAeropuerto))
}
