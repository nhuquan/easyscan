package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"strconv"
	"time"
	_ "time"

	_ "github.com/mattn/go-sqlite3"
	"misoda.fr/easyscan/pkg/domain"
)

var dbEngine = "sqlite3"
var dbPath = "./easyscan_2020_05.db"

func init() {
	var currentMonth string
	log.Info("Initializing database repositories")
	currentMonth = strconv.Itoa(time.Now().Year()) + "_" + time.Now().Month().String()
	log.Info("Current date is: %s", currentMonth)
	// TODO check if database file is exist for current month, if not create it:
}

type Repo interface {
	GetOne(id string) (interface{}, error)
	GetAll() ([]interface{}, error)
	AddOne(item interface{}) (interface{}, error)
	Update(id string, item interface{}) (interface{}, error)
}

type DocRepo struct{}

func (d DocRepo) GetOne(id string) (interface{}, error) {
	db, err := sql.Open(dbEngine, dbPath)
	if err != nil {
		return nil, err
	}
	q := "SELECT id, name, docType, creationDate FROM documents WHERE id=?"
	rows, err := db.Query(q, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var docID int64
	var name string
	var doctype string
	var creationDate string
	var dd *domain.Document
	if rows.Next() {
		err = rows.Scan(&docID, &name, &doctype, &creationDate)
		if err != nil {
			return nil, err
		}
		dd = &domain.Document{
			ID:           docID,
			DocType:      doctype,
			Name:         name,
			CreationDate: creationDate,
		}
		fmt.Println(docID)
		fmt.Println(name)
		fmt.Println(doctype)
		fmt.Println(creationDate)
	}

	return dd, nil
}

func (d DocRepo) GetAll() (interface{}, error) {
	return nil, nil
}

func (d DocRepo) AddOne(item interface{}) (interface{}, error) {
	// Check if the item is of type Document
	dd, ok := item.(*domain.Document)
	if !ok {
		return nil, errors.New("not a document")
	}

	db, err := sql.Open(dbEngine, dbPath)
	checkerr(err)

	q := "INSERT INTO documents(docType, name, creationDate) values (?, ?, ?)"
	stmt, err := db.Prepare(q)
	checkerr(err)

	res, err := stmt.Exec(dd.DocType, dd.Name, dd.CreationDate)
	checkerr(err)

	id, err := res.LastInsertId()
	checkerr(err)

	dd.ID = id

	return dd, nil
}

func (d DocRepo) Update(id string, item interface{}) (interface{}, error) {
	return nil, nil
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
