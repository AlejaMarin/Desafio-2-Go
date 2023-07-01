package store

import "github.com/AlejaMarin/Desafio-2-Go/internal/domain"

type StoreInterface interface {
	/* ----- ODONTÓLOGO ----- */
	GetDentistById(id int) (domain.Dentista, error)
	CreateDentist(d domain.Dentista) (int, error)
	UpdateDentist(d domain.Dentista) error
	DeleteDentist(id int) error
	ExistsDentistByMatricula(Matricula string) bool
	/* ----- PACIENTE ----- */
	GetPatientById(id int) (domain.Paciente, error)
	CreatePatient(p domain.Paciente) (int, error)
	UpdatePatient(p domain.Paciente) error
	DeletePatient(id int) error
	ExistsPatientByDNI(DNI string) bool
	/* ----- TURNO ----- */
	GetShiftById(id int) (domain.Turno, error)
	CreateShift(t domain.Turno) (int, error)
	UpdateShift(t domain.Turno) error
	ExistsShift(f, h string, idD int) bool
	DeleteShift(id int) error
	GetPatientIdByDni(dni string) (int, error)
	GetDentistByMatricula(matricula string) (int, error)
	GetShiftsByDniPatient(dni string) ([]domain.TurnoByDni, error)
}
