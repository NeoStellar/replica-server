package telemetri_gonder

import (
	"encoding/json"
	"iharacee/server"
	"iharacee/server/routes/api/sunucusaati"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Telemetri Gönder
//
//	@Summary      Telemetri Göndermek için kullanılır
//	@Description  Hava aracının bilgilerini anlık olarak sunucuya göndermek ve diğer takımların bilgilerini almak için kullanılır
//	@Tags         API
//	@Accept       json
//	@Produce      json
//	@Param        TelemetryData body TelemetryData true "Telemetri Bilgileri"
//	@Success      200  {object}  TelemetryDataResponse
//	@Failure      400
//	@Router       /api/telemetri_gonder [post]
func SendTelemetryData(ctx *fiber.Ctx) error {
	var requestBody TelemetryData
	var takimData server.UserData
	if err := ctx.BodyParser(&requestBody); err != nil {
		return ctx.SendStatus(400)
	}
	sess, err := server.SessionStore.Get(ctx)
	if err != nil {
		log.Println(err)
	}
	takim := sess.Get("takim")
	if err := json.Unmarshal([]byte(takim.(string)), &takimData); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "parseError: " + err.Error(),
		})
	}
	takim_no := takimData.Takim_no
	if requestBody.Takim_numarasi != takim_no {
		return ctx.SendStatus(400)
	}
	response := TelemetryDataResponse{
		Sunucusaati:    sunucusaati.GetServerTime(),
		KonumBilgileri: PushTelemetryData(requestBody),
	}
	return ctx.JSON(response)
}
