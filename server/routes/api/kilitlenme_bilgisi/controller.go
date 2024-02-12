package kilitlenme_bilgisi

import (
	"encoding/json"
	"iharacee/server"

	"github.com/gofiber/fiber/v2"
)

// /api/kilitlenme_bilgisi
//
//	@Summary      	Kilitlenme Bilgisi Gönderir
//	@Description  	Takımlar gerçekleştirdikleri her başarılı kilitlenmenin ardından sunucuya kilitlenme bilgisi göndermelidir
//	@Tags			API
//	@Param			KilitlenmeVerisi 	body	LockInfo	true	"Örnek Kilitlenme Verisi"
//	@Accept       	json
//	@Produce      	json
//	@Success      	200
//	@Router       	/api/kilitlenme_bilgisi [post]
func sendLockInfo(ctx *fiber.Ctx) error {
	var LockData LockInfo
	if err := ctx.BodyParser(&LockData); err != nil {
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
	kamikazeDataDocument := &LockInfoDocument{
		Data:       LockData,
		TeamNumber: takimData.Takim_no,
	}
	if _, err := kilitlenmeCollection.InsertOne(ctx.Context(), kamikazeDataDocument); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": "Kamikaze verisi kaydedilemedi",
		})
	}
	return ctx.SendStatus(200)
}
