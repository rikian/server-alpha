package repository

import (
	"go/service1/src/config"
	"go/service1/src/entities"
	"log"

	pb "go/service1/src/protos"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Products(input *pb.User) ([]entities.Product, error)
	Product(input *pb.RequestProduct) (*entities.Product, error)
}

type productImpl struct {
	products []entities.Product
	product  *entities.Product
}

func NewProductRepo() ProductRepository {
	return &productImpl{}
}

func (p *productImpl) Products(input *pb.User) ([]entities.Product, error) {
	var db gorm.DB = *config.DB

	result := db.Find(&p.products)

	if result.Error != nil {
		log.Print(result.Error.Error())
		return nil, result.Error
	}

	return p.products, nil
}

func (p *productImpl) Product(input *pb.RequestProduct) (*entities.Product, error) {
	var db gorm.DB = *config.DB
	result := db.Model(&entities.Product{}).Where("product_id = ? AND product_name = ?", input.ProductId, input.ProductName).First(&p.product)

	if result.Error != nil {
		log.Print(result.Error.Error())
		return nil, result.Error
	}

	return p.product, nil
}
