package handlers

import (
	"admin/api/utils"
	"encoding/json"
	"net/http"
)

// AeropuertoCreate representa la estructura para la creación de un nuevo aeropuerto
type AeropuertoCreate struct {
	Nombre   string `json:"nombre"`
	CiudadID int    `json:"ciudad_id"`
}

func CreateAeropuerto(w http.ResponseWriter, r *http.Request) {
	// Obtener los datos para la creación del aeropuerto desde el cuerpo de la solicitud
	var aeropuertoCreate AeropuertoCreate

	err := json.NewDecoder(r.Body).Decode(&aeropuertoCreate)
	if err != nil {
		handleError(w, "Error al decodificar los datos de creación", http.StatusBadRequest, err)
		return
	}

	// Crear el nuevo aeropuerto en la base de datos
	err = createAeropuerto(aeropuertoCreate)
	if err != nil {
		handleError(w, "Error al crear el aeropuerto", http.StatusInternalServerError, err)
		return
	}

	// Responder con éxito y proporcionar una respuesta JSON
	response := map[string]interface{}{
		"status":  "success",
		"message": "Aeropuerto creado con éxito",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func createAeropuerto(aeropuerto AeropuertoCreate) error {
	// Realizar la verificación en la base de datos
	db, err := utils.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Realizar la inserción en la base de datos
	_, err = db.Exec(`
        INSERT INTO aeropuerto (nombre, ciudad_id)
        VALUES ($1, $2)
    `, aeropuerto.Nombre, aeropuerto.CiudadID)

	return err
}
