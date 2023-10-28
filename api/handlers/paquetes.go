package handlers

import (
	"admin/api/models"
	"admin/api/utils"
	"encoding/json"
	"net/http"
)

// GetAllPaquetes obtiene todos los paquetes y los devuelve como JSON.
func GetAllPaquetes(w http.ResponseWriter, r *http.Request) {
	paquetes, err := fetchPaquetes()
	if err != nil {
		handleError(w, "Error al obtener los paquetes", http.StatusInternalServerError, err)
		return
	}

	paquetesJSON, err := json.Marshal(paquetes)
	if err != nil {
		handleError(w, "Error al convertir a JSON", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(paquetesJSON)
}

func fetchPaquetes() ([]models.PaqueteInfo, error) {
	db, err := utils.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT DISTINCT
			paquete.*,
			COALESCE(total_personas, 0) AS total_personas,
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
		INNER JOIN unnest(paquete.id_hh) WITH ORDINALITY t(habitacion_id, ord) ON TRUE
		INNER JOIN habitacionhotel ON t.habitacion_id = habitacionhotel.id
		INNER JOIN hotel ON habitacionhotel.hotel_id = hotel.id
		INNER JOIN opcionhotel ON habitacionhotel.opcion_hotel_id = opcionhotel.id
		INNER JOIN ciudad ciudad_origen ON paquete.id_origen = ciudad_origen.id
		INNER JOIN ciudad ciudad_destino ON paquete.id_destino = ciudad_destino.id
		LEFT JOIN (
			SELECT
				paquete.id AS paquete_id,
				SUM(opcionhotel.cantidad) AS total_personas
			FROM
				paquete
				INNER JOIN unnest(paquete.id_hh) WITH ORDINALITY t(habitacion_id, ord) ON TRUE
				INNER JOIN habitacionhotel ON t.habitacion_id = habitacionhotel.id
				INNER JOIN opcionhotel ON habitacionhotel.opcion_hotel_id = opcionhotel.id
			GROUP BY
				paquete.id
		) AS subquery ON paquete.id = subquery.paquete_id
	ORDER BY
		paquete.id;
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	paquetes := make([]models.PaqueteInfo, 0)

	for rows.Next() {
		var paqueteInfo models.PaqueteInfo
		var infoPaquete models.PaqueteInfoAdicional
		var hotelInfo models.HotelInfo

		err := rows.Scan(
			&paqueteInfo.ID,
			&paqueteInfo.Nombre,
			&paqueteInfo.IdOrigen,
			&paqueteInfo.IdDestino,
			&paqueteInfo.Descripcion,
			&paqueteInfo.Detalles,
			&paqueteInfo.Precio,
			&paqueteInfo.IdHabitacionHotel,
			&paqueteInfo.Imagenes,
			&paqueteInfo.TotalPersonas,
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

		if err != nil {
			return nil, err
		}

		infoPaquete.HotelInfo = hotelInfo
		paqueteInfo.InfoPaquete = infoPaquete

		paquetes = append(paquetes, paqueteInfo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return paquetes, nil
}
