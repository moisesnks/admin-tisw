package routes

import (
	"admin/api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigureRoutes(r *mux.Router) {
	r.Handle("/", http.HandlerFunc(handlers.HomeHandler))

	// Rutas para las all
	r.Handle("/all_aeropuertos", http.HandlerFunc(handlers.All_Aeropuertos))
	r.Handle("/all_ciudades", http.HandlerFunc(handlers.All_Ciudades))
	r.Handle("/all_paises", http.HandlerFunc(handlers.All_Paises))
	r.Handle("/all_paquetes", http.HandlerFunc(handlers.All_Paquetes))
}
