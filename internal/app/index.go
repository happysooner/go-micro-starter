package app

import (
	"fmt"
	"github.com/micro/go-micro"
	"go-micro-starter/api/proto/app"
	"go-micro-starter/internal/app/handler"
	"time"
)

func Init() {

	appService := micro.NewService(
		micro.Name("app"),
		micro.Version("v1.0.0"),
		micro.RegisterInterval(15*time.Second),
		micro.RegisterTTL(time.Second*60),
	)
	appService.Init()
	app.RegisterAppServiceHandler(appService.Server(), new(handler.AppHandler))
	fmt.Println("start happy coding ")
	fmt.Println("App starting")
	err := appService.Run()
	if err != nil {
		fmt.Println("err string app service" + err.Error())
		return
	}

}
