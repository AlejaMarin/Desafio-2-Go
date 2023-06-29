package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/AlejaMarin/Desafio-2-Go/internal/domain"
	"github.com/AlejaMarin/Desafio-2-Go/internal/paciente"
	"github.com/AlejaMarin/Desafio-2-Go/pkg/web"
	"github.com/gin-gonic/gin"
)

type patientHandler struct {
	s paciente.Service
}

func NewPatientHandler(s paciente.Service) *patientHandler {
	return &patientHandler{
		s: s,
	}
}

func (h *patientHandler) GetByID() gin.HandlerFunc {

	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("Id Inválido"))
			return
		}
		p, err := h.s.GetPatientById(id)
		if err != nil {
			web.Failure(c, http.StatusNotFound, errors.New("El ID del Paciente ingresado No Existe"))
			return
		}
		web.Success(c, http.StatusOK, p)
	}
}

func validateEmptys(patient *domain.Paciente) (bool, error) {

	if patient.Nombre == "" || patient.Apellido == "" || patient.Domicilio == "" || patient.DNI == "" || patient.FechaAlta == "" {
		return false, errors.New("Los campos no pueden estar vacíos")
	}
	return true, nil
}

func validateDate(date string) (bool, error) {

	dates := strings.Split(date, "/")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("Fecha de Alta Inválida, debe tener Formato: dd/mm/yyyy")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("Fecha de Alta Inválida, deben ser Números")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 31) && (list[1] < 1 || list[1] > 12) && (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("Fecha de Alta Inválida, debe estar entre 01/01/0001 y 31/12/9999")
	}
	return true, nil
}

func (h *patientHandler) Post() gin.HandlerFunc {

	return func(c *gin.Context) {
		var p domain.Paciente
		err := c.ShouldBindJSON(&p)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("JSON Inválido"))
			return
		}
		valid, err := validateEmptys(&p)
		if !valid {
			web.Failure(c, http.StatusBadRequest, err)
			return
		}
		valid, err = validateDate(p.FechaAlta)
		if !valid {
			web.Failure(c, http.StatusBadRequest, err)
			return
		}
		patient, err := h.s.CreatePatient(p)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, err)
			return
		}
		web.Success(c, http.StatusCreated, patient)
	}
}

func (h *patientHandler) Put() gin.HandlerFunc {

	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("Id Inválido"))
			return
		}
		_, err = h.s.GetPatientById(id)
		if err != nil {
			web.Failure(c, http.StatusNotFound, errors.New("Paciente No Encontrado"))
			return
		}
		var patient domain.Paciente
		err = c.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("JSON Inválido"))
			return
		}
		valid, err := validateEmptys(&patient)
		if !valid {
			web.Failure(c, http.StatusBadRequest, err)
			return
		}
		valid, err = validateDate(patient.FechaAlta)
		if !valid {
			web.Failure(c, http.StatusBadRequest, err)
			return
		}
		p, err := h.s.UpdatePatient(id, patient)
		if err != nil {
			web.Failure(c, http.StatusConflict, err)
			return
		}
		web.Success(c, http.StatusOK, p)
	}
}

func (h *patientHandler) Patch() gin.HandlerFunc {

	type Request struct {
		Nombre    string `json:"nombre,omitempty"`
		Apellido  string `json:"apellido,omitempty"`
		Domicilio string `json:"domicilio,omitempty"`
		DNI       string `json:"dni,omitempty"`
		FechaAlta string `json:"fechaAlta,omitempty"`
	}

	return func(c *gin.Context) {
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("Id Inválido"))
			return
		}
		_, err = h.s.GetPatientById(id)
		if err != nil {
			web.Failure(c, http.StatusNotFound, errors.New("Paciente No Encontrado"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("JSON Inválido"))
			return
		}
		update := domain.Paciente{
			Nombre:    r.Nombre,
			Apellido:  r.Apellido,
			Domicilio: r.Domicilio,
			DNI:       r.DNI,
			FechaAlta: r.FechaAlta,
		}
		if update.FechaAlta != "" {
			valid, err := validateDate(update.FechaAlta)
			if !valid {
				web.Failure(c, http.StatusBadRequest, err)
				return
			}
		}
		p, err := h.s.UpdatePatient(id, update)
		if err != nil {
			web.Failure(c, http.StatusConflict, err)
			return
		}
		web.Success(c, http.StatusOK, p)
	}
}

func (h *patientHandler) Delete() gin.HandlerFunc {

	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("Id Inválido"))
			return
		}
		err = h.s.DeletePatient(id)
		if err != nil {
			web.Failure(c, http.StatusNotFound, err)
			return
		}
		web.Success(c, http.StatusNoContent, nil)
	}
}
