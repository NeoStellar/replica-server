package sunucusaati

import (
	"iharacee/server"
	"time"
)

func init() {
	app := server.API
	router := app.Group("/sunucusaati")
	router.Get("/", SendServerTime)
}

type ServerTime struct {
	Day        int `json:"gun" bson:"gun"`
	Hour       int `json:"saat" bson:"saat"`
	Minute     int `json:"dakika" bson:"dakika"`
	Second     int `json:"saniye" bson:"saniye"`
	Milisecond int `json:"milisaniye" bson:"milisaniye"`
}

func GetServerTime() ServerTime {
	return ServerTime{
		Day:        time.Now().Day(),
		Hour:       time.Now().Hour(),
		Minute:     time.Now().Minute(),
		Second:     time.Now().Second(),
		Milisecond: time.Now().Nanosecond() / 1000000,
	}
}
