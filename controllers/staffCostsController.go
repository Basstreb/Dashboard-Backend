package controllers

import (
	"dashboard/database"
	"dashboard/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SendStaffCostData(c *fiber.Ctx) error {
	var result []models.StaffCosts

	database.DB.Raw(`SELECT * FROM staff_costs`).Scan(&result)
	return c.JSON(result)
}

func FilterStaffCostData(c *fiber.Ctx) error {
	var query []models.IvaAmountModel

	database.DB.Raw(`
	SELECT MONTH(sc.pay_date) AS 'month', SUM(sc.cost) AS 'amount'
	FROM staff_costs sc
	WHERE sc.deleted_at IS NULL
	GROUP BY MONTH(sc.pay_date);`).Scan(&query)
	return c.JSON(query)
}

func CreateStaffCostData(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		logrus.Info(err)
		return err
	}

	amount, err := strconv.ParseFloat(string(data["amount"]), 64)

	if err != nil {
		logrus.Info(err)

		return err
	}

	cost, err := strconv.ParseFloat(string(data["cost"]), 64)

	if err != nil {
		logrus.Info(err)
		return err
	}

	social := amount - cost

	// social, err := strconv.ParseFloat(string(data["socialInsurances"]), 64)

	// if err != nil {
	// 	logrus.Info(err)
	// 	return err
	// }

	// Parse "0001-01-01" to time.Time type
	lyt := "2006-01-02T15:04:05.000Z"
	date, err := time.Parse(lyt, data["payDate"]+"T00:00:00.000Z")

	if err != nil {
		logrus.Info(err)
		return err
	}

	per1, err := strconv.ParseFloat(string(data["per1"]), 64)

	if err != nil {
		logrus.Info(err)
		return err
	}

	per2, err := strconv.ParseFloat(string(data["per2"]), 64)

	if err != nil {
		logrus.Info(err)
		return err
	}

	per3, err := strconv.ParseFloat(string(data["per3"]), 64)

	if err != nil {
		logrus.Info(err)
		return err
	}

	per4, err := strconv.ParseFloat(string(data["per4"]), 64)

	if err != nil {
		logrus.Info(err)
		return err
	}

	staff := models.StaffCosts{
		StaffName:        data["staffName"],
		Amount:           amount,
		Cost:             cost,
		SocialInsurances: social,
		Project1:         data["project1"],
		Percentage1:      per1,
		Project2:         data["project2"],
		Percentage2:      per2,
		Project3:         data["project3"],
		Percentage3:      per3,
		Project4:         data["project4"],
		Percentage4:      per4,
		PayDate:          date,
		CreatedAt:        time.Now().UTC(),
	}

	database.DB.Create(&staff)

	return c.JSON(staff)
}

func UpdateStaffCostData(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	amount, err := strconv.ParseFloat(string(data["amount"]), 64)

	if err != nil {
		return err
	}

	cost, err := strconv.ParseFloat(string(data["cost"]), 64)

	if err != nil {
		return err
	}

	social := amount - cost

	// social, err := strconv.ParseFloat(string(data["socialInsurances"]), 64)

	// if err != nil {
	// 	return err
	// }

	// Parse "0001-01-01" to time.Time type
	lyt := "2006-01-02T15:04:05.000Z"
	date, err := time.Parse(lyt, data["payDate"]+"T00:00:00.000Z")

	if err != nil {
		return err
	}

	per1, err := strconv.ParseFloat(string(data["per1"]), 64)

	if err != nil {
		logrus.Info(err)
		return err
	}

	per2, err := strconv.ParseFloat(string(data["per2"]), 64)

	if err != nil {
		logrus.Info(err)
		return err
	}

	per3, err := strconv.ParseFloat(string(data["per3"]), 64)

	if err != nil {
		logrus.Info(err)
		return err
	}

	per4, err := strconv.ParseFloat(string(data["per4"]), 64)

	if err != nil {
		logrus.Info(err)
		return err
	}

	staff := models.StaffCosts{
		StaffName:        data["staffName"],
		Amount:           amount,
		Cost:             cost,
		SocialInsurances: social,
		Project1:         data["project1"],
		Percentage1:      per1,
		Project2:         data["project2"],
		Percentage2:      per2,
		Project3:         data["project3"],
		Percentage3:      per3,
		Project4:         data["project4"],
		Percentage4:      per4,
		PayDate:          date,
	}

	id, err := strconv.Atoi(data["id"])

	if err != nil {
		return err
	}

	database.DB.Where("id = ?", id).Model(&staff).Updates(staff)

	return c.JSON(staff)
}

func DeleteStaffData(c *fiber.Ctx) error {
	var data map[string]uint64

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	erase := models.StaffCosts{
		Id: data["id"],
	}

	database.DB.Delete(&erase)

	return c.JSON(erase)
}
