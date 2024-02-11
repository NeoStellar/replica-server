package users

import (
	"iharacee/server"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// Kullanıcı Oluşturmak
//
//	@Summary      Kullanıcı Oluşturmak
//	@Description  Sadece adminlerin kullanabileceği bir endpointtir.
//	@Tags         ADMIN
//	@Accept       json
//	@Produce      json
//	@Param        AuthObject body AuthObject true "Giriş Bilgileri"
//	@Success      200
//	@Failure      400
//	@Failure      500
//	@Router       /admin/users [post]
func createUser(ctx *fiber.Ctx) error {
	var userDoc, testDoc server.UserData
	if err := ctx.BodyParser(&userDoc); err != nil {
		return ctx.SendStatus(400)
	}
	userDoc.Roles = []string{"user"}
	if err := server.UserCollection.FindOne(ctx.Context(), bson.D{
		{Key: "kadi", Value: userDoc.Username},
	}).Decode(&testDoc); err == nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bu kullanıcı adı zaten alınmış",
		})
	}
	if userCount, err := server.UserCollection.CountDocuments(ctx.Context(), bson.D{}); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	} else {
		userDoc.Takim_no = userCount
		if _, err := server.UserCollection.InsertOne(ctx.Context(), userDoc); err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return ctx.SendStatus(200)
	}
}

func getUsers(ctx *fiber.Ctx) error {
	var users []server.UserData
	cursor, err := server.UserCollection.Find(ctx.Context(), bson.D{})
	if err != nil {
		log.Println(err)
		return ctx.SendStatus(500)
	}
	if err = cursor.All(ctx.Context(), &users); err != nil {
		log.Println(err)
		return ctx.SendStatus(500)
	}
	return ctx.JSON(users)
}
