package dentista

import "github.com/AlejaMarin/Desafio-2-Go/internal/domain"

type Service interface {
	GetDentistById(id int) (domain.Dentista, error)
	CreateDentist(d domain.Dentista) (domain.Dentista, error)
	UpdateDentist(id int, d domain.Dentista) (domain.Dentista, error)
	DeleteDentist(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

// CreateDentist implements Service.
func (s *service) CreateDentist(d domain.Dentista) (domain.Dentista, error) {
	d, err := s.r.CreateDentist(d)
	if err != nil {
		return domain.Dentista{}, err
	}
	return d, nil
}

func (s *service) GetDentistById(id int) (domain.Dentista, error) {

	d, err := s.r.GetDentistById(id)
	if err != nil {
		return domain.Dentista{}, err
	}
	return d, nil
}

func (s *service) UpdateDentist(id int, d domain.Dentista) (domain.Dentista, error) {
	dentist, err := s.r.GetDentistById(id)

	if err != nil {
		return domain.Dentista{}, err
	}

	if d.Apellido != "" {
		dentist.Apellido = d.Apellido
	}

	if d.Nombre != "" {
		dentist.Nombre = d.Nombre
	}

	if d.Matricula != "" {
		dentist.Matricula = d.Matricula
	}

	d, err = s.r.UpdateDentist(id, dentist)
	if err != nil {
		return domain.Dentista{}, err
	}

	return d, nil
}

func (s *service) DeleteDentist(id int) error {
	err := s.r.DeleteDentist(id)
	if err != nil {
		return err
	}
	return nil
}
