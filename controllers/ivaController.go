package controllers

import (
	"dashboard/database"
	"dashboard/models"

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
	SELECT od.offer_name AS 'name', od.price_iva - od.price AS 'iva', od.created_at AS 'date' 
FROM offers_data od
WHERE od.deleted_at IS NULL;`).Scan(&query)
	return c.JSON(query)
}

func IvaDataRepercuted(c *fiber.Ctx) error {
	var query []models.IvaModel
	database.DB.Raw(`
	SELECT cc.common_costs_name  AS 'name', cc.amount - cc.amount_w AS 'iva', cc.created_at AS 'date'
FROM common_costs cc
WHERE cc.deleted_at IS NULL;`).Scan(&query)
	return c.JSON(query)
}

func IvaDataAcumulative(c *fiber.Ctx) error {
	var ivaSopor float64
	var ivaReper float64

	database.DB.Raw(`
	SELECT ROUND(SUM(od.price_iva - od.price) ,2)
	FROM offers_data od
	WHERE od.deleted_at IS NULL;`).Scan(&ivaSopor)

	database.DB.Raw(`
	SELECT ROUND(SUM(cc.amount - cc.amount_w) ,2)
	FROM common_costs cc
	WHERE cc.deleted_at IS NULL;`).Scan(&ivaReper)

	totalIva := ivaSopor + ivaReper

	return c.JSON(totalIva)
}
