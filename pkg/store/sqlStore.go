package store

import (
	"database/sql"

	"github.com/AlejaMarin/Desafio-2-Go/internal/domain"
)

type sqlStore struct {
	DB *sql.DB
}

func NewSqlStore(database *sql.DB) StoreInterface {
	return &sqlStore{
		DB: database,
	}
}

func (s *sqlStore) GetPatientById(id int) (domain.Paciente, error) {

	var patient domain.Paciente

	query := "SELECT id, nombre, apellido, domicilio, dni, fechaAlta FROM paciente WHERE id = ? AND activo = ?;"
	row := s.DB.QueryRow(query, id, true)
	err := row.Scan(&patient.Id, &patient.Nombre, &patient.Apellido, &patient.Domicilio, &patient.DNI, &patient.FechaAlta)
	if err != nil {
		return domain.Paciente{}, err
	}
	return patient, nil

}

func (s *sqlStore) CreatePatient(p domain.Paciente) (int, error) {

	query := "INSERT INTO paciente(nombre, apellido, domicilio, dni, fechaAlta, activo) VALUES(?, ?, ?, ?, ?, ?)"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(p.Nombre, p.Apellido, p.Domicilio, p.DNI, p.FechaAlta, true)
	if err != nil {
		return 0, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return 0, err
	}

	var id int

	q := "SELECT LAST_INSERT_ID();"
	row := s.DB.QueryRow(q)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (s *sqlStore) UpdatePatient(p domain.Paciente) error {

	query := "UPDATE paciente SET nombre = ?, apellido = ?, domicilio = ?, dni = ?, fechaAlta = ? WHERE id = ? AND activo = ?;"

	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(p.Nombre, p.Apellido, p.Domicilio, p.DNI, p.FechaAlta, p.Id, true)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil

}

func (s *sqlStore) DeletePatient(id int) error {

	query := "UPDATE paciente SET activo = ? WHERE id = ?"
	_, err := s.DB.Exec(query, false, id)
	if err != nil {
		return err
	}
	return nil

}

func (s *sqlStore) ExistsPatientByDNI(DNI string) bool {

	query := "SELECT id FROM paciente WHERE dni = ?"
	row := s.DB.QueryRow(query, DNI)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return false
	}

	if id > 0 {
		return true
	}
	return false
}

func (s *sqlStore) GetDentistById(id int) (domain.Dentista, error) {

	var dentist domain.Dentista

	query := "SELECT id, apellido, nombre, matricula FROM dentista WHERE id = ? AND activo = ?;"
	row := s.DB.QueryRow(query, id, true)
	err := row.Scan(&dentist.Id, &dentist.Apellido, &dentist.Nombre, &dentist.Matricula)
	if err != nil {
		return domain.Dentista{}, err
	}
	return dentist, nil

}

func (s *sqlStore) CreateDentist(d domain.Dentista) (int, error) {

	query := "INSERT INTO dentista(apellido, nombre, matricula, activo) VALUES(?, ?, ?, ?)"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(d.Apellido, d.Nombre, d.Matricula, true)
	if err != nil {
		return 0, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return 0, err
	}

	var id int

	q := "SELECT MAX(id) FROM dentista;"
	row := s.DB.QueryRow(q)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (s *sqlStore) UpdateDentist(d domain.Dentista) error {

	query := "UPDATE dentista SET apellido = ?, nombre = ?, matricula = ? WHERE id = ? AND activo = ?;"

	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(d.Apellido, d.Nombre, d.Matricula, d.Id, true)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil

}

func (s *sqlStore) DeleteDentist(id int) error {

	query := "UPDATE dentista SET activo = ? WHERE id = ?"
	_, err := s.DB.Exec(query, false, id)
	if err != nil {
		return err
	}
	return nil

}

func (s *sqlStore) ExistsDentistByMatricula(Matricula string) bool {

	query := "SELECT id FROM dentista WHERE matricula = ?"
	row := s.DB.QueryRow(query, Matricula)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return false
	}

	if id > 0 {
		return true
	}
	return false
}
