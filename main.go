package main

import (
	"fmt"

	"github.com/LebrancWorkshop/Go-Hexagonal-Architecture-Codebangkok/repository"
	"github.com/LebrancWorkshop/Go-Hexagonal-Architecture-Codebangkok/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := sqlx.Open("mysql", "root:P@ssw0rd@tcp(0.0.0.0:3306)/lingo_quest");
	if err != nil {
		fmt.Println("[ERROR] Open SQL error")
		panic(err);
	}

	userRepository := repository.NewUserRepositoryDB(db);
	userService := service.NewUserService(userRepository);

	_ = userService;

	app := fiber.New();
	app.Get("/users", func(c *fiber.Ctx) error {
		return c.SendString("Coming Soon."); 
	})

	app.Listen(":8010");

}