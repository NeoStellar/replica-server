package kamikaze_bilgisi

import (
	"github.com/gofiber/fiber/v2"
)

// /api/kamikaze_bilgisi
//
//	@Summary      	Kilitlenme Bilgisi Gönderir
//	@Description  	Takımlar gerçekleştirdikleri her başarılı kilitlenmenin ardından sunucuya kilitlenme bilgisi göndermelidir
//	@Tags			kilitlenme bilgisi
//	@Param			KilitlenmeVerisi 	body	LockInfo	true	"Örnek Kilitlenme Verisi"
//	@Accept       	json
//	@Produce      	json
//	@Success      	200
//	@Router       	/api/kamikaze_bilgisi [post]
func sendLockInfo(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
