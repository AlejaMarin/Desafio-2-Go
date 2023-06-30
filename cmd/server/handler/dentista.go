package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/AlejaMarin/Desafio-2-Go/internal/dentista"
	"github.com/AlejaMarin/Desafio-2-Go/internal/domain"
	"github.com/AlejaMarin/Desafio-2-Go/pkg/web"
	"github.com/gin-gonic/gin"
)

type dentistHandler struct {
	s dentista.Service
}

func NewDentistHandler(s dentista.Service) *dentistHandler {
	return &dentistHandler{
		s: s,
	}
}

func (h *dentistHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("id inválido"))
			return
		}
		d, err := h.s.GetDentistById(id)
		if err != nil {
			web.Failure(c, http.StatusNotFound, errors.New("el id ingresado no existe"))
			return
		}
		web.Success(c, http.StatusOK, d)
	}
}

func validateEmptys(dentist *domain.Dentista) (bool, error) {
	if dentist.Apellido == "" || dentist.Nombre == "" || dentist.Matricula == "" {
		return false, errors.New("todos los campos deben estar completos")
	}
	return true, nil
}

func (h *dentistHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var d domain.Dentista
		err := c.ShouldBindJSON(&d)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("json inválido"))
			return
		}
		valid, err := validateEmptys(&d)
		if !valid {
			web.Failure(c, http.StatusBadRequest, err)
			return
		}
		dentist, err := h.s.CreateDentist(d)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, err)
			return
		}
		web.Success(c, http.StatusCreated, dentist)
	}
}

func (h *dentistHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("id inválido"))
			return
		}
		_, err = h.s.GetDentistById(id)
		if err != nil {
			web.Failure(c, http.StatusNotFound, errors.New("odontólogo no encontrado"))
			return
		}
		var dentist domain.Dentista
		err = c.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("JSON Inválido"))
			return
		}
		valid, err := validateEmptys(&dentist)
		if !valid {
			web.Failure(c, http.StatusBadRequest, err)
			return
		}
		p, err := h.s.UpdateDentist(id, dentist)
		if err != nil {
			web.Failure(c, http.StatusConflict, err)
			return
		}
		web.Success(c, http.StatusOK, p)
	}
}

func (h *dentistHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Apellido  string `json:"apellido,omitempty"`
		Nombre    string `json:"nombre,omitempty"`
		Matricula string `json:"matricula,omitempty"`
	}

	return func(c *gin.Context) {
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("id inválido"))
			return
		}
		_, err = h.s.GetDentistById(id)
		if err != nil {
			web.Failure(c, http.StatusNotFound, errors.New("odontólogo no encontrado"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("json inválido"))
			return
		}
		update := domain.Dentista{
			Apellido:  r.Apellido,
			Nombre:    r.Nombre,
			Matricula: r.Matricula,
		}
		p, err := h.s.UpdateDentist(id, update)
		if err != nil {
			web.Failure(c, http.StatusConflict, err)
			return
		}
		web.Success(c, http.StatusOK, p)
	}
}

func (h *dentistHandler) Delete() gin.HandlerFunc {

	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("id inválido"))
			return
		}
		err = h.s.DeleteDentist(id)
		if err != nil {
			web.Failure(c, http.StatusNotFound, err)
			return
		}
		web.Success(c, http.StatusNoContent, nil)
	}
}
