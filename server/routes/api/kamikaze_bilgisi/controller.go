package kamikaze_bilgisi

import (
	"encoding/json"
	"iharacee/server"

	"github.com/gofiber/fiber/v2"
)

// /api/kamikaze_bilgisi
//
//	@Summary      	Kilitlenme Bilgisi Gönderir
//	@Description  	Takımlar gerçekleştirdikleri başarılı kamikaze görevi ardından sunucuya kamikaze bilgisi göndermelidir.
//	@Tags			API
//	@Param			KamikazeVerisi 	body	KamikazeData	true	"Örnek Kamikaze Verisi"
//	@Accept       	json
//	@Produce      	json
//	@Success      	200
//	@Router       	/api/kamikaze_bilgisi [post]
func sendLockInfo(ctx *fiber.Ctx) error {
	var kamikazeData KamikazeData
	if err := ctx.BodyParser(&kamikazeData); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Kamikaze verisi alınamadı",
		})
	}
	session, err := server.SessionStore.Get(ctx)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": "Session alınamadı",
		})
	}
	takim := session.Get("takim")
	if takim == nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": "Takım alınamadı",
		})
	}
	var takimData server.UserData
	if err := json.Unmarshal([]byte(takim.(string)), &takimData); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "parseError: " + err.Error(),
		})
	}
	// Note the method has pointer receiver, so use a pointer to your value:
	kamikazeDataDocument := &KamikazeDataDocument{
		KamikazeStartTime: kamikazeData.KamikazeStartTime,
		KamikazeEndTime:   kamikazeData.KamikazeEndTime,
		QrText:            kamikazeData.QrText,
		TeamNumber:        takimData.Takim_no,
	}
	if _, err := kamikazeCollection.InsertOne(ctx.Context(), kamikazeDataDocument); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": "Kamikaze verisi kaydedilemedi",
		})
	}
	return ctx.SendStatus(200)
}
