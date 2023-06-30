package store

import "github.com/AlejaMarin/Desafio-2-Go/internal/domain"

type StoreInterface interface {
	GetPatientById(id int) (domain.Paciente, error)
	CreatePatient(p domain.Paciente) (int, error)
	UpdatePatient(p domain.Paciente) error
	DeletePatient(id int) error
	ExistsPatientByDNI(DNI string) bool
}
