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

func NewShiftHandler(s turno.Service) *shiftHandler {
	return &shiftHandler{
		s: s,
	}
}

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
