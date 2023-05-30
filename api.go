package main

import (
	"database/sql"
	"net/http"

	"log"
	"os"

	"github.com/labstack/echo/v4"

	_ "github.com/lib/pq"
)

func main() {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DSN environment variable not set")
	}

	db, err := sql.Open("postgres", dsn+"?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	e := echo.New()

	e.GET("/:kind", func(c echo.Context) error {
		kind := c.Param("kind")
		var info string
		err := db.QueryRow("SELECT info FROM mytable WHERE kind = $1", kind).Scan(&info)
		switch err {
		case nil:
			return c.String(http.StatusOK, info)
		case sql.ErrNoRows:
			return echo.NewHTTPError(http.StatusNotFound, "Такую породу мы не нашли")
		default:
			log.Println(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "На сервере что-то пошло не так(")
		}
	})

	e.Logger.Fatal(e.Start(":8081"))
}
