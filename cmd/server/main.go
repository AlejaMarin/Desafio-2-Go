package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/AlejaMarin/Desafio-2-Go/cmd/server/handler"
	"github.com/AlejaMarin/Desafio-2-Go/internal/dentista"
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

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })
	dentists := r.Group("/dentistas")
	{
		dentists.GET(":id", dentistHandler.GetByID())
		dentists.POST("", dentistHandler.Post())
		dentists.PUT(":id", dentistHandler.Put())
		dentists.PATCH(":id", dentistHandler.Patch())
		dentists.DELETE(":id", dentistHandler.Delete())
	}

	r.Run(":8080")

}
