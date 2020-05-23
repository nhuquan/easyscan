package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Document
type Document struct {
	Name         string `json:"name" xml:"name"`
	Type         string `json:"type" xml:"type"`
	Size         int64  `json:"size" xml:"size"`
	CreationDate string `json:"creationDate" xml:"creationDate"`
}

// Handler
func returnJson(c echo.Context) error {
	doc := &Document{
		Name:         "test",
		Type:         "Facture",
		Size:         1000,
		CreationDate: "2020",
	}
	return c.JSONPretty(http.StatusOK, doc, " ")
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	e.GET("/docs", returnJson)

	e.Logger.Fatal(e.Start(":1323"))
}
