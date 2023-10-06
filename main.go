package main

import (
	"log"
	"uts/Uts-Sds/database"
	"uts/Uts-Sds/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Menghubungkan ke database
	database.Connect()

	// Membuat aplikasi Fiber
	app := fiber.New()

	// Endpoint untuk menambahkan data
	app.Post("/insert", route.InsertData)

	// Endpoint untuk mendapatkan semua data
	app.Get("/getData", route.GetAllData)

	// Endpoint untuk mendapatkan data berdasarkan ID
	app.Get("/getDataUser/:id_user", route.GetUserByID)

	// Endpoint untuk menghapus data
	app.Get("/delete/:id_user", route.Delete)

	// Endpoint untuk mengupdate data
	app.Put("/update/:id_user", route.Update)

	// Menjalankan aplikasi pada port 3000
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
