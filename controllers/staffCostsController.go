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

	database.DB.Raw(`SELECT * FROM staff_costs WHERE deleted_at IS NULL`).Scan(&result)
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

	var project1Name string

	database.DB.Raw(`
SELECT
	od.offer_name
FROM
	offers_data od
WHERE
	od.deleted_at IS NULL
	AND id = ?`, data["project1"]).Scan(&project1Name)

	var project2Name string

	database.DB.Raw(`
SELECT
	od.offer_name
FROM
	offers_data od
WHERE
	od.deleted_at IS NULL
	AND id = ?`, data["project2"]).Scan(&project2Name)

	var project3Name string

	database.DB.Raw(`
SELECT
	od.offer_name
FROM
	offers_data od
WHERE
	od.deleted_at IS NULL
	AND id = ?`, data["project3"]).Scan(&project3Name)

	var project4Name string

	database.DB.Raw(`
SELECT
	od.offer_name
FROM
	offers_data od
WHERE
	od.deleted_at IS NULL
	AND id = ?`, data["project4"]).Scan(&project4Name)

	staff := models.StaffCosts{
		StaffName:        data["staffName"],
		Amount:           amount,
		Cost:             cost,
		SocialInsurances: social,
		Project1:         data["project1"],
		Project1Name:     project1Name,
		Percentage1:      per1,
		Project2:         data["project2"],
		Project2Name:     project2Name,
		Percentage2:      per2,
		Project3:         data["project3"],
		Project3Name:     project3Name,
		Percentage3:      per3,
		Project4:         data["project4"],
		Project4Name:     project4Name,
		Percentage4:      per4,
		PayDate:          date,
		CreatedAt:        time.Now().UTC(),
	}

	database.DB.Create(&staff)

	return c.JSON(staff)
}

func CreateStaffProjectData(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		logrus.Info(err)
		return err
	}

	var id uint64

	database.DB.Raw(`
	SELECT id
	FROM staff_costs sc
	ORDER BY id DESC
	LIMIT 1`).Scan(&id)

	amount, err := strconv.ParseFloat(string(data["amount"]), 64)

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

	total1 := amount * per1 / 100
	total2 := amount * per2 / 100
	total3 := amount * per3 / 100
	total4 := amount * per4 / 100

	if id == 0 {
		id = 1
	} else {
		id = id + 1
	}

	staff1 := models.StaffProject{
		IdStaff:    id,
		Project:    data["project1"],
		Amount:     amount,
		Percentage: per1,
		Total:      total1,
		CreatedAt:  time.Now().UTC(),
	}

	staff2 := models.StaffProject{
		IdStaff:    id,
		Project:    data["project2"],
		Amount:     amount,
		Percentage: per2,
		Total:      total2,
		CreatedAt:  time.Now().UTC(),
	}

	staff3 := models.StaffProject{
		IdStaff:    id,
		Project:    data["project3"],
		Amount:     amount,
		Percentage: per3,
		Total:      total3,
		CreatedAt:  time.Now().UTC(),
	}

	staff4 := models.StaffProject{
		IdStaff:    id,
		Project:    data["project4"],
		Amount:     amount,
		Percentage: per4,
		Total:      total4,
		CreatedAt:  time.Now().UTC(),
	}

	database.DB.Create(&staff1)
	database.DB.Create(&staff2)
	database.DB.Create(&staff3)
	database.DB.Create(&staff4)

	return c.JSON(staff1)
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

	var project1Name string

	database.DB.Raw(`
SELECT
	od.offer_name
FROM
	offers_data od
WHERE
	od.deleted_at IS NULL
	AND id = ?`, data["project1"]).Scan(&project1Name)

	var project2Name string

	database.DB.Raw(`
SELECT
	od.offer_name
FROM
	offers_data od
WHERE
	od.deleted_at IS NULL
	AND id = ?`, data["project2"]).Scan(&project2Name)

	var project3Name string

	database.DB.Raw(`
SELECT
	od.offer_name
FROM
	offers_data od
WHERE
	od.deleted_at IS NULL
	AND id = ?`, data["project3"]).Scan(&project3Name)

	var project4Name string

	database.DB.Raw(`
SELECT
	od.offer_name
FROM
	offers_data od
WHERE
	od.deleted_at IS NULL
	AND id = ?`, data["project4"]).Scan(&project4Name)

	staff := models.StaffCosts{
		StaffName:        data["staffName"],
		Amount:           amount,
		Cost:             cost,
		SocialInsurances: social,
		Project1:         data["project1"],
		Project1Name:     project1Name,
		Percentage1:      per1,
		Project2:         data["project2"],
		Project2Name:     project2Name,
		Percentage2:      per2,
		Project3:         data["project3"],
		Project3Name:     project3Name,
		Percentage3:      per3,
		Project4:         data["project4"],
		Project4Name:     project4Name,
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

func UpdateStaffProjectData(c *fiber.Ctx) error {
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

	total1 := amount * per1 / 100
	total2 := amount * per2 / 100
	total3 := amount * per3 / 100
	total4 := amount * per4 / 100

	idR, err := strconv.ParseUint(data["id"], 10, 64)

	logrus.Info(idR)

	if err != nil {
		return err
	}

	staff1 := models.StaffProject{
		IdStaff:    idR,
		Project:    data["project1"],
		Amount:     amount,
		Percentage: per1,
		Total:      total1,
		CreatedAt:  time.Now().UTC(),
	}

	staff2 := models.StaffProject{
		IdStaff:    idR,
		Project:    data["project2"],
		Amount:     amount,
		Percentage: per2,
		Total:      total2,
		CreatedAt:  time.Now().UTC(),
	}

	staff3 := models.StaffProject{
		IdStaff:    idR,
		Project:    data["project3"],
		Amount:     amount,
		Percentage: per3,
		Total:      total3,
		CreatedAt:  time.Now().UTC(),
	}

	staff4 := models.StaffProject{
		IdStaff:    idR,
		Project:    data["project4"],
		Amount:     amount,
		Percentage: per4,
		Total:      total4,
		CreatedAt:  time.Now().UTC(),
	}

	database.DB.Exec("DELETE FROM staff_projects WHERE id_staff = ?", idR)

	database.DB.Create(&staff1)
	database.DB.Create(&staff2)
	database.DB.Create(&staff3)
	database.DB.Create(&staff4)

	return c.JSON(staff1)
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

func DeleteStaffProjectData(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		logrus.Info(err)
		return err
	}

	idR, err := strconv.ParseUint(data["prId"], 10, 64)

	if err != nil {
		return err
	}

	database.DB.Exec("DELETE FROM staff_projects WHERE id_staff = ?", idR+1)

	return c.JSON(idR)
}
