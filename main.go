package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/kriangkrai/GoFiber/Controller"
	"github.com/kriangkrai/GoFiber/Router"
)

// type MongoInstance struct {
// 	Client *mongo.Client
// 	Db     *mongo.Database
// }

// var mg MongoInstance

// const dbName = "mongodb+srv://Meeci:Meeci50026@meego.biqun.mongodb.net/Mee?retryWrites=true&w=majority"

// func Connect() error {
// 	client, err := mongo.NewClient(options.Client().ApplyURI(dbName))

// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()

// 	err = client.Connect(ctx)
// 	db := client.Database("Mee")

// 	if err != nil {
// 		return err
// 	}
// 	mg = MongoInstance{
// 		Client: client,
// 		Db:     db,
// 	}
// 	return nil
// }
func main() {

	if err := Controller.Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	api := app.Group("/api") // /api

	v1 := api.Group("/v1") // /api/v1
	v1.Get("/get/:device", Router.Get)
	v1.Get("/gets", Router.Gets)
	v1.Post("/insert", Router.Insert)
	v1.Put("/update", Router.Update)
	v1.Delete("/delete/:device", Router.Delete)

	app.Listen(":3000")
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
		//fmt.Println("No Port In Heroku" + port)
	}

	log.Fatal(app.Listen(port))
}
