package dentista

import (
	"errors"

	"github.com/AlejaMarin/Desafio-2-Go/internal/domain"
	"github.com/AlejaMarin/Desafio-2-Go/pkg/store"
)

type Repository interface {
	GetDentistById(id int) (domain.Dentista, error)
	CreateDentist(d domain.Dentista) (domain.Dentista, error)
	UpdateDentist(id int, d domain.Dentista) (domain.Dentista, error)
	DeleteDentist(id int) error
}

type repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetDentistById(id int) (domain.Dentista, error) {

	dentist, err := r.storage.GetDentistById(id)
	if err != nil {
		return domain.Dentista{}, errors.New("odontólogo no encontrado")
	}
	return dentist, nil

}

func (r *repository) CreateDentist(d domain.Dentista) (domain.Dentista, error) {

	if r.storage.ExistsDentistByMatricula(d.Matricula) {
		return domain.Dentista{}, errors.New("ya existe un odontólogo con la matrícula ingresada")
	}
	id, err := r.storage.CreateDentist(d)
	if err != nil {
		return domain.Dentista{}, errors.New("no se pudo registrar al odontólogo")
	}
	d.Id = id
	return d, nil

}

func (r *repository) UpdateDentist(id int, d domain.Dentista) (domain.Dentista, error) {

	dentist, err := r.GetDentistById(id)
	if err != nil {
		return domain.Dentista{}, err
	}

	if r.storage.ExistsDentistByMatricula(d.Matricula) && d.Matricula != dentist.Matricula {
		return domain.Dentista{}, errors.New("ya existe un odontólogo con la matrícula ingresada")
	}
	err = r.storage.UpdateDentist(d)
	if err != nil {
		return domain.Dentista{}, errors.New("no se pudo actualizar al odontólogo")
	}
	return d, nil

}

func (r *repository) DeleteDentist(id int) error {

	err := r.storage.DeleteDentist(id)
	if err != nil {
		return err
	}
	return nil

}
