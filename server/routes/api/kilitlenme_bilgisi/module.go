package kilitlenme_bilgisi

import (
	"iharacee/server"
)

func init() {
	app := server.API
	router := app.Group("/kilitlenme_bilgisi")
	router.Post("/", sendLockInfo)
}

type LockInfo struct {
	LockStartTime struct {
		Hour       int `json:"saat" bson:"saat"`
		Minute     int `json:"dakika" bson:"dakika"`
		Second     int `json:"saniye" bson:"saniye"`
		Milisecond int `json:"milisaniye" bson:"milisaniye"`
	} `json:"kilitlenmeBaslangicZamani" bson:"kilitlenmeBaslangicZamani"`
	LockEndTime struct {
		Hour       int `json:"saat" bson:"saat"`
		Minute     int `json:"dakika" bson:"dakika"`
		Second     int `json:"saniye" bson:"saniye"`
		Milisecond int `json:"milisaniye" bson:"milisaniye"`
	} `json:"kilitlenmeBitisZamani" bson:"kilitlenmeBitisZamani"`
	AutonomousLock int `json:"otonom_kilitlenme" bson:"otonom_kilitlenme"`
}
