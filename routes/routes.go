package routes

import (
	"dashboard/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Get("/user", controllers.User)
	app.Post("/logout", controllers.Logout)

	app.Post("/create_client", controllers.CreateClient)
	app.Post("/delete_client", controllers.DeleteClientData)
	app.Post("/update_client", controllers.UpdateClientData)
	app.Post("/delete_client_offers", controllers.DeleteClientOffers)
	app.Post("/update_client_offers", controllers.UpdateClientOffers)
	app.Get("/list_client", controllers.SendClientData)

	app.Post("/create_offer", controllers.CreateOffer)
	app.Post("/delete_offer", controllers.DeleteOfferData)
	app.Post("/update_offer", controllers.UpdateOfferData)
	app.Get("/list_offer", controllers.SendOfferData)
	app.Get("/price_offer", controllers.FilterPriceDataMonth)

	app.Get("/list_common", controllers.SendCommonCostData)
	app.Post("/create_common", controllers.CreateCommonCost)
	app.Post("/delete_common", controllers.DeleteCommonData)
	app.Post("/update_common", controllers.UpdateCommonCost)
	app.Get("/price_cost", controllers.FilterCostDataMonth)

	app.Get("/list_staff", controllers.SendStaffCostData)
	app.Post("/create_staff", controllers.CreateStaffCostData)
	app.Post("/update_staff", controllers.UpdateStaffCostData)
	app.Post("/delete_staff", controllers.DeleteStaffData)

	app.Post("/upload_pdf", controllers.UploadPdf)
	app.Post("/upload_update_pdf", controllers.UpdateUploadPdf)
	app.Post("/upload_img", controllers.UploadImg)
	app.Post("/upload_update_img", controllers.UpdateUploadImg)

	// app.Post("/test", controllers.Test)
}
