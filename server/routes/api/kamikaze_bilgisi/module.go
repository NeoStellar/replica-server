package kamikaze_bilgisi

import (
	"iharacee/server"
)

func init() {
	app := server.API
	router := app.Group("/kamikaze_bilgisi")
	router.Post("/", sendLockInfo)
}

type Time struct {
	Hour       int `json:"saat" bson:"saat"`
	Minute     int `json:"dakika" bson:"dakika"`
	Second     int `json:"saniye" bson:"saniye"`
	Milisecond int `json:"milisaniye" bson:"milisaniye"`
}

type KamikazeData struct {
	KamikazeStartTime Time   `json:"kamikazeBaslangicZamani" bson:"kamikazeBaslangicZamani"`
	KamikazeEndTime   Time   `json:"kamikaBitisZamani" bson:"kamikaBitisZamani"`
	QrText            string `json:"qrMetni" bson:"qrMetni"`
}
