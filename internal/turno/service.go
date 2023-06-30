package turno

import "github.com/AlejaMarin/Desafio-2-Go/internal/domain"

type Service interface {
	GetShiftById(id int) (domain.Turno, error)
	CreateShift(t domain.Turno) (domain.Turno, error)
	UpdateShift(id int, x domain.Turno) (domain.Turno, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetShiftById(id int) (domain.Turno, error) {

	t, err := s.r.GetShiftById(id)
	if err != nil {
		return domain.Turno{}, err
	}
	return t, nil
}

func (s *service) CreateShift(t domain.Turno) (domain.Turno, error) {

	t, err := s.r.CreateShift(t)
	if err != nil {
		return domain.Turno{}, err
	}
	return t, nil
}

func (s *service) UpdateShift(id int, x domain.Turno) (domain.Turno, error) {

	shift, err := s.r.GetShiftById(id)
	if err != nil {
		return domain.Turno{}, err
	}
	if x.IdPaciente > 0 {
		shift.IdPaciente = x.IdPaciente
	}
	if x.IdDentista > 0 {
		shift.IdDentista = x.IdDentista
	}
	if x.Fecha != "" {
		shift.Fecha = x.Fecha
	}
	if x.Hora != "" {
		shift.Hora = x.Hora
	}
	if x.Descripcion != "" {
		shift.Descripcion = x.Descripcion
	}
	t, err := s.r.UpdateShift(id, shift)
	if err != nil {
		return domain.Turno{}, err
	}
	return t, nil
}
