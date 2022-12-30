package lead

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/gofiber/fiber"
	"github.com/jesseinvent/go-fiber-crm/database"
)

type Lead struct {
	gorm.Model
	Name 	string  `json:"name"`
	Company string	`json:"company"`
	Email 	string 	`json:"email"`
	Phone 	string	`json:"phone"`
}


func GetLeads(c *fiber.Ctx) {

	db := database.DBConn;

	var leads []Lead;

	db.Find(&leads);

	c.JSON(leads)
}

func GetLead(c *fiber.Ctx) {

	db := database.DBConn;

	id := c.Params("id");

	var lead Lead;

	db.Find(&lead, id);

	c.JSON(lead)
}

func NewLead(c *fiber.Ctx) {

	db := database.DBConn;

	lead := new(Lead);

	err := c.BodyParser(lead);

	if err != nil {
		c.Status(503).Send(err);
		return  
	}

	db.Create(&lead);

	c.JSON(lead);
}

func DeleteLead(c *fiber.Ctx) {

	db := database.DBConn;

	id := c.Params("id");

	var lead Lead;

	db.First(&lead, id);

	if lead.Name == "" {
		c.Status(400).Send("No lead found");
		return
	}

	db.Delete(&lead);

	c.Send("Lead succesfully deleted");
}