package domain

type Paciente struct {
	Id        int    `json:"id" swaggerignore:"true"`
	Nombre    string `json:"nombre" binding:"required"`
	Apellido  string `json:"apellido" binding:"required"`
	Domicilio string `json:"domicilio" binding:"required"`
	DNI       string `json:"dni" binding:"required"`
	FechaAlta string `json:"fechaAlta" binding:"required"`
}
