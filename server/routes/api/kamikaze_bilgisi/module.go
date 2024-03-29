package kamikaze_bilgisi

import (
	"iharacee/server"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	kamikazeCollection = server.Mongo.Collection("api.kamikaze_bilgisi")
)

func init() {
	app := server.API
	router := app.Group("/kamikaze_bilgisi")
	router.Post("/", sendLockInfo)
}

type Servertime struct {
	Hour       int `json:"saat" bson:"saat"`
	Minute     int `json:"dakika" bson:"dakika"`
	Second     int `json:"saniye" bson:"saniye"`
	Milisecond int `json:"milisaniye" bson:"milisaniye"`
}

type KamikazeData struct {
	KamikazeStartTime Servertime `json:"kamikazeBaslangicZamani" bson:"kamikazeBaslangicZamani"`
	KamikazeEndTime   Servertime `json:"kamikazeBitisZamani" bson:"kamikazeBitisZamani"`
	QrText            string     `json:"qrMetni" bson:"qrMetni"`
}

type KamikazeDataDocument struct {
	ID         string       `json:"id,omitempty" bson:"_id,omitempty"`
	Data       KamikazeData `json:"data" bson:"data"`
	TeamNumber int64        `json:"takim_no" bson:"takim_no"`
	CreatedAt  time.Time    `json:"created_at" bson:"created_at"`
}

func (u *KamikazeDataDocument) MarshalBSON() ([]byte, error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	type my KamikazeDataDocument
	return bson.Marshal((*my)(u))
}
