package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Repo interface {
	GetOne(id string) (interface{}, error)
	GetAll() ([]interface{}, error)
	AddOne(item interface{}) (interface{}, error)
	Update(id string, item interface{}) (interface{}, error)
}

type Doc struct{}

func (d Doc) GetOne(id string) (interface{}, error) {
	// TODO
	db, err := sql.Open("sqlite3", "./easyscan_2020_05.db")
	if err != nil {
		return nil, err
	}
	q := "Select * from documents where id=?"
	rows, err := db.Query(q, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var id int
	var name string
	var doctype string
	var creationDate string
	if rows.Next() {
		err = rows.Scan(&id, &name, &doctype, &creationDate)
		if err != nil {
			return nil, err
		}
		fmt.Println(id)
		fmt.Println(name)
		fmt.Println(doctype)
		fmt.Println(creationDate)
	}

	return nil, nil
}

func (d Doc) GetAll() (interface{}, error) {
	// TODO
	return nil, nil
}

func (d Doc) AddOne(item interface{}) (interface{}, error) {
	// TODO
	return nil, nil
}

func (d Doc) Update(id string, item interface{}) (interface{}, error) {
	// TODO
	return nil, nil
}
