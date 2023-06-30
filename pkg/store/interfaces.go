package store

import "github.com/AlejaMarin/Desafio-2-Go/internal/domain"

type StoreInterface interface {
	GetDentistById(id int) (domain.Dentista, error)
	CreateDentist(d domain.Dentista) (int, error)
	UpdateDentist(d domain.Dentista) error
	DeleteDentist(id int) error
	ExistsDentistByMatricula(Matricula string) bool
}