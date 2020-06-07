package main

import (
	"fmt"
	"github.com/tpphu/gobox"
)

func main() {
	app := gobox.Default()
	app.Load()
	app.Up()
}
