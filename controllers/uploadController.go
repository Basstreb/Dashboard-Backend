package controllers

import (
	"dashboard/database"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func UploadPdf(c *fiber.Ctx) error {

	var code string
	database.DB.Raw(`SELECT code FROM offers_data od ORDER BY od.code DESC LIMIT 1;`).Scan(&code)

	if code == "" {
		codes = "OF_21_001.pdf"
		// Parse the multipart form:
		if form, err := c.MultipartForm(); err == nil {
			// => *multipart.Form
			// Get all files from "documents" key:
			files := form.File["docs"]
			// => []*multipart.FileHeader

			// Loop through files:
			for _, file := range files {
				fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
				// => "tutorial.pdf" 360641 "application/pdf"

				if file.Header["Content-Type"][0] != "application/pdf" {
					return err
				}

				// Save the files to disk:
				if err := c.SaveFile(file, fmt.Sprintf("../frontend/public/pdf/%s", codes)); err != nil {
					return err
				}
			}
			return err
		} else {
			return err
		}
	} else {
		res1 := strings.Split(code, "_")
		num, err := strconv.Atoi(res1[2])

		if err != nil {
			return err
		}

		nums := strconv.Itoa(num)

		if num > 99 {
			codes = "OF_21_" + nums + ".pdf"
		} else if num > 9 {
			codes = "OF_21_0" + nums + ".pdf"
		} else {
			codes = "OF_21_00" + nums + ".pdf"
		}

		// Parse the multipart form:
		if form, err := c.MultipartForm(); err == nil {
			// => *multipart.Form
			// Get all files from "documents" key:
			files := form.File["docs"]
			// => []*multipart.FileHeader

			// Loop through files:
			for _, file := range files {
				fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
				// => "tutorial.pdf" 360641 "application/pdf"

				if file.Header["Content-Type"][0] != "application/pdf" {
					return err
				}

				// Save the files to disk:
				if err := c.SaveFile(file, fmt.Sprintf("../frontend/public/pdf/%s", codes)); err != nil {
					return err
				}
			}
			return err
		} else {
			return err
		}
	}
}

func UpdateUploadPdf(c *fiber.Ctx) error {

	var text string

	if form, err := c.MultipartForm(); err == nil {
		// => *multipart.Form
		// Get all files from "documents" key:
		files := form.File["docs"]
		// => []*multipart.FileHeader
		id, err := strconv.Atoi(form.Value["docs"][0])

		if err != nil {
			return err
		}

		database.DB.Raw(`SELECT code FROM offers_data od WHERE id = ?`, id).Scan(&text)
		text = text + ".pdf"
		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "application/pdf"

			if file.Header["Content-Type"][0] != "application/pdf" {
				return err
			}

			// Save the files to disk:
			if err := c.SaveFile(file, fmt.Sprintf("../frontend/public/pdf/%s", text)); err != nil {
				return err
			}
		}
		return err
	} else {
		return err
	}
}

func UploadImg(c *fiber.Ctx) error {

	var id int

	database.DB.Raw(`SELECT id
	FROM common_costs cc
	ORDER BY id DESC
	LIMIT 1;`).Scan(&id)

	imgName := "IMG_21_" + strconv.Itoa(id+1) + ".png"

	// Parse the multipart form:
	if form, err := c.MultipartForm(); err == nil {
		// => *multipart.Form
		// Get all files from "documents" key:
		files := form.File["imgs"]
		// => []*multipart.FileHeader

		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "image/type"

			imageCheck := strings.Split(file.Header["Content-Type"][0], "/")
			if imageCheck[0] != "image" {
				return err
			}

			// Save the files to disk:
			if err := c.SaveFile(file, fmt.Sprintf("../frontend/public/img/%s", imgName)); err != nil {
				return err
			}
		}
		return err
	} else {
		return err
	}
}

func UpdateUploadImg(c *fiber.Ctx) error {

	// Parse the multipart form:
	if form, err := c.MultipartForm(); err == nil {
		// => *multipart.Form
		// Get all files from "documents" key:
		files := form.File["imgs"]
		// => []*multipart.FileHeader
		id, err := strconv.Atoi(form.Value["imgs"][0])

		imgName := "IMG_21_" + strconv.Itoa(id) + ".png"
		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "image/type"

			imageCheck := strings.Split(file.Header["Content-Type"][0], "/")
			if imageCheck[0] != "image" {
				return err
			}
			// Save the files to disk:
			if err := c.SaveFile(file, fmt.Sprintf("../frontend/public/img/%s", imgName)); err != nil {
				return err
			}
		}
		return err
	} else {
		return err
	}
}
