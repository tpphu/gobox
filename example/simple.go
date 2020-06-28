package main

import (
	"flag"
	"github.com/tpphu/gobox/app"
	"github.com/tpphu/gobox/service/http"
)

func main() {
	var raz string
	var raz1 string
	flag.StringVar(&raz, "raz-value", "bar", "set the raz")
	flag.StringVar(&raz1, "raz1-value", "foo", "set the raz1")

	app := app.NewApp(
		app.Name("test"),
		app.Description("test"),
		//app.WithHTTPService(":3000"),
	)
	//app.Init()
	httpService := http.NewGinService(app.Name)
	//product  := &handler.Product{}
	//app.Provide(product)
	//httpService.GET("/product/:id", product.Get)
	app.AddService(httpService)
	app.Run()
}
