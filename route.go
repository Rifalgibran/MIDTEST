package route

import (
	"strconv"
	"uts/Uts-Sds/database"
	"uts/Uts-Sds/models"

	"github.com/gofiber/fiber/v2"
)

func InsertData(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	err := database.CreateUser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"Pesan": "Data telah berhasil ditambahkan",
	})
}

func GetAllData(c *fiber.Ctx) error {
	users, err := database.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": users,
	})
}

func GetUserByID(c *fiber.Ctx) error {
	idParam := c.Params("id_user")
	idUser, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	user, err := database.GetUserByID(uint(idUser))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": user,
	})
}

func Delete(c *fiber.Ctx) error {
	// Mendapatkan id_user dari URL
	idParam := c.Params("id_user")
	idUser, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	err = database.DeleteUserByID(uint(idUser))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"Pesan": "Data telah dihapus",
	})
}

func Update(c *fiber.Ctx) error {
	// Mendapatkan id_user dari URL
	idParam := c.Params("id_user")
	_, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	update := models.User{
		Nama:     data["nama"],
		Email:    data["email"],
		Password: data["password"],
	}

	err = database.UpdateUser(&update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"Pesan": "Data User telah diupdate",
	})
}
