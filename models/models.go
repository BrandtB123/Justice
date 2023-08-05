package models

import "github.com/go-pg/pg"

var DB *pg.DB

func NewDB() {
	db := pg.Connect(&pg.Options{
		Addr:     ".cs96vtk3gnbe.us-east-2.rds.amazonaws.com:5432",
		User:     "postgres",
		Password: "",
		Database: "postgres",
	})
	DB = db
}
