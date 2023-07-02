package domain

type Dentista struct {
	Id        int    `json:"id" swaggerignore:"true"` 
	Apellido  string `json:"apellido" binding:"required"`
	Nombre    string `json:"nombre" binding:"required"`
	Matricula string `json:"matricula" binding:"required"`
}
