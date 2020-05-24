package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

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

func upload(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	d := docs[id]
	if d == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	// Read form field
	name := c.FormValue("name")
	docType := c.FormValue("type")
	creationDate := time.Now().String()

	// update metadata
	d.Type = docType
	d.CreationDate = creationDate

	// Read file

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, fmt.Sprintf("File %s is uploaded successfully with fields name=%s and email=%s.", file.Filename, name), " ")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Routes
	e.GET("/docs", getAllDocuments)
	e.GET("/docs/:id", getDocument)
	e.POST("/docs/new", createDocument)
	e.POST("/docs/:id/upload", upload)
	e.PUT("/docs/:id", updateDocument)
	e.DELETE("/docs/:id", deleteDocument)

	e.Logger.Fatal(e.Start(":1323"))
}
