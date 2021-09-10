package main

import "fmt"

type Pusa struct {
	Name     string
	Laziness int
}

type Database struct {
	Puso []Pusa
}

func NewDataBase() *Database {
	return &Database{
		Puso: make([]Pusa, 0),
	}
}

func (db *Database) Save(p Pusa) {
	db.Puso = append(db.Puso, p)
}

func (db *Database) Dump() []Pusa {
	fmt.Printf("db.Puso: %v\n", db.Puso)
	return db.Puso
}
