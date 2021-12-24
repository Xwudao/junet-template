package main

import (
	"github.com/Xwudao/junet/app"
	"github.com/gin-gonic/gin"
)

func main() {
	r := app.NewApp(
		app.SetUse("junet"),
		app.SetShort("junet project template"),
		app.SetLong("junet project template"),
	)

	r.GET("/", func(c *gin.Context) {
		app.SendJson(c, app.OkRtn("hello junet!"))
	})
	panic(r.Start(3000))
}
