package main

import (
	_ "github.com/AlexSilverson/Dojdik/docs"
	"github.com/AlexSilverson/Dojdik/src/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

// @title			Dojdik API
// @version		1.0
// @description	This is a sample swagger for Dojdik
// @termsOfService	http://swagger.io/terms/

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	//routes.Setup(app)

	app.Get("/swagger/*", swagger.HandlerDefault)
	controllers.CalcByGivenData(app)
	app.Listen(":3001")
	/*var n int
	var h float64
	fmt.Scan(&n, &h)
	n++
	a := make([]geom.Point, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i].X, &a[i].Y)
	}
	inp := entity.Input{
		N:      n,
		H:      h,
		Points: a,
	}
	str, _ := json.Marshal(inp)
	fmt.Println(string(str))*/
}
