package qr_koordinati

import (
	"github.com/gofiber/fiber/v2"
)

// /api/qr_koordinati
//
//	@Summary      	Qr Koordinatı Gösterir
//	@Description  	Takımlar sunucuya gönderecekleri sorgu ile müsabakada kullanılacak olan QR kodunun konumunu alabilmektedir
//	@Tags			API
//	@Accept       	json
//	@Produce      	json
//	@Success      	200 {object}  QrCoordinates
//	@Router       	/api/qr_koordinati [get]
func sendLockInfo(c *fiber.Ctx) error {
	return c.JSON(QrCoordinates{
		QrEnlem:  39.970612,
		QrBoylam: 32.630087,
	})
}
