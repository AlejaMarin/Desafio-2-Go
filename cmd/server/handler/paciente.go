package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/AlejaMarin/Desafio-2-Go/internal/domain"
	"github.com/AlejaMarin/Desafio-2-Go/internal/paciente"
	"github.com/AlejaMarin/Desafio-2-Go/pkg/web"
	"github.com/gin-gonic/gin"
)

type patientHandler struct {
	s paciente.Service
}

type RequestPaciente struct {
	Nombre    string `json:"nombre,omitempty"`
	Apellido  string `json:"apellido,omitempty"`
	Domicilio string `json:"domicilio,omitempty"`
	DNI       string `json:"dni,omitempty"`
	FechaAlta string `json:"fechaAlta,omitempty"`
}

func NewPatientHandler(s paciente.Service) *patientHandler {
	return &patientHandler{
		s: s,
	}
}

// GetPatientById godoc
// @Summary Obtener paciente
// @Description Obtener paciente por ID
// @Tags Pacientes
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} domain.Paciente
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /pacientes/{id} [get]
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
	//list := []int{}
	if len(dates) != 3 {
		return false, errors.New("Fecha Inválida, debe tener Formato: dd/mm/yyyy")
	}
	for value := range dates {
		_, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("Fecha Inválida, deben ser Números")
		}
		//list = append(list, number)
	}
	/* condition := (list[0] < 1 || list[0] > 31) && (list[1] < 1 || list[1] > 12) && (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("Fecha de Alta Inválida, debe estar entre 01/01/0001 y 31/12/9999")
	} */
	validDate := dates[2] + "-" + dates[1] + "-" + dates[0]
	_, err := time.Parse("2006-01-02", validDate)
	if err != nil {
		return false, errors.New("Fecha Inválida |" + err.Error())
	}
	return true, nil
}

// SavePatient godoc
// @Summary Agregar paciente
// @Description Crear paciente
// @Tags Pacientes
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param paciente body domain.Paciente true "Body patient"
// @Success 201 {object} domain.Paciente
// @Failure 400 {object} web.errorResponse
// @Router /pacientes [post]
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

// UpdatePatient godoc
// @Summary Actualizar paciente
// @Description Actualizar paciente por ID
// @Tags Pacientes
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param token header string true "token"
// @Param dentista body domain.Paciente true "Body patient"
// @Success 200 {object} domain.Paciente
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Router /pacientes/{id} [put]
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

// PartialUpdatePatient godoc
// @Summary Actualizar paciente
// @Description Actualizar un paciente por alguno de sus campos
// @Tags Pacientes
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param token header string true "token"
// @Param request body handler.RequestPaciente true "Request body"
// @Success 200 {object} domain.Paciente
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Router /pacientes/{id} [patch]
func (h *patientHandler) Patch() gin.HandlerFunc {

	return func(c *gin.Context) {
		var r RequestPaciente
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

// DeletePatient godoc
// @Summary Eliminar paciente
// @Description Eliminar un paciente por ID
// @Tags Pacientes
// @Param id path int true "id"
// @Param token header string true "token"
// @Success 204
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /pacientes/{id} [delete]
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
