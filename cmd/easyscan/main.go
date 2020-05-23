package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Document
type Document struct {
	ID           int    `json:"id", xml:"id"`
	Name         string `json:"name" xml:"name"`
	Type         string `json:"type" xml:"type"`
	Size         int    `json:"size" xml:"size"`
	CreationDate string `json:"creationDate" xml:"creationDate"`
}

var (
	docs = map[int]*Document{}
	seq  = 1
)

//----------
// Handler
//----------

func createDocument(c echo.Context) error {
	d := &Document{
		ID: seq,
	}
	if err := c.Bind(d); err != nil {
		return err
	}
	docs[d.ID] = d
	seq++
	return c.JSONPretty(http.StatusCreated, d, " ")
}

func getDocument(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSONPretty(http.StatusOK, docs[id], " ")
}

func getAllDocuments(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, docs, " ")
}

func updateDocument(c echo.Context) error {
	d := new(Document)
	if err := c.Bind(d); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	docs[id].Name = d.Name
	return c.JSONPretty(http.StatusOK, docs[id], " ")
}

func deleteDocument(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(docs, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/docs", getAllDocuments)
	e.POST("/docs", createDocument)
	e.GET("/docs/:id", getDocument)
	e.PUT("/docs/:id", updateDocument)
	e.DELETE("/docs/:id", deleteDocument)

	e.Logger.Fatal(e.Start(":1323"))
}
