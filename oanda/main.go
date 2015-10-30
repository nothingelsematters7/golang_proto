package main

import (
	"io/ioutil"
	"database/sql"
	"time"
	"log"
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/nothingelsematters7/golang_proto/config"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/claudiu/gocron"
	"fmt"
)

type Response struct {
	BaseCurrency string `json:"base_currency"`
	Quotes       map[string]Quote `json:"quotes"`
}

type Quote struct {
	Ask       string `json:"ask"`
	Bid       string `json:"bid"`
	Finalized string `json:"date"`
}

type Rate struct {
	ID        uint `gorm:"primary_key"`
	Currency  string
	Value     sql.NullFloat64
	Finalized time.Time
	CreatedAt time.Time
}

const OandaUrl = "https://www.oanda.com/rates/api/v1/rates/DZD.json"

func task(db gorm.DB) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", OandaUrl, nil)
	log.Print(fmt.Sprintf("Bearer %s", config.Conf.OANDA_KEY))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Conf.OANDA_KEY))
	response, err := client.Do(req)
	defer response.Body.Close()
	if err != nil {
		fmt.Errorf("Error while getting rates from oanda")
		fmt.Print(err)
		return
	}

	log.Print(response)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Errorf("Error while getting bytes from oanda response")
		fmt.Print(err)
		return
	}

	var r Response
	json.Unmarshal(body, &r)
	log.Print(r)

	for currency := range r.Quotes {
		date, err := time.Parse("2006-01-02T15:04:05", r.Quotes[currency].Finalized[:19])
		if err != nil {
			log.Print(err)
			continue
		}
		bid, err := strconv.ParseFloat(r.Quotes[currency].Bid, 64)
		rate := Rate{Currency:currency, Value:sql.NullFloat64{Float64:bid, Valid:true}, Finalized:date}
		db.Create(&rate)

	}
}

func main() {
	log.Print(config.Conf.MysqlArgs())
	db, err := gorm.Open("mysql", config.Conf.MysqlArgs())
	if err != nil {
		log.Print(err)
		log.Fatal("cannot connect to database...")
	}
	db.DB()
	db.CreateTable(&Rate{})

	task(db)

	s := gocron.NewScheduler()
	s.Every(5).Minutes().Do(task, db)
	<- s.Start()
}
