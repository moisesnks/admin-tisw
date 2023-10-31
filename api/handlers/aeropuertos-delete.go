package handlers

import (
	"admin/api/utils"
	"encoding/json"
	"net/http"
)

// AeropuertoDeleteRequest representa la estructura para eliminar un aeropuerto por su ID
type AeropuertoDeleteRequest struct {
	ID int `json:"id"`
}

func DeleteAeropuerto(w http.ResponseWriter, r *http.Request) {
	// Obtener los datos para eliminar el aeropuerto desde el cuerpo de la solicitud
	var deleteRequest AeropuertoDeleteRequest

	err := json.NewDecoder(r.Body).Decode(&deleteRequest)
	if err != nil {
		handleError(w, "Error al decodificar los datos de eliminación", http.StatusBadRequest, err)
		return
	}

	// Realizar la eliminación en la base de datos
	err = deleteAeropuerto(deleteRequest.ID)
	if err != nil {
		handleError(w, "Error al eliminar el aeropuerto", http.StatusInternalServerError, err)
		return
	}

	// Responder con éxito y proporcionar una respuesta JSON
	response := map[string]interface{}{
		"status":  "success",
		"message": "Aeropuerto eliminado con éxito",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func deleteAeropuerto(id int) error {
	// Realizar la eliminación en la base de datos
	db, err := utils.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Realizar la consulta DELETE en la base de datos utilizando el ID
	_, err = db.Exec("DELETE FROM aeropuerto WHERE id = $1", id)

	return err
}
