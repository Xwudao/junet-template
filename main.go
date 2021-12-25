package main

import (
	"github.com/Xwudao/junet/app"

	"github.com/Xwudao/junet-template/pkg/routes"
)

func main() {
	r := app.NewApp(
		app.SetUse("junet"),
		app.SetShort("junet project template"),
		app.SetLong("junet project template"),
	)

	routes.Setup(r)

	panic(r.Start(3000))
}
