package sunucusaati

import (
	"github.com/gofiber/fiber/v2"
)

// /api/sunucusaati
//
//	@Summary      Sunucu Saatini Gönderir
//	@Description  Sunucu saatini sorgulamak için kullanılır
//	@Tags         sunucu saati
//	@Accept       json
//	@Produce      json
//	@Success      200  {object}  ServerTime
//	@Router       /api/sunucusaati [get]
func SendServerTime(c *fiber.Ctx) error {
	return c.JSON(GetServerTime())
}
