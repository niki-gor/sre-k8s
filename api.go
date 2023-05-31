package main

import (
	"database/sql"
	"fmt"
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
	animals := os.Getenv("ANIMALS")
	if animals == "" {
		log.Fatal("ANIMALS environment variable not set")
	}

	db, err := sql.Open("postgres", dsn+"?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	e := echo.New()

	route := fmt.Sprintf("/%s/:kind", animals)
	e.GET(route, func(c echo.Context) error {
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
