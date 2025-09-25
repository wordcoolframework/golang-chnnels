package provider

import (
	"fiber-gorm-channel-ecommerce/src/infrastructure/database/seeders"
	"fiber-gorm-channel-ecommerce/src/pkg/databaseCore"
)

var SeederRegister = map[string]databaseCore.Seeder{
	"CategorySeeder": &seeders.CategorySeeder{},
}
