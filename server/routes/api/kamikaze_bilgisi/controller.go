package kamikaze_bilgisi

import (
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
func sendLockInfo(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
