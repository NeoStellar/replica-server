package qr_koordinati

import (
	"github.com/gofiber/fiber/v2"
)

// /api/qr_koordinati
//
//	@Summary      	Qr Koordinatı Gösterir
//	@Description  	Takımlar sunucuya gönderecekleri sorgu ile müsabakada kullanılacak olan QR kodunun konumunu alabilmektedir
//	@Tags			Qr Koordinatı
//	@Accept       	json
//	@Produce      	json
//	@Success      	200
//	@Router       	/api/qr_koordinati [get]
func sendLockInfo(c *fiber.Ctx) error {
	return c.JSON(QrKoordinati{
		QrEnlem:  0.0,
		QrBoylam: 0.0,
	})
}
