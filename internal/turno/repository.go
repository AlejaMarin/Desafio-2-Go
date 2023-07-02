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
	DeleteShift(id int) error
	CreateShiftByDniAndEnrollment(t2 domain.TurnoDos) (domain.Turno, error)
	GetShiftsByDniPatient(dni string) ([]domain.TurnoByDni, error)
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

	shift, err := r.GetShiftById(id)
	if err != nil {
		return domain.Turno{}, err
	}

	if r.storage.ExistsShift(t.Fecha, t.Hora, t.IdDentista) {
		if t.Fecha != shift.Fecha && t.Hora != shift.Hora && t.IdDentista != shift.IdDentista {
			return domain.Turno{}, errors.New("Ya hay un Turno asignado para esa Fecha, Hora y Dentista")
		}
		return domain.Turno{}, errors.New("Ya hay un Turno asignado para esa Fecha, Hora y Dentista")
	}

	err = r.storage.UpdateShift(t)
	if err != nil {
		return domain.Turno{}, errors.New("No se pudo actualizar el turno")
	}
	return t, nil

}

func (r *repository) DeleteShift(id int) error {
	
	_, err := r.GetShiftById(id)
	if err != nil {
		return err
	}
	
	err = r.storage.DeleteShift(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) CreateShiftByDniAndEnrollment(t2 domain.TurnoDos) (domain.Turno, error) {
	var shift domain.Turno
	idP, err := r.storage.GetPatientIdByDni(t2.DniPaciente)
	if err != nil {
		return domain.Turno{}, errors.New("El paciente no existe")
	}
	idD, err := r.storage.GetDentistByMatricula(t2.MatriculaDentista)
	if err != nil {
		return domain.Turno{}, errors.New("El odontólogo no existe")
	}
	shift = domain.Turno{
		IdPaciente:  idP,
		IdDentista:  idD,
		Fecha:       t2.Fecha,
		Hora:        t2.Hora,
		Descripcion: t2.Descripcion,
	}

	s, err := r.CreateShift(shift)
	if err != nil {
		return domain.Turno{}, err
	}

	return s, nil

}

func (r *repository) GetShiftsByDniPatient(dni string) ([]domain.TurnoByDni, error) {

	if !r.storage.ExistsPatientByDNI(dni) {
		return nil, errors.New("No existe un paciente con el DNI ingresado")
	}

	shifts, err := r.storage.GetShiftsByDniPatient(dni)
	if err != nil {
		return nil, errors.New("No se pudo obtener el/los turno/s")
	}
	return shifts, nil

}
