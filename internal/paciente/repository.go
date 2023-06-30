package paciente

import (
	"errors"

	"github.com/AlejaMarin/Desafio-2-Go/internal/domain"
	"github.com/AlejaMarin/Desafio-2-Go/pkg/store"
)

type Repository interface {
	GetPatientById(id int) (domain.Paciente, error)
	CreatePatient(p domain.Paciente) (domain.Paciente, error)
	UpdatePatient(id int, p domain.Paciente) (domain.Paciente, error)
	DeletePatient(id int) error
}

type repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetPatientById(id int) (domain.Paciente, error) {

	patient, err := r.storage.GetPatientById(id)
	if err != nil {
		return domain.Paciente{}, errors.New("Paciente No Encontrado")
	}
	return patient, nil

}

func (r *repository) CreatePatient(p domain.Paciente) (domain.Paciente, error) {

	if r.storage.ExistsPatientByDNI(p.DNI) {
		return domain.Paciente{}, errors.New("Ya existe un Paciente con el DNI ingresado")
	}
	id, err := r.storage.CreatePatient(p)
	if err != nil {
		return domain.Paciente{}, errors.New("No se pudo registrar el paciente")
	}
	p.Id = id
	return p, nil

}

func (r *repository) UpdatePatient(id int, p domain.Paciente) (domain.Paciente, error) {

	patient, err := r.GetPatientById(id)
	if err != nil {
		return domain.Paciente{}, err
	}

	if r.storage.ExistsPatientByDNI(p.DNI) && p.DNI != patient.DNI {
		return domain.Paciente{}, errors.New("Ya existe un Paciente con el DNI ingresado")
	}
	err = r.storage.UpdatePatient(p)
	if err != nil {
		return domain.Paciente{}, errors.New("No se pudo actualizar el paciente")
	}
	return p, nil

}

func (r *repository) DeletePatient(id int) error {

	err := r.storage.DeletePatient(id)
	if err != nil {
		return err
	}
	return nil

}
