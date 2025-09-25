package contracts

import (
	"fiber-gorm-channel-ecommerce/src/domain/model/Dto"
	"fiber-gorm-channel-ecommerce/src/domain/model/entity"
)

type CategoryContract interface {
	Create(dto Dto.CreateCategoryDto) error
	getAll() ([]entity.Category, error)
	FindById(id uint) (entity.Category, error)
}
