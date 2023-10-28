package models

// PaqueteInfo representa la información de un paquete turístico.
type PaqueteInfo struct {
	ID                  int                  `json:"id"`
	Nombre              string               `json:"nombre"`
	IdOrigen            int                  `json:"id_origen"`
	IdDestino           int                  `json:"id_destino"`
	Descripcion         string               `json:"descripcion"`
	Detalles            string               `json:"detalles"`
	Precio              int                  `json:"precio_viaje"`
	IdHabitacionHotel   string               `json:"id_hh"`
	Imagenes            string               `json:"imagenes"`
	TotalPersonas       int                  `json:"total_personas"`
	NombreCiudadOrigen  string               `json:"nombre_ciudad_origen"`
	NombreCiudadDestino string               `json:"nombre_ciudad_destino"`
	InfoPaquete         PaqueteInfoAdicional `json:"info_paquete"`
}

// PaqueteInfoAdicional representa la información adicional de habitación y hotel en un paquete turístico.
type PaqueteInfoAdicional struct {
	HabitacionID          int       `json:"habitacion_id"`
	OpcionHotelID         int       `json:"opcion_hotel_id"`
	NombreOpcionHotel     string    `json:"nombre_opcion_hotel"`
	DescripcionHabitacion string    `json:"descripcion_habitacion"`
	ServiciosHabitacion   string    `json:"servicios_habitacion"`
	PrecioNoche           int       `json:"precio_noche"`
	HotelInfo             HotelInfo `json:"hotel_info"`
}

// HotelInfo representa la información de un hotel.
type HotelInfo struct {
	ID                int     `json:"id"`
	Nombre            string  `json:"nombre"`
	CiudadID          int     `json:"ciudad_id"`
	Direccion         string  `json:"direccion"`
	Valoracion        float64 `json:"valoracion"`
	DescripcionHotel  string  `json:"descripcion_hotel"`
	ServiciosHotel    string  `json:"servicios_hotel"`
	Telefono          string  `json:"telefono"`
	CorreoElectronico string  `json:"correo_electronico"`
	SitioWeb          string  `json:"sitio_web"`
}
