package handlers

import (
	"admin/api/utils"
	"encoding/json"
	"net/http"
)

// PaisUpdate representa la estructura para la actualización parcial de un país
type PaisUpdate struct {
	ID          int     `json:"id"`
	Nombre      *string `json:"nombre,omitempty"`
	Abreviacion *string `json:"abreviacion,omitempty"`
	Imagenes    *string `json:"imagenes,omitempty"`
}

// Handler que llama a la función updatePais y devuelve una respuesta JSON
func UpdatePais(w http.ResponseWriter, r *http.Request) {
	// Obtener los datos actualizados del país desde el cuerpo de la solicitud
	var paisUpdate PaisUpdate

	err := json.NewDecoder(r.Body).Decode(&paisUpdate)
	if err != nil {
		handleError(w, "Error al decodificar los datos de actualización", http.StatusBadRequest, err)
		return
	}

	// Actualizar el país en la base de datos
	err = updatePais(paisUpdate.ID, paisUpdate.Nombre, paisUpdate.Abreviacion, paisUpdate.Imagenes)
	if err != nil {
		handleError(w, "Error al actualizar el país", http.StatusInternalServerError, err)
		return
	}

	// Responder con éxito
	w.WriteHeader(http.StatusOK)
}

// Realizar la actualización en la base de datos
func updatePais(id int, nombre *string, abreviacion *string, imagenes *string) error {

	db, err := utils.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Construir la consulta SQL basada en los campos no nulos
	query := "UPDATE Pais SET"
	args := make([]interface{}, 0)

	if nombre != nil {
		query += " nombre = $1,"
		args = append(args, *nombre)
	}

	if abreviacion != nil {
		query += " abreviacion = $2,"
		args = append(args, *abreviacion)
	}

	if imagenes != nil {
		query += " imagenes = $3,"
		args = append(args, *imagenes)
	}

	// Eliminar la última coma
	query = query[:len(query)-1]

	// Agregar la condición WHERE
	query += " WHERE id = $4"
	args = append(args, id)

	// Ejecutar la consulta
	_, err = db.Exec(query, args...)
	return err
}
