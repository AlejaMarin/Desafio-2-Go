package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/AlejaMarin/Desafio-2-Go/cmd/server/handler"
	"github.com/AlejaMarin/Desafio-2-Go/internal/dentista"
	"github.com/AlejaMarin/Desafio-2-Go/internal/paciente"
	"github.com/AlejaMarin/Desafio-2-Go/internal/turno"
	"github.com/AlejaMarin/Desafio-2-Go/pkg/middleware"
	"github.com/AlejaMarin/Desafio-2-Go/pkg/store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	database := os.Getenv("DATABASE")
	dataSourceName := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	storage := store.NewSqlStore(db)

	dentistRepo := dentista.NewRepository(storage)
	denstistService := dentista.NewService(dentistRepo)
	dentistHandler := handler.NewDentistHandler(denstistService)
	patientRepo := paciente.NewRepository(storage)
	patientService := paciente.NewService(patientRepo)
	patientHandler := handler.NewPatientHandler(patientService)
	shiftRepo := turno.NewRepository(storage)
	shiftService := turno.NewService(shiftRepo)
	shiftHandler := handler.NewShiftHandler(shiftService)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

	dentists := r.Group("/dentistas")
	{
		dentists.GET(":id", dentistHandler.GetByID())
		dentists.POST("", middleware.TokenAuthMiddleware(), dentistHandler.Post())
		dentists.PUT(":id", middleware.TokenAuthMiddleware(), dentistHandler.Put())
		dentists.PATCH(":id", middleware.TokenAuthMiddleware(), dentistHandler.Patch())
		dentists.DELETE(":id", middleware.TokenAuthMiddleware(), dentistHandler.Delete())
	}
	patients := r.Group("/pacientes")
	{
		patients.GET(":id", patientHandler.GetByID())
		patients.POST("", middleware.TokenAuthMiddleware(), patientHandler.Post())
		patients.PUT(":id", middleware.TokenAuthMiddleware(), patientHandler.Put())
		patients.PATCH(":id", middleware.TokenAuthMiddleware(), patientHandler.Patch())
		patients.DELETE(":id", middleware.TokenAuthMiddleware(), patientHandler.Delete())
	}
	shifts := r.Group("/turnos")
	{
		shifts.GET(":id", shiftHandler.GetByID())
		shifts.POST("", middleware.TokenAuthMiddleware(), shiftHandler.Post())
		shifts.PUT(":id", middleware.TokenAuthMiddleware(), shiftHandler.Put())
		shifts.PATCH(":id", middleware.TokenAuthMiddleware(), shiftHandler.Patch())
		shifts.DELETE(":id", middleware.TokenAuthMiddleware(), shiftHandler.Delete())
		shifts.POST("/pacientedentista", middleware.TokenAuthMiddleware(), shiftHandler.PostDos())
		shifts.GET("", shiftHandler.GetByDni())
	}

	r.Run(":8080")

}
