package routes

import (
	"github.com/Xwudao/junet/app"

	v1 "github.com/Xwudao/junet-template/pkg/routes/v1"
)

func Setup(a *app.App) {
	v1HomeRoutes := v1.HomeRoutes{Engine: a}
	v1HomeRoutes.SetUpRoutes()
}
