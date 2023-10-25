package middleware

import (
	"net/http"
)

// CorsMiddleware verifica el origen de la solicitud y permite el acceso solo a orígenes autorizados.
func CorsMiddleware(allowedOrigins []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			if isOriginAllowed(origin, allowedOrigins) {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET")
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Not allowed", http.StatusForbidden)
			}
		})
	}
}

// isOriginAllowed verifica si el origen de la solicitud está en la lista de orígenes permitidos.
func isOriginAllowed(origin string, allowedOrigins []string) bool {
	for _, allowedOrigin := range allowedOrigins {
		if allowedOrigin == origin {
			return true
		}
	}
	return false
}
