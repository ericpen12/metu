package main

import (
	"dev_ihan/api"
	"dev_ihan/pkg"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
)

func main() {
	err := pkg.InitSettings()
	if err != nil {
		log.Fatal("cannot initialize setting, ", err)
	}
	pkg.Init()
	r := api.InitRoute()
	err = r.Run(fmt.Sprintf(":%s", viper.GetString("app.port")))
	if err != nil {
		zap.L().Error("cannot run web service, " + err.Error())
	}
}
