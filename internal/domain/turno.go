package domain

type Turno struct {
	Id          int    `json:"id"`
	IdPaciente  int    `json:"idPaciente" binding:"required"`
	IdDentista  int    `json:"idDentista" binding:"required"`
	Fecha       string `json:"fecha" binding:"required"`
	Hora        string `json:"hora" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
}
