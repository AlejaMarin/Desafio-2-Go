package store

import (
	"database/sql"
	"log"

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

	q := "SELECT MAX(id) FROM paciente;"
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

func (s *sqlStore) GetShiftById(id int) (domain.Turno, error) {

	var shift domain.Turno

	query := "SELECT t.id, t.idPaciente, t.idDentista, t.fecha, t.hora, t.descripcion FROM turno t LEFT JOIN paciente p ON p.id = t.idPaciente LEFT JOIN dentista d ON d.id = t.idDentista WHERE t.id = ? AND p.activo = ? AND d.activo = ?;"
	row := s.DB.QueryRow(query, id, true, true)
	err := row.Scan(&shift.Id, &shift.IdPaciente, &shift.IdDentista, &shift.Fecha, &shift.Hora, &shift.Descripcion)
	if err != nil {
		return domain.Turno{}, err
	}
	return shift, nil

}

func (s *sqlStore) CreateShift(t domain.Turno) (int, error) {

	query := "INSERT INTO turno(idPaciente, idDentista, fecha, hora, descripcion) VALUES(?, ?, ?, ?, ?)"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(t.IdPaciente, t.IdDentista, t.Fecha, t.Hora, t.Descripcion)
	if err != nil {
		return 0, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return 0, err
	}

	var id int

	q := "SELECT MAX(id) FROM turno;"
	row := s.DB.QueryRow(q)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (s *sqlStore) UpdateShift(t domain.Turno) error {

	query := "UPDATE turno SET idPaciente = ?, idDentista = ?, fecha = ?, hora = ?, descripcion = ? WHERE id = ?;"

	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(t.IdPaciente, t.IdDentista, t.Fecha, t.Hora, t.Descripcion, t.Id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil

}

func (s *sqlStore) ExistsShift(f, h string, idD int) bool {

	query := "SELECT id FROM turno WHERE fecha = ? AND hora = ? AND idDentista = ?"
	row := s.DB.QueryRow(query, f, h, idD)
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

func (s *sqlStore) DeleteShift(id int) error {
	query := "DELETE FROM turno WHERE id = ?"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil

}

func (s *sqlStore) GetPatientIdByDni(dni string) (int, error) {
	query := "SELECT id FROM paciente WHERE dni = ?;"
	row := s.DB.QueryRow(query, dni)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *sqlStore) GetDentistByMatricula(matricula string) (int, error) {
	query := "SELECT id FROM dentista WHERE matricula = ?;"
	row := s.DB.QueryRow(query, matricula)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *sqlStore) GetShiftsByDniPatient(dni string) ([]domain.TurnoByDni, error) {
	query := "SELECT t.fecha, t.hora, t.descripcion, p.id, p.nombre, p.apellido, p.domicilio, p.dni, d.id, d.apellido, d.nombre, d.matricula FROM turno t LEFT JOIN paciente p ON p.id = t.idPaciente LEFT JOIN dentista d ON d.id = t.idDentista WHERE p.dni = ?;"
	rows, err := s.DB.Query(query, dni)
	if err != nil {
		return nil, err
	}
	var t domain.TurnoByDni
	var turnos []domain.TurnoByDni

	for rows.Next() {
		err := rows.Scan(&t.Fecha, &t.Hora, &t.Descripcion, &t.Paciente.Id, &t.Paciente.Nombre, &t.Paciente.Apellido, &t.Paciente.Domicilio, &t.Paciente.DNI, &t.Dentista.Id, &t.Dentista.Apellido, &t.Dentista.Nombre, &t.Dentista.Matricula)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		} else {
			turnos = append(turnos, t)
		}
	}
	return turnos, nil
}
