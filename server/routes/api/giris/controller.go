package giris

import (
	"encoding/json"
	"iharacee/server"
	"log"

	"github.com/gofiber/fiber/v2"
)

type girisResult200 struct {
	Takim_no string `json:"takim_no" bson:"takim_no"`
}

type girisResult500 struct {
	Error string `json:"error" bson:"error"`
}

// Sunucuya Giriş Yapmak
//
//	@Summary      Giriş Yapmak için kullanılır
//	@Description  Status Code'lar kesin olarak bu şekilde, 200 dönütünde aldığımız obje farklı olabilir (!)
//	@Tags         giris
//	@Accept       json
//	@Produce      json
//	@Param        AuthObject body AuthObject true "Giriş Bilgileri"
//	@Success      200  {object}  girisResult200
//	@Failure      400
//	@Failure      500  {object} girisResult500
//	@Router       /api/giris [post]
func signSessionUser(ctx *fiber.Ctx) error {
	var requestBody server.AuthObject
	sess, err := server.SessionStore.Get(ctx)
	if err != nil {
		log.Println(err)
		return ctx.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := ctx.BodyParser(&requestBody); err != nil {
		log.Println("parseError: " + err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	userDoc, err := server.GetUserData(ctx, requestBody)
	if err != nil {
		log.Println("getUserDataError: " + err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"error": "no user found with these credentials",
		})
	}
	//takim_no = countTeams()
	out, _ := json.Marshal(userDoc)
	sess.Set("takim", string(out))
	if err := sess.Save(); err != nil {
		log.Println(err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"takim_no": userDoc.Takim_no,
	})
}

/* func countTeams() int64 {
	redisString := server.Redis.Get("takim_sayisi")
	i, _ := strconv.Atoi(redisString)
	server.Redis.Set("takim_sayisi", strconv.Itoa(i+1))
	return int64(i + 1)
} */
