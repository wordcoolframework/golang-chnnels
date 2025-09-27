package boot

import (
	"fiber-gorm-channel-ecommerce/src/pkg/configCore"
	"fiber-gorm-channel-ecommerce/src/pkg/databaseCore"
	"fiber-gorm-channel-ecommerce/src/pkg/qrybldr"
	"fiber-gorm-channel-ecommerce/src/pkg/routing"
)

func Serve() {
	configCore.ConfigurationSet()

	databaseCore.ConnectMysqlDB()

	qrybldr.Init(databaseCore.Connection())

	routing.Init()

	routing.RegisterRoutes()

	routing.RunServer()

}
