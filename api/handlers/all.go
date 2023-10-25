package handlers

import (
	"admin/api/models"
	"admin/api/utils"
	"encoding/json"
	"log"
	"net/http"
)

// All_Aeropuertos maneja las solicitudes a la ruta "/all_aeropuertos"
func All_Aeropuertos(w http.ResponseWriter, r *http.Request) {
	// Abrir conexión
	db, err := utils.OpenDB()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error al abrir la conexión a la base de datos", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Consulta SQL
	rows, err := db.Query(`
	SELECT aeropuerto.*, ciudad.nombre as nombre_ciudad,ciudad.pais_id as id_pais, pais.nombre as nombre_pais, pais.abreviacion as abrev_pais
	FROM aeropuerto
	INNER JOIN ciudad ON aeropuerto.ciudad_id = ciudad.id
	INNER JOIN pais ON ciudad.pais_id = pais.id;
	`)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error al consultar la base de datos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Crear un slice de Aeropuerto
	aeropuertos := make([]models.Aeropuerto, 0)

	// Iterar por cada fila
	for rows.Next() {
		// Crear un nuevo Aeropuerto
		aeropuerto := models.Aeropuerto{}

		// Llenar el Aeropuerto con los datos de la fila
		err := rows.Scan(&aeropuerto.ID, &aeropuerto.Nombre, &aeropuerto.CiudadID, &aeropuerto.NombreCiudad, &aeropuerto.PaisID, &aeropuerto.NombrePais, &aeropuerto.AbrevPais)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Error al escanear la base de datos", http.StatusInternalServerError)
			return
		}

		// Agregar el Aeropuerto al slice
		aeropuertos = append(aeropuertos, aeropuerto)
	}

	// Verificar si hubo un error al iterar por las filas
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		http.Error(w, "Error al iterar por las filas", http.StatusInternalServerError)
		return
	}

	// Convertir el slice de Aeropuerto a JSON
	aeropuertosJSON, err := json.Marshal(aeropuertos)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		return
	}

	// Devolver el JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(aeropuertosJSON)

}

// All_Paises maneja las solicitudes a la ruta "/all_paises"
func All_Paises(w http.ResponseWriter, r *http.Request) {
	// Abrir conexión
	db, err := utils.OpenDB()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error al abrir la conexión a la base de datos", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Consulta SQL
	rows, err := db.Query(`
	SELECT *
	FROM pais;
	`)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error al consultar la base de datos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Crear un slice de Pais
	paises := make([]models.Pais, 0)

	// Iterar por cada fila
	for rows.Next() {
		// Crear un nuevo Pais
		pais := models.Pais{}

		// Llenar el Pais con los datos de la fila
		err := rows.Scan(&pais.ID, &pais.Nombre, &pais.Abreviacion)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Error al escanear la base de datos", http.StatusInternalServerError)
			return
		}

		// Agregar el Pais al slice
		paises = append(paises, pais)
	}

	// Verificar si hubo un error al iterar por las filas
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		http.Error(w, "Error al iterar por las filas", http.StatusInternalServerError)
		return
	}

	// Convertir el slice de Pais a JSON
	paisesJSON, err := json.Marshal(paises)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		return
	}

	// Devolver el JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(paisesJSON)

}

func All_Ciudades(w http.ResponseWriter, r *http.Request) {
	// Abrir conexión
	db, err := utils.OpenDB()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error al abrir la conexión a la base de datos", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Consulta SQL
	rows, err := db.Query(`
	SELECT ciudad.*, pais.nombre as nombre_pais, pais.abreviacion as abrev_pais
	FROM ciudad
	INNER JOIN pais ON ciudad.pais_id = pais.id;
	`)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error al consultar la base de datos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Crear un slice de Ciudad
	ciudades := make([]models.Ciudad, 0)

	// Iterar por cada fila
	for rows.Next() {
		// Crear un nuevo Ciudad
		ciudad := models.Ciudad{}

		// Llenar el Ciudad con los datos de la fila
		err := rows.Scan(&ciudad.ID, &ciudad.Nombre, &ciudad.PaisID, &ciudad.NombrePais, &ciudad.AbrevPais)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Error al escanear la base de datos", http.StatusInternalServerError)
			return
		}

		// Agregar el Ciudad al slice
		ciudades = append(ciudades, ciudad)
	}

	// Verificar si hubo un error al iterar por las filas
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		http.Error(w, "Error al iterar por las filas", http.StatusInternalServerError)
		return
	}

	// Convertir el slice de Ciudad a JSON
	ciudadesJSON, err := json.Marshal(ciudades)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		return
	}

	// Devolver el JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(ciudadesJSON)

}

func All_Paquetes(w http.ResponseWriter, r *http.Request) {
	// Abrir conexión
	db, err := utils.OpenDB()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error al abrir la conexión a la base de datos", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Consulta SQL
	rows, err := db.Query(`
    SELECT
		paquete.*,
		ciudad_origen.nombre AS nombre_ciudad_origen,
		ciudad_destino.nombre AS nombre_ciudad_destino,
		habitacionhotel.id AS habitacion_id,
		habitacionhotel.opcion_hotel_id AS opcion_hotel_id,
		opcionhotel.nombre AS nombre_opcion_hotel,
		habitacionhotel.descripcion AS descripcion_habitacion,
		habitacionhotel.servicios AS servicios_habitacion,
		habitacionhotel.precio_noche AS precio_noche,
		hotel.id AS hotel_id,
		hotel.nombre AS nombre_hotel,
		hotel.ciudad_id AS ciudad_id_hotel,
		hotel.direccion AS direccion_hotel,
		hotel.valoracion AS valoracion_hotel,
		hotel.descripcion AS descripcion_hotel,
		hotel.servicios AS servicios_hotel,
		hotel.telefono AS telefono_hotel,
		hotel.correo_electronico AS correo_electronico_hotel,
		hotel.sitio_web AS sitio_web_hotel
	FROM
		paquete
		INNER JOIN habitacionhotel ON paquete.id_hh = habitacionhotel.id
		INNER JOIN hotel ON habitacionhotel.hotel_id = hotel.id
		INNER JOIN opcionhotel ON habitacionhotel.opcion_hotel_id = opcionhotel.id
		INNER JOIN ciudad ciudad_origen ON paquete.id_origen = ciudad_origen.id
		INNER JOIN ciudad ciudad_destino ON paquete.id_destino = ciudad_destino.id;
    `)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error al consultar la base de datos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Crear un slice de PaqueteInfo
	paquetes := make([]models.PaqueteInfo, 0)

	// Iterar por cada fila
	for rows.Next() {
		var paqueteInfo models.PaqueteInfo
		var infoPaquete models.PaqueteInfoAdicional
		var hotelInfo models.HotelInfo

		// Escanear la fila y verificar errores
		err := rows.Scan(
			&paqueteInfo.ID,
			&paqueteInfo.Nombre,
			&paqueteInfo.IdOrigen,
			&paqueteInfo.IdDestino,
			&paqueteInfo.IdHabitacionHotel,
			&paqueteInfo.Descripcion,
			&paqueteInfo.Detalles,
			&paqueteInfo.Precio,
			&paqueteInfo.NombreCiudadOrigen,
			&paqueteInfo.NombreCiudadDestino,
			&infoPaquete.HabitacionID,
			&infoPaquete.OpcionHotelID,
			&infoPaquete.NombreOpcionHotel,
			&infoPaquete.DescripcionHabitacion,
			&infoPaquete.ServiciosHabitacion,
			&infoPaquete.PrecioNoche,
			&hotelInfo.ID,
			&hotelInfo.Nombre,
			&hotelInfo.CiudadID,
			&hotelInfo.Direccion,
			&hotelInfo.Valoracion,
			&hotelInfo.DescripcionHotel,
			&hotelInfo.ServiciosHotel,
			&hotelInfo.Telefono,
			&hotelInfo.CorreoElectronico,
			&hotelInfo.SitioWeb,
		)

		// Verificar si hubo un error al escanear la fila
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Error al escanear la base de datos", http.StatusInternalServerError)
			return
		}

		// Asignar la información de la habitación y hotel a la estructura principal
		infoPaquete.HotelInfo = hotelInfo
		paqueteInfo.InfoPaquete = infoPaquete

		// Agregar el PaqueteInfo al slice
		paquetes = append(paquetes, paqueteInfo)
	}

	// Convertir el slice de PaqueteInfo a JSON
	paquetesJSON, err := json.Marshal(paquetes)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		return
	}

	// Devolver el JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(paquetesJSON)
}
