package handlers

import (
	"admin/api/utils"
	"encoding/json"
	"net/http"
)

// PaisCreate representa la estructura para la creación de un nuevo país
type PaisCreate struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Abreviacion string `json:"abreviacion"`
	// Otras propiedades de país, si las hubiera
}

func CreatePais(w http.ResponseWriter, r *http.Request) {
	// Obtener los datos para la creación del país desde el cuerpo de la solicitud
	var paisCreate PaisCreate

	err := json.NewDecoder(r.Body).Decode(&paisCreate)
	if err != nil {
		handleError(w, "Error al decodificar los datos de creación", http.StatusBadRequest, err)
		return
	}

	// Crear el nuevo país en la base de datos
	err = createPais(paisCreate.ID, paisCreate.Nombre, paisCreate.Abreviacion)
	if err != nil {
		handleError(w, "Error al crear el país", http.StatusInternalServerError, err)
		return
	}

	// Responder con éxito
	w.WriteHeader(http.StatusCreated)
}

func createPais(id int, nombre string, abreviacion string) error {
	// Realizar la inserción en la base de datos
	db, err := utils.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`
        INSERT INTO pais (id, nombre, abreviacion)
        VALUES ($1, $2, $3)
    `, id, nombre, abreviacion)

	return err
}
