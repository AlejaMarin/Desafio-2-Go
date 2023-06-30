package store

import "github.com/AlejaMarin/Desafio-2-Go/internal/domain"

type StoreInterface interface {
	GetDentistById(id int) (domain.Dentista, error)
	CreateDentist(d domain.Dentista) (int, error)
	UpdateDentist(d domain.Dentista) error
	DeleteDentist(id int) error
	ExistsDentistByMatricula(Matricula string) bool
	GetPatientById(id int) (domain.Paciente, error)
	CreatePatient(p domain.Paciente) (int, error)
	UpdatePatient(p domain.Paciente) error
	DeletePatient(id int) error
	ExistsPatientByDNI(DNI string) bool
	GetShiftById(id int) (domain.Turno, error)
	CreateShift(t domain.Turno) (int, error)
	UpdateShift(t domain.Turno) error
	ExistsShift(f, h string, idD int) bool
}
