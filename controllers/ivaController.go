package controllers

import (
	"dashboard/database"
	"dashboard/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func IvaPriceSupported(c *fiber.Ctx) error {
	var query []models.IvaAmountModel
	database.DB.Raw(`
	SELECT MONTH(od.created_at) AS 'month', SUM(od.price_iva - od.price) AS 'amount'
FROM offers_data od
WHERE od.deleted_at IS NULL
AND od.status = 'APPROVED'
OR od.status = 'PAYMENT_PENDING'
OR od.status = 'PAYD'
GROUP BY MONTH(od.created_at);`).Scan(&query)
	return c.JSON(query)
}

func IvaPriceRepercuted(c *fiber.Ctx) error {
	var query []models.IvaAmountModel
	database.DB.Raw(`
	SELECT MONTH(cc.created_at) AS 'month', SUM(cc.amount - cc.amount_w) AS 'amount'
FROM common_costs cc
WHERE cc.deleted_at IS NULL
GROUP BY MONTH(cc.created_at);`).Scan(&query)
	return c.JSON(query)
}

func IvaDataSoported(c *fiber.Ctx) error {
	var query []models.IvaModel
	database.DB.Raw(`
	SELECT od.id ,od.offer_name AS 'name', od.price_iva - od.price AS 'iva', od.created_at AS 'date'
FROM offers_data od
WHERE od.deleted_at IS NULL
AND od.status = 'APPROVED'
OR od.status = 'PAYMENT_PENDING'
OR od.status = 'PAYD';`).Scan(&query)
	return c.JSON(query)
}

func IvaDataRepercuted(c *fiber.Ctx) error {
	var query []models.IvaModel
	database.DB.Raw(`
	SELECT cc.id ,cc.common_costs_name  AS 'name', cc.amount - cc.amount_w AS 'iva', cc.created_at AS 'date'
FROM common_costs cc
WHERE cc.deleted_at IS NULL;`).Scan(&query)
	return c.JSON(query)
}

func IvaDataAcumulative(c *fiber.Ctx) error {
	var ivaSopor float64
	var ivaReper float64
	var ivaPaid float64

	database.DB.Raw(`
	SELECT ROUND(SUM(od.price_iva - od.price) ,2)
	FROM offers_data od
	WHERE od.deleted_at IS NULL
	AND status = 'APPROVED'
	OR status = 'PAYMENT_PENDING'
	OR status = 'PAYD';`).Scan(&ivaSopor)

	database.DB.Raw(`
	SELECT ROUND(SUM(cc.amount - cc.amount_w) ,2)
	FROM common_costs cc
	WHERE cc.deleted_at IS NULL;`).Scan(&ivaReper)

	database.DB.Raw(`
	SELECT ROUND(SUM(ip.amount), 2)
	FROM iva_paids ip
	WHERE ip.deleted_at IS NULL;`).Scan(&ivaPaid)

	totalIva := ivaSopor + ivaReper - ivaPaid

	return c.JSON(totalIva)
}

func CreateIvaPaid(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	iva, err := strconv.ParseFloat(data["iva"], 64)

	if err != nil {
		return err
	}

	lyt := "2006-01-02T15:04:05.000Z"
	date, err := time.Parse(lyt, data["date"]+"T00:00:00.000Z")

	if err != nil {
		return err
	}

	ivaP := models.IvaPaid{
		Amount: iva,
		Date:   date,
	}

	database.DB.Create(&ivaP)

	return c.JSON(ivaP)
}

func ListIvaPaid(c *fiber.Ctx) error {
	var query []models.IvaPaidModel
	database.DB.Raw(`
	SELECT ip.id, ip.amount, ip.date
FROM iva_paids ip
WHERE ip.deleted_at IS NULL;`).Scan(&query)
	return c.JSON(query)
}

func IvaPricePaid(c *fiber.Ctx) error {
	var query []models.IvaAmountModel
	database.DB.Raw(`
	SELECT MONTH(ip.date) AS 'month', SUM(ip.amount) AS 'amount'
FROM iva_paids ip
WHERE ip.deleted_at IS NULL
GROUP BY MONTH(ip.date);`).Scan(&query)
	return c.JSON(query)
}
