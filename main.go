package main

import (
	"github.com/Xwudao/junet/app"
	"github.com/Xwudao/junet/confx"
	"github.com/Xwudao/junet/logx"

	"github.com/Xwudao/junet-template/pkg/models"
	"github.com/Xwudao/junet-template/pkg/routes"
)

func main() {
	logx.Init()
	confx.Init()
	models.InitDB()

	r := app.NewApp(
		app.SetUse("junet"),
		app.SetShort("junet project template"),
		app.SetLong("junet project template"),
	)

	routes.Setup(r)

	err := r.Start(3000)
	if err != nil {
		logx.Panic(err)
	}
}
