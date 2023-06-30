package turno

import (
	"errors"

	"github.com/AlejaMarin/Desafio-2-Go/internal/domain"
	"github.com/AlejaMarin/Desafio-2-Go/pkg/store"
)

type Repository interface {
	GetShiftById(id int) (domain.Turno, error)
	CreateShift(t domain.Turno) (domain.Turno, error)
	UpdateShift(id int, t domain.Turno) (domain.Turno, error)
}

type repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetShiftById(id int) (domain.Turno, error) {

	shift, err := r.storage.GetShiftById(id)
	if err != nil {
		return domain.Turno{}, errors.New("Turno No Encontrado")
	}
	return shift, nil

}

func (r *repository) CreateShift(t domain.Turno) (domain.Turno, error) {

	if r.storage.ExistsShift(t.Fecha, t.Hora, t.IdDentista) {
		return domain.Turno{}, errors.New("Ya hay un Turno asignado para esa Fecha, Hora y Dentista")
	}
	id, err := r.storage.CreateShift(t)
	if err != nil {
		return domain.Turno{}, errors.New("No se pudo registrar el turno")
	}
	t.Id = id
	return t, nil
}

func (r *repository) UpdateShift(id int, t domain.Turno) (domain.Turno, error) {

	if r.storage.ExistsShift(t.Fecha, t.Hora, t.IdDentista) {
		return domain.Turno{}, errors.New("Ya hay un Turno asignado para esa Fecha, Hora y Dentista")
	}
	err := r.storage.UpdateShift(t)
	if err != nil {
		return domain.Turno{}, errors.New("No se pudo actualizar el turno")
	}
	return t, nil

}
