package main

import(
	"fmt"
	"github.com/jesseinvent/go-fiber-crm/lead"
	"github.com/jesseinvent/go-fiber-crm/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/lead",lead.GetLeads);
	app.Get("/api/v1/lead/:id",lead.GetLead);
	app.Post("/api/v1/lead", lead.NewLead);
	app.Delete("/api/v1/lead/:id", lead.DeleteLead);
}

func initDatabase() {

	var err error;

	database.DBConn, err = gorm.Open("sqlite3", "database/leads.sqlite");

	if err != nil {
		panic("Failed to connect database");
	}

	fmt.Println("Connection to database establised");

	database.DBConn.AutoMigrate(&lead.Lead{});

	fmt.Println("Database Migrated");
}

func main() {

	initDatabase()

	app := fiber.New();

	setUpRoutes(app);

	app.Listen(3000);

	defer database.DBConn.Close();
}