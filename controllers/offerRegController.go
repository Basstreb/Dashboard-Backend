package controllers

import (
	"dashboard/database"
	"dashboard/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateOfferReg(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	price, err := strconv.ParseFloat(string(data["price"]), 64)

	if err != nil {
		return err
	}

	priceIva := price * 1.21

	percent, err := strconv.ParseFloat(data["percentN"], 64)

	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(string(data["id"]), 10, 64)

	if err != nil {
		return err
	}

	finalPrice := priceIva * (percent / 100)

	staff := models.OffersRegData{
		OfferId:    id,
		Price:      priceIva,
		Percent:    percent,
		PriceFinal: finalPrice,
		CreatedAt:  time.Now().UTC(),
	}

	database.DB.Create(&staff)

	return c.JSON(staff)
}

func UpdateOfferReg(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	price, err := strconv.ParseFloat(string(data["price"]), 64)

	if err != nil {
		return err
	}

	priceIva := price * 1.21

	percent, err := strconv.ParseFloat(data["percent"], 64)

	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(string(data["id"]), 10, 64)

	if err != nil {
		return err
	}

	finalPrice := price * (percent / 100)

	offerReg := models.OffersRegData{
		Price:      priceIva,
		Percent:    percent,
		PriceFinal: finalPrice,
	}

	database.DB.Where("offer_id = ?", id).Model(&offerReg).Updates(offerReg)

	return c.JSON(offerReg)
}

func SendOfferReg(c *fiber.Ctx) error {
	var result []models.OffersRegData

	database.DB.Raw("SELECT * FROM offers_reg_data").Scan(&result)
	return c.JSON(result)
}
