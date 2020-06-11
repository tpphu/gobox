package main

import (
	"github.com/tpphu/gobox/app"
	"github.com/tpphu/gobox/example/handler"
)

func main() {
	app := app.NewApp(
		app.Name("test"),
		app.Description("test"),
		app.WithHTTPService(":3000"))
	app.Init()
	httpService := app.GetHTTPService()
	product  := &handler.Product{}
	app.Provide(product)
	httpService.GET("/product/:id", product.Get)
	app.Run()
}
