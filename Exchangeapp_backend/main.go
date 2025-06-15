package main

import (
	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/config"
	"github.com/DavidLee0620/ExchangeApp/Exchangeapp_backend/route"
)

func main() {
	config.InitConfig()
	r := route.SetupRouter()

	r.Run(config.AppConfig.App.Port) // listen and serve
}
