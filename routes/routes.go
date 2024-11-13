package routes

import (
	"github.com/gofiber/fiber/v2"
	"referral_system/handlers"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/register", handlers.RegisterUser)
	app.Post("/login", handlers.LoginUser)

	app.Post("/referral-code", handlers.CreateRefCode)
	app.Delete("/referral-code", handlers.DeleteRefCode)
	app.Get("/referrals/:id", handlers.GetRefCodeByReferrerId)

}
