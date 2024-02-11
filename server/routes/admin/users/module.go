package users

import (
	"encoding/json"
	"iharacee/server"
	"slices"

	"github.com/gofiber/fiber/v2"
)

func init() {
	app := server.ADMIN
	router := app.Group("/users")
	router.Use(func(ctx *fiber.Ctx) error {
		if !checkLoggedUserIsAdmin(ctx) {
			return ctx.Status(403).JSON(fiber.Map{
				"message": "Bu işlemi yapmaya yetkiniz yok",
			})
		}
		return ctx.Next()
	})
	router.Post("/", createUser)
	router.Get("/", getUsers)
}

func checkLoggedUserIsAdmin(ctx *fiber.Ctx) bool {
	var userDoc server.UserData
	sess, err := server.SessionStore.Get(ctx)
	if err != nil {
		return false
	}
	token := sess.Get("takim")
	if token == nil || token.(string) == "" {
		return false
	}
	json.Unmarshal([]byte(token.(string)), &userDoc)
	idx := slices.IndexFunc(userDoc.Roles, func(role string) bool {
		return role == "admin"
	})
	return idx != -1
}

type AuthObject struct {
	Username string `json:"kadi" bson:"kadi"`
	Password string `json:"sifre" bson:"sifre"`
}
