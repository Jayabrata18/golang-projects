package main

import (
	"fmt"

	"github.com/Jayabrata18/golang-projects/go-fiber-crm/database"
	"github.com/Jayabrata18/golang-projects/go-fiber-crm/lead"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func intDatabase() {
	var err error
	database.DBConn, err = gorm.Open("splite3", "leads.db")
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("database is connected")
	databse.DBConn.AutoMigrate(&lead.lead{})
	fmt.Println("database Migrated")

}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.listen(3000)
	defer database.DBConn.Close()

}
