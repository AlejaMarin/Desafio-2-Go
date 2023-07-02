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

type Request struct {
	Apellido  string `json:"apellido,omitempty"`
	Nombre    string `json:"nombre,omitempty"`
	Matricula string `json:"matricula,omitempty"`
}

func NewDentistHandler(s dentista.Service) *dentistHandler {
	return &dentistHandler{
		s: s,
	}
}

// GetDentistById godoc
// @Summary Obtener dentista
// @Description Obtener dentista por ID
// @Tags Dentistas
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} domain.Dentista 
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /dentistas/{id} [get]
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

func validateEmptysDentist(dentist *domain.Dentista) (bool, error) {
	if dentist.Apellido == "" || dentist.Nombre == "" || dentist.Matricula == "" {
		return false, errors.New("todos los campos deben estar completos")
	}
	return true, nil
}

// SaveDentist godoc
// @Summary Agregar dentista
// @Description Crear dentista
// @Tags Dentistas
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param dentista body domain.Dentista true "Body dentist"
// @Success 201 {object} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Router /dentistas [post]
func (h *dentistHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var d domain.Dentista
		err := c.ShouldBindJSON(&d)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("json inválido"))
			return
		}
		valid, err := validateEmptysDentist(&d)
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

// UpdateDentist godoc
// @Summary Actualizar dentista
// @Description Actualizar dentista por ID
// @Tags Dentistas
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param token header string true "token"
// @Param dentista body domain.Dentista true "Body dentist"
// @Success 200 {object} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Router /dentistas/{id} [put]
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
		valid, err := validateEmptysDentist(&dentist)
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

// PartialUpdateDentist godoc
// @Summary Actualizar dentista
// @Description Actualizar un dentista por alguno de sus campos
// @Tags Dentistas
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param token header string true "token"
// @Param request body handler.Request true "Request body"
// @Success 200 {object} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Router /dentistas/{id} [patch]
func (h *dentistHandler) Patch() gin.HandlerFunc {
	
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

// DeleteDentist godoc
// @Summary Eliminar dentista
// @Description Eliminar un dentista por ID
// @Tags Dentistas
// @Param id path int true "id"
// @Param token header string true "token"
// @Success 204
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /dentistas/{id} [delete]
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
