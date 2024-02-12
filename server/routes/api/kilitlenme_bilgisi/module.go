package kilitlenme_bilgisi

import (
	"iharacee/server"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var kilitlenmeCollection = server.Mongo.Collection("kilitlenme_bilgisi")

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

type LockInfoDocument struct {
	ID            string `json:"id,omitempty" bson:"_id,omitempty"`
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
	AutonomousLock int       `json:"otonom_kilitlenme" bson:"otonom_kilitlenme"`
	TeamNumber     int64     `json:"takim_no" bson:"takim_no"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
}

func (u *LockInfoDocument) MarshalBSON() ([]byte, error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	type my LockInfoDocument
	return bson.Marshal((*my)(u))
}
