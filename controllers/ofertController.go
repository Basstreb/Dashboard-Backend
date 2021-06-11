package controllers

import (
	"dashboard/database"
	"dashboard/models"
	"reflect"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var codes string

func CreateOffer(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	// price, err := strconv.ParseUint(string(data["price"]), 10, 64)
	//Parse string to float
	price, err := strconv.ParseFloat(string(data["price"]), 64)

	if err != nil {
		return err
	}

	//Parse string to uint
	clientId, err := strconv.ParseUint(string(data["clientId"]), 10, 64)

	if err != nil {
		return err
	}

	//Format Date
	lyt := "2006-01-02T15:04:05.000Z"
	date, err := time.Parse(lyt, data["decisionDate"]+"T00:00:00.000Z")

	if err != nil {
		return err
	}

	var result []models.OffersData
	database.DB.Raw("SELECT id FROM offers_data ORDER BY id DESC LIMIT 1").Scan(&result)

	if len(result) > 0 {
		id_check := (result[0].Id + 1)
		i := strconv.FormatUint(id_check, 10)

		if id_check > 99 {
			codes = "OF_21_" + i
		} else if id_check > 9 {
			codes = "OF_21_0" + i
		} else {
			codes = "OF_21_00" + i
		}

		offer := models.OffersData{
			ClientId:     clientId,
			CompanyName:  data["companyName"],
			Offer:        data["offer"],
			OfferName:    data["offerName"],
			Code:         codes,
			DecisionDate: date,
			Status:       "PENDING",
			Price:        price,
			Percent:      0,
			CreatedAt:    time.Now().UTC(),
		}

		database.DB.Create(&offer)

		return c.JSON(offer)
	} else {
		var id_check uint64 = 1
		i := strconv.FormatUint(id_check, 10)

		if id_check > 99 {
			codes = "OF_21_" + i
		} else if id_check > 9 {
			codes = "OF_21_0" + i
		} else {
			codes = "OF_21_00" + i
		}

		offer := models.OffersData{
			ClientId:     clientId,
			CompanyName:  data["companyName"],
			Offer:        data["offer"],
			OfferName:    data["offerName"],
			Code:         codes,
			DecisionDate: date,
			Status:       "PENDING",
			Price:        price,
			Percent:      0,
			CreatedAt:    time.Now().UTC(),
		}

		database.DB.Create(&offer)

		return c.JSON(offer)
	}
}

func SendOfferData(c *fiber.Ctx) error {
	var result []models.OffersData

	database.DB.Raw("SELECT * FROM offers_data").Scan(&result)
	logrus.Info(result)
	return c.JSON(result)
}

func DeleteOfferData(c *fiber.Ctx) error {
	var data map[string]uint64

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	erase := models.OffersData{
		Id: data["id"],
	}

	database.DB.Delete(&erase)

	return c.JSON(erase)
}

func UpdateOfferData(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	// price, err := strconv.ParseUint(string(data["price"]), 10, 64)
	price, err := strconv.ParseFloat(string(data["price"]), 10)

	if err != nil {
		return err
	}

	// Parse "0001-01-01" to time.Time type
	lyt := "2006-01-02T15:04:05.000Z"
	date, err := time.Parse(lyt, data["decisionDate"]+"T00:00:00.000Z")

	if err != nil {
		return err
	}

	// Parse string to uint8
	percent, err := strconv.ParseUint(data["percent"], 10, 8)

	if err != nil {
		return err
	}

	logrus.Info(reflect.TypeOf(percent))
	logrus.Info(percent)

	offer := models.OffersData{
		Offer:        data["offer"],
		OfferName:    data["offerName"],
		DecisionDate: date,
		Status:       data["status"],
		Price:        price,
		Percent:      percent,
	}

	id, err := strconv.Atoi(data["id"])

	if err != nil {
		return err
	}

	database.DB.Where("id = ?", id).Model(&offer).Updates(offer)

	return c.JSON(offer)
}

func FilterPriceDataMonth(c *fiber.Ctx) error {

	var query []models.MonthPriceOffer
	database.DB.Raw(`
	SELECT MONTH(decision_date) AS 'month', SUM(price) AS 'price'
	FROM offers_data
	WHERE deleted_at IS NULL
	AND status = 'APPROVED'
	OR status = 'PAYMENT_PENDING'
	OR status = 'PAYD'
	GROUP BY MONTH(decision_date);`).Scan(&query)
	logrus.Info(query)
	return c.JSON(query)
}
