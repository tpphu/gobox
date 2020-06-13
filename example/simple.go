package main

import(
	"flag"
	"github.com/tpphu/gobox/app"
	"github.com/tpphu/gobox/example/handler"
)

func main() {
	var raz string
	var raz1 string
	flag.StringVar(&raz, "raz-value", "bar", "set the raz")
	flag.StringVar(&raz1, "raz1-value", "foo", "set the raz1")

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
