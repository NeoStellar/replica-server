package giris

import (
	"iharacee/server"
)

var (
	username = "kadi"
	password = "sifre"
	takim_no int64
)

func init() {
	app := server.API
	router := app.Group("/giris")
	router.Post("/", signSessionUser)
}

type AuthObject struct {
	Username string `json:"kadi" bson:"kadi"`
	Password string `json:"sifre" bson:"sifre"`
}
