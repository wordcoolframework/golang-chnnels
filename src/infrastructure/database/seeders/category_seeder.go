package seeders

import (
	"fiber-gorm-channel-ecommerce/src/domain/model/entity"
	"fiber-gorm-channel-ecommerce/src/pkg"
	"fiber-gorm-channel-ecommerce/src/pkg/databaseCore"
	"fiber-gorm-channel-ecommerce/src/pkg/qrybldr"
	"fmt"
	"strconv"
)

type CategorySeeder struct {
}

func (c *CategorySeeder) GetName() string {
	return "CategorySeeder"
}

func (c *CategorySeeder) Run() error {
	databaseCore.ConnectMysqlDB()
	db := databaseCore.Connection()

	if db == nil {
		return fmt.Errorf("DB connection is nil")
	}

	qb := qrybldr.Instance(db)

	for i := 0; i < 10; i++ {

		category := entity.Category{
			Name: "Category" + strconv.Itoa(i),
		}

		if err := qb.Model(&entity.Category{}).Create(&category); err != nil {
			return err
		}
	}

	table := pkg.GetTableNameWithModel(entity.Category{})

	fmt.Println("✅ Seed Data on Table", table)

	return nil
}
