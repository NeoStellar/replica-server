package telemetri_gonder

import (
	"iharacee/server"
	"iharacee/server/routes/api/sunucusaati"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	telemetryCollection = server.Mongo.Collection("telemetri_gonder")
	redis               server.RedisInstance
)

func init() {
	app := server.API
	redis = server.Redis
	router := app.Group("/telemetri_gonder")
	router.Post("/", SendTelemetryData)
}

type TelemetryData struct {
	Takim_numarasi  int64   `json:"takim_numarasi" bson:"takim_numarasi"`
	Iha_enlem       float64 `json:"iha_enlem" bson:"iha_enlem"`
	Iha_boylam      float64 `json:"iha_boylam" bson:"iha_boylam"`
	Iha_irtifa      int     `json:"iha_irtifa" bson:"iha_irtifa"`
	Iha_dikilme     int     `json:"iha_dikilme" bson:"iha_dikilme"`
	Iha_yonelme     int     `json:"iha_yonelme" bson:"iha_yonelme"`
	Iha_yatis       int     `json:"iha_yatis" bson:"iha_yatis"`
	Iha_hiz         int     `json:"iha_hiz" bson:"iha_hiz"`
	Iha_batarya     int     `json:"iha_batarya" bson:"iha_batarya"`
	Iha_otonom      int     `json:"iha_otonom" bson:"iha_otonom"`
	Iha_kilitlenme  int     `json:"iha_kilitlenme" bson:"iha_kilitlenme"`
	Hedef_merkez_X  int     `json:"hedef_merkez_X" bson:"hedef_merkez_X"`
	Hedef_merkez_Y  int     `json:"hedef_merkez_Y" bson:"hedef_merkez_Y"`
	Hedef_genislik  int     `json:"hedef_genislik" bson:"hedef_genislik"`
	Hedef_yukseklik int     `json:"hedef_yukseklik" bson:"hedef_yukseklik"`
	Gps_saati       struct {
		Saat       int `json:"saat" bson:"saat"`
		Dakika     int `json:"dakika" bson:"dakika"`
		Saniye     int `json:"saniye" bson:"saniye"`
		Milisaniye int `json:"milisaniye" bson:"milisaniye"`
	} `json:"gps_saati" bson:"gps_saati"`
}

type TelemetryDataResponse struct {
	Sunucusaati    sunucusaati.ServerTime `json:"sunucusaati" bson:"sunucusaati"`
	KonumBilgileri []TelemetryData        `json:"konumBilgileri" bson:"konumBilgileri"`
}

type TelemetryDataDocument struct {
	ID         string        `json:"id,omitempty" bson:"_id,omitempty"`
	Data       TelemetryData `json:"data" bson:"data"`
	TeamNumber int64         `json:"takim_no" bson:"takim_no"`
	CreatedAt  time.Time     `json:"created_at" bson:"created_at"`
}

func (u *TelemetryDataDocument) MarshalBSON() ([]byte, error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	type my TelemetryDataDocument
	return bson.Marshal((*my)(u))
}
