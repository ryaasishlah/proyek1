package url

import (
	"github.com/ryaasishlah/proyek1/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Homepage)
	page.Get("/Lapangan", controller.GetLapangan) //API from user whatsapp message from iteung gowa
	page.Get("/Kategori", controller.GetKategori) //API from user whatsapp message from iteung gowa
	page.Get("/Diskon", controller.GetDiskon)     //API from user whatsapp message from iteung gowa
	page.Get("/Bank", controller.GetBank)         //API from user whatsapp message from iteung gowa
	page.Get("/Kontak", controller.GetKontak)     //API from user whatsapp message from iteung gowa

	// page.Get("/", controller.Sink)
	// page.Post("/", controller.Sink)
	// page.Put("/", controller.Sink)
	// page.Patch("/", controller.Sink)
	// page.Delete("/", controller.Sink)
	// page.Options("/", controller.Sink)

}
