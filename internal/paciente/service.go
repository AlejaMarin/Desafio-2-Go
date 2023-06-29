package paciente

import "github.com/AlejaMarin/Desafio-2-Go/internal/domain"

type Service interface {
	GetPatientById(id int) (domain.Paciente, error)
	CreatePatient(p domain.Paciente) (domain.Paciente, error)
	UpdatePatient(id int, p domain.Paciente) (domain.Paciente, error)
	DeletePatient(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetPatientById(id int) (domain.Paciente, error) {

	p, err := s.r.GetPatientById(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil

}
func (s *service) CreatePatient(p domain.Paciente) (domain.Paciente, error) {

	p, err := s.r.CreatePatient(p)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil

}
func (s *service) UpdatePatient(id int, p domain.Paciente) (domain.Paciente, error) {

	patient, err := s.r.GetPatientById(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	if p.Nombre != "" {
		patient.Nombre = p.Nombre
	}
	if p.Apellido != "" {
		patient.Apellido = p.Apellido
	}
	if p.Domicilio != "" {
		patient.Domicilio = p.Domicilio
	}
	if p.DNI != "" {
		patient.DNI = p.DNI
	}
	if p.FechaAlta != "" {
		patient.FechaAlta = p.FechaAlta
	}
	p, err = s.r.UpdatePatient(id, patient)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil

}
func (s *service) DeletePatient(id int) error {

	err := s.r.DeletePatient(id)
	if err != nil {
		return err
	}
	return nil

}
