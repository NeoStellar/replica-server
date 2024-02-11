package giris

import (
	"iharacee/server"
)

var (
	UserData server.UserData
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
