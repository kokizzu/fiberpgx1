package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"fiberpgx1/business"
	"fiberpgx1/model"
	"fiberpgx1/presentation"
)

func main() {
	pg, err := model.ConnectPostgres(`root`, `password`, `127.0.0.1`, 5432, "root")
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Close()

	model.DoMigrate(pg)

	app := fiber.New()

	adapter := &model.Adapter{Pool: pg}
	admin := &business.Admin{Pg: adapter}

	view := &presentation.Presenter{Admin: admin}

	app.Post("/admin/article", view.AdminCreateArticle)

	log.Fatal(app.Listen(":3000"))
}
