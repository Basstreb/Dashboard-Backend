package controllers

import (
	"dashboard/database"
	"dashboard/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func CreateClient(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	client := models.ClientData{
		CompanyName:   data["companyName"],
		ManagerName:   data["managerName"],
		ManagerCharge: data["managerCharge"],
		PhoneNumber:   data["phoneNumber"],
		Cif:           data["cif"],
		Direction:     data["direction"],
		CreatedAt:     time.Now().UTC(),
	}

	database.DB.Create(&client)

	return c.JSON(client)
}

func SendClientData(c *fiber.Ctx) error {
	var result []models.ClientData

	database.DB.Raw("SELECT * FROM client_data").Scan(&result)
	logrus.Info(result)
	return c.JSON(result)
}

func DeleteClientData(c *fiber.Ctx) error {
	var data map[string]uint64

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	erase := models.ClientData{
		Id: data["id"],
	}

	database.DB.Delete(&erase)

	return c.JSON(erase)
}

func DeleteClientOffers(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	eraseOffers := models.OffersData{}

	id := data["clientId"]

	database.DB.Where("client_id = ?", id).Delete(&eraseOffers)

	return c.JSON(eraseOffers)
}

func UpdateClientData(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	client := models.ClientData{
		CompanyName:   data["companyName"],
		ManagerName:   data["managerName"],
		ManagerCharge: data["managerCharge"],
		PhoneNumber:   data["phoneNumber"],
		Cif:           data["cif"],
		Direction:     data["direction"],
	}

	id, err := strconv.Atoi(data["id"])

	if err != nil {
		return err
	}

	database.DB.Where("id = ?", id).Model(&client).Updates(client)

	return c.JSON(client)
}

func UpdateClientOffers(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	client := models.OffersData{
		CompanyName: data["companyName"],
	}

	id := data["clientId"]

	database.DB.Where("client_id = ?", id).Model(&client).Updates(client)

	return c.JSON(client)
}
