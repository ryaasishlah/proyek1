package controller

import (
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/ryaasishlah/proyek1/config"
	"github.com/ryaasishlah/ryaasbackend"
	"github.com/whatsauth/whatsauth"
)

var DataLapangan = "Lapangan"
var DataKategori = "Kategori"
var DataDiskon = "Diskon"
var DataBank = "Bank"
var DataKontak = "Kontak"

// type HTTPRequest struct {
// 	Header string `json:"header"`
// 	Body   string `json:"body"`
// }

// func Sink(c *fiber.Ctx) error {
// 	var req HTTPRequest
// 	req.Header = string(c.Request().Header.Header())
// 	req.Body = string(c.Request().Body())
// 	return c.JSON(req)
// }

func WsWhatsAuthQR(c *websocket.Conn) {
	whatsauth.RunSocket(c, config.PublicKey, config.Usertables[:], config.Ulbimariaconn)
}

func PostWhatsAuthRequest(c *fiber.Ctx) error {
	if string(c.Request().Host()) == config.Internalhost {
		var req whatsauth.WhatsauthRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}
		ntfbtn := whatsauth.RunModuleLegacy(req, config.PrivateKey, config.Usertables[:], config.Ulbimariaconn)
		return c.JSON(ntfbtn)
	} else {
		var ws whatsauth.WhatsauthStatus
		ws.Status = string(c.Request().Host())
		return c.JSON(ws)
	}

}

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetLapangan(c *fiber.Ctx) error {
	getstatus := ryaasbackend.GetDataLapangan("Farhan Rizki Maulana")
	return c.JSON(getstatus)
}

func GetKategori(c *fiber.Ctx) error {
	getstatus := ryaasbackend.GetDataKategori("tournaments")
	return c.JSON(getstatus)
}

func GetDiskon(c *fiber.Ctx) error {
	getstatus := ryaasbackend.GetDataDiskon("35000")
	return c.JSON(getstatus)
}

func GetBank(c *fiber.Ctx) error {
	getstatus := ryaasbackend.GetDataBank("Microtfon")
	return c.JSON(getstatus)
}
func GetKontak(c *fiber.Ctx) error {
	getstatus := ryaasbackend.GetDataKontak("0821247838192")
	return c.JSON(getstatus)
}
