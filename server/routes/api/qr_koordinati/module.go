package qr_koordinati

import (
	"iharacee/server"
)

func init() {
	app := server.API
	router := app.Group("/qr_koordinati")
	router.Get("/", sendLockInfo)
}

type QrCoordinates struct {
	QrEnlem  float64 `json:"qrEnlem" bson:"qrEnlem"`
	QrBoylam float64 `json:"qrBoylam" bson:"qrBoylam"`
}
