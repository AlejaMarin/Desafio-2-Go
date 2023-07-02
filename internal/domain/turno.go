package domain

type Turno struct {
	Id          int    `json:"id" swaggerignore:"true"`
	IdPaciente  int    `json:"idPaciente" binding:"required"`
	IdDentista  int    `json:"idDentista" binding:"required"`
	Fecha       string `json:"fecha" binding:"required"`
	Hora        string `json:"hora" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
}

type TurnoDos struct {
	DniPaciente       string `json:"dniPaciente" binding:"required"`
	MatriculaDentista string `json:"matriculaDentista" binding:"required"`
	Fecha             string `json:"fecha" binding:"required"`
	Hora              string `json:"hora" binding:"required"`
	Descripcion       string `json:"descripcion" binding:"required"`
}

type TurnoByDni struct {
	Fecha       string   `json:"fecha" binding:"required"`
	Hora        string   `json:"hora" binding:"required"`
	Descripcion string   `json:"descripcion" binding:"required"`
	Paciente    Paciente `json:"paciente" binding:"required"`
	Dentista    Dentista `json:"dentista" binding:"required"`
}
