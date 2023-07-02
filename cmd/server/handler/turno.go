package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/AlejaMarin/Desafio-2-Go/internal/domain"
	"github.com/AlejaMarin/Desafio-2-Go/internal/turno"
	"github.com/AlejaMarin/Desafio-2-Go/pkg/web"
	"github.com/gin-gonic/gin"
)

type shiftHandler struct {
	s turno.Service
}

type RequestTurno struct {
	IdPaciente  int    `json:"idPaciente,omitempty"`
	IdDentista  int    `json:"idDentista,omitempty"`
	Fecha       string `json:"fecha,omitempty"`
	Hora        string `json:"hora,omitempty"`
	Descripcion string `json:"descripcion,omitempty"`
}

func NewShiftHandler(s turno.Service) *shiftHandler {
	return &shiftHandler{
		s: s,
	}
}

// GetShiftById godoc
// @Summary Obtener turno
// @Description Obtener turno por ID
// @Tags Turnos
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} domain.Turno
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /turnos/{id} [get]
func (h *shiftHandler) GetByID() gin.HandlerFunc {

	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("Id Inválido"))
			return
		}
		t, err := h.s.GetShiftById(id)
		if err != nil {
			web.Failure(c, http.StatusNotFound, errors.New("El ID del Turno ingresado No Existe"))
			return
		}
		web.Success(c, http.StatusOK, t)
	}
}

func validateEmptysShift(shift *domain.Turno) (bool, error) {

	switch {
	case shift.Fecha == "" || shift.Hora == "" || shift.Descripcion == "":
		return false, errors.New("Los campos no pueden estar vacíos")
	case shift.IdPaciente <= 0 || shift.IdDentista <= 0:
		if shift.IdPaciente <= 0 {
			return false, errors.New("El ID del Paciente es inválido")
		}
		if shift.IdDentista <= 0 {
			return false, errors.New("El ID del Dentista es inválido")
		}
	}
	return true, nil
}

// SaveShift godoc
// @Summary Agregar turno
// @Description Agregar turno
// @Tags Turnos
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param turno body domain.Turno true "Body shift"
// @Success 201 {object} domain.Turno
// @Failure 400 {object} web.errorResponse
// @Router /turnos [post]
func (h *shiftHandler) Post() gin.HandlerFunc {

	return func(c *gin.Context) {
		var t domain.Turno
		err := c.ShouldBindJSON(&t)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("JSON Inválido"))
			return
		}
		valid, err := validateEmptysShift(&t)
		if !valid {
			web.Failure(c, http.StatusBadRequest, err)
			return
		}
		valid, err = validateDate(t.Fecha)
		if !valid {
			web.Failure(c, http.StatusBadRequest, err)
			return
		}
		shift, err := h.s.CreateShift(t)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, err)
			return
		}
		web.Success(c, http.StatusCreated, shift)
	}
}

// UpdateShift godoc
// @Summary Actualizar turno
// @Description Actualizar turno por ID
// @Tags Turnos
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param token header string true "token"
// @Param turno body domain.Turno true "Body shift"
// @Success 200 {object} domain.Turno
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Router /turnos/{id} [put]
func (h *shiftHandler) Put() gin.HandlerFunc {

	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("Id Inválido"))
			return
		}
		_, err = h.s.GetShiftById(id)
		if err != nil {
			web.Failure(c, http.StatusNotFound, errors.New("Turno No Encontrado"))
			return
		}
		var shift domain.Turno
		err = c.ShouldBindJSON(&shift)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("JSON Inválido"))
			return
		}
		valid, err := validateEmptysShift(&shift)
		if !valid {
			web.Failure(c, http.StatusBadRequest, err)
			return
		}
		valid, err = validateDate(shift.Fecha)
		if !valid {
			web.Failure(c, http.StatusBadRequest, err)
			return
		}
		t, err := h.s.UpdateShift(id, shift)
		if err != nil {
			web.Failure(c, http.StatusConflict, err)
			return
		}
		web.Success(c, http.StatusOK, t)
	}
}

// PartialUpdateShift godoc
// @Summary Actualizar turno
// @Description Actualizar un turno por alguno de sus campos
// @Tags Turnos
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param token header string true "token"
// @Param request body handler.RequestTurno true "Request body"
// @Success 200 {object} domain.Turno
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Router /turnos/{id} [patch]
func (h *shiftHandler) Patch() gin.HandlerFunc {

	return func(c *gin.Context) {
		var r RequestTurno
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("id inválido"))
			return
		}
		_, err = h.s.GetShiftById(id)
		if err != nil {
			web.Failure(c, http.StatusNotFound, errors.New("turno no encontrado"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("json inválido"))
			return
		}
		update := domain.Turno{
			IdPaciente:  r.IdPaciente,
			IdDentista:  r.IdDentista,
			Fecha:       r.Fecha,
			Hora:        r.Hora,
			Descripcion: r.Descripcion,
		}
		if update.Fecha != "" {
			valid, err := validateDate(update.Fecha)
			if !valid {
				web.Failure(c, http.StatusBadRequest, err)
				return
			}
		}
		/*
			if update.IdDentista != 0 || update.IdPaciente <= 0 {
				web.Failure(c, http.StatusBadRequest, err)
				return
			}
		*/
		t, err := h.s.UpdateShift(id, update)
		if err != nil {
			web.Failure(c, http.StatusConflict, err)
			return
		}
		web.Success(c, http.StatusOK, t)
	}

}

// DeleteTurno godoc
// @Summary Eliminar turno
// @Description Eliminar un turno por ID
// @Tags Turnos
// @Param id path int true "id"
// @Param token header string true "token"
// @Success 204
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /turnos/{id} [delete]
func (h *shiftHandler) Delete() gin.HandlerFunc {

	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("Id Inválido"))
			return
		}
		err = h.s.DeleteShift(id)

		if err != nil {
			web.Failure(c, http.StatusNotFound, err)
			return
		}
		web.Success(c, http.StatusNoContent, nil)
	}
}

// SaveShiftTwo godoc
// @Summary Agregar turno
// @Description Agregar turno por DNI del paciente y matrícula del dentista
// @Tags Turnos
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param turno body domain.TurnoDos true "Body shift"
// @Success 201 {object} domain.Turno
// @Failure 400 {object} web.errorResponse
// @Router /turnos/pacientedentista [post]
func (h *shiftHandler) PostDos() gin.HandlerFunc {

	return func(c *gin.Context) {
		var t domain.TurnoDos
		err := c.ShouldBindJSON(&t)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("JSON Inválido"))
			return
		}
		valid, err := validateEmptysShiftTwo(&t)
		if !valid {
			web.Failure(c, http.StatusBadRequest, err)
			return
		}
		valid, err = validateDate(t.Fecha)
		if !valid {
			web.Failure(c, http.StatusBadRequest, err)
			return
		}
		shift, err := h.s.CreateShiftByDniAndEnrollment(t)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, err)
			return
		}
		web.Success(c, http.StatusCreated, shift)
	}
}

func validateEmptysShiftTwo(shift *domain.TurnoDos) (bool, error) {

	if shift.DniPaciente == "" || shift.MatriculaDentista == "" || shift.Fecha == "" || shift.Hora == "" || shift.Descripcion == "" {
		return false, errors.New("Los campos no pueden estar vacíos")
	}

	return true, nil
}

// GetShiftByDni godoc
// @Summary Obtener turno
// @Description Obtener turno por DNI del paciente
// @Tags Turnos
// @Produce json
// @Param dni query string true "dni"
// @Success 200 {array} domain.TurnoByDni
// @Failure 400 {object} web.errorResponse
// @Router /turnos [get]
func (h *shiftHandler) GetByDni() gin.HandlerFunc {
	return func(c *gin.Context) {
		dni := c.Query("dni")
		turnos, err := h.s.GetShiftsByDniPatient(dni)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, err)
			return
		}
		web.Success(c, http.StatusOK, turnos)
	}
}
