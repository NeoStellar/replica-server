package telemetri_gonder

import (
	"iharacee/server"
	"iharacee/server/routes/api/sunucusaati"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Telemetri Gönder
//
//	@Summary      Telemetri Göndermek için kullanılır
//	@Description  Hava aracının bilgilerini anlık olarak sunucuya göndermek ve diğer takımların bilgilerini almak için kullanılır
//	@Tags         Telemetri Gönder
//	@Accept       json
//	@Produce      json
//	@Param        TelemetryData body TelemetryData true "Telemetri Bilgileri"
//	@Success      200  {object}  TelemetryDataResponse
//	@Failure      400
//	@Router       /api/telemetri_gonder [post]
func SendTelemetryData(ctx *fiber.Ctx) error {
	var requestBody TelemetryData
	if err := ctx.BodyParser(&requestBody); err != nil {
		return ctx.SendStatus(400)
	}
	sess, err := server.SessionStore.Get(ctx)
	if err != nil {
		log.Println(err)
	}
	takim_no := sess.Get("takim_no").(int64)
	if requestBody.Takim_numarasi != takim_no {
		return ctx.SendStatus(400)
	}
	response := TelemetryDataResponse{
		Sunucusaati:    sunucusaati.GetServerTime(),
		KonumBilgileri: PushTelemetryData(requestBody),
	}
	return ctx.JSON(response)
}
