package controllers

import (
	"dashboard/database"
	"dashboard/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SendCommonCostData(c *fiber.Ctx) error {
	var result []models.CostsCom

	database.DB.Raw(`SELECT * FROM common_costs cc JOIN typo_costs tc ON cc.id = tc.id;`).Scan(&result)
	logrus.Info(result)
	return c.JSON(result)
}

func CreateCommonCost(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	clientId, err := strconv.ParseUint(string(data["clientId"]), 10, 64)

	if err != nil {
		return err
	}

	lyt := "2006-01-02T15:04:05.000Z"
	date, err := time.Parse(lyt, data["costDate"]+"T00:00:00.000Z")

	if err != nil {
		return err
	}

	iva, err := strconv.ParseUint(string(data["iva"]), 10, 64)

	if err != nil {
		return err
	}

	amount, err := strconv.ParseFloat(string(data["amount"]), 64)

	if err != nil {
		return err
	}

	commonCost := models.CommonCosts{
		CommonCostsName: data["commonCostsName"],
		ClientId:        clientId,
		ClientName:      data["clientName"],
		Amount:          amount,
		CostDate:        date,
		CreatedAt:       time.Now().UTC(),
	}

	typoCost := models.TypoCosts{
		Typo:      data["typo"],
		Iva:       iva,
		CreatedAt: time.Now().UTC(),
	}

	database.DB.Create(&commonCost)
	database.DB.Create(&typoCost)

	return c.JSON(commonCost)
}

func UpdateCommonCost(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	clientId, err := strconv.ParseUint(string(data["clientId"]), 10, 64)

	if err != nil {
		return err
	}

	lyt := "2006-01-02T15:04:05.000Z"
	date, err := time.Parse(lyt, data["costDate"]+"T00:00:00.000Z")

	if err != nil {
		return err
	}

	iva, err := strconv.ParseUint(string(data["iva"]), 10, 64)

	if err != nil {
		return err
	}

	amount, err := strconv.ParseFloat(string(data["amount"]), 64)

	if err != nil {
		return err
	}

	commonCost := models.CommonCosts{
		CommonCostsName: data["commonCostsName"],
		ClientId:        clientId,
		ClientName:      data["clientName"],
		Amount:          amount,
		CostDate:        date,
		CreatedAt:       time.Now().UTC(),
	}

	typoCost := models.TypoCosts{
		Typo:      data["typo"],
		Iva:       iva,
		CreatedAt: time.Now().UTC(),
	}

	id, err := strconv.Atoi(data["id"])

	if err != nil {
		return err
	}

	database.DB.Where("id = ?", id).Model(&commonCost).Updates(commonCost)
	database.DB.Where("id = ?", id).Model(&typoCost).Updates(typoCost)

	return c.JSON(commonCost)
}

func DeleteCommonData(c *fiber.Ctx) error {
	var data map[string]uint64

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	eraseComm := models.CommonCosts{
		Id: data["id"],
	}

	eraseTypo := models.TypoCosts{
		Id: data["id"],
	}

	database.DB.Delete(&eraseComm)
	database.DB.Delete(&eraseTypo)

	return c.JSON(eraseComm)
}

func FilterCostDataMonth(c *fiber.Ctx) error {
	var query []models.MonthPriceCost
	database.DB.Raw(`
	SELECT MONTH(cost_date) AS 'month', SUM(amount) AS 'amount'
	FROM common_costs cc
	JOIN typo_costs tc
	ON cc.id = tc.id
	WHERE cc.deleted_at IS NULL
	GROUP BY MONTH(cost_date);`).Scan(&query)
	logrus.Info(query)
	return c.JSON(query)
}
