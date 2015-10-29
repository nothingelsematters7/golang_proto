package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"

	"github.com/nothingelsematters7/golang_proto/config"
)

type Email struct {
	ID         int
	UserID     int     `sql:"index"`                          // Foreign key (belongs to), tag `index` will create index for this field when using AutoMigrate
	Email      string  `sql:"type:varchar(100);unique_index"` // Set field's sql type, tag `unique_index` will create unique index
	Subscribed bool
}

func main() {
	db, err := gorm.Open("mysql", config.Conf.MysqlArgs())
	if err != nil {
		log.Fatal("cannot connect to database...")
	}
	db.DB()

	db.CreateTable(&Email{})
}
