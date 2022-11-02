package repository

import (
	"errors"
	"go/service1/src/entities"
	"go/service1/src/http"
	"go/service1/src/listener/postgres"
	"log"
	"time"

	pb "go/service1/src/protos"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Products(input *pb.User) ([]entities.Product, error)
	Product(input *pb.RequestProduct) (*entities.Product, error)
	InsertProduct(input *pb.DataInsertProduct) (*entities.ResponseInsertProduct, error)
	DeleteProduct(i *pb.DataDeleteProduct) (*entities.ResponseDeleteProduct, error)
	UpdateProduct(i *pb.DataUpdateProduct) (*pb.ResponseUpdateProduct, error)
}

type productImpl struct {
	h                     http.HttpRequest
	products              []entities.Product
	product               *entities.Product
	dataInsertProduct     *entities.Product
	responseInsertProduct *entities.ResponseInsertProduct
	responseDeleteProduct *entities.ResponseDeleteProduct
}

func NewProductRepo() ProductRepository {
	return &productImpl{
		h: http.NewHttpRequest(),
	}
}

func (p *productImpl) Products(input *pb.User) ([]entities.Product, error) {
	var db gorm.DB = *postgres.DB

	result := db.Find(&p.products)

	if result.Error != nil {
		log.Print(result.Error.Error())
		return nil, result.Error
	}

	return p.products, nil
}

func (p *productImpl) Product(input *pb.RequestProduct) (*entities.Product, error) {
	var db gorm.DB = *postgres.DB
	result := db.Model(&entities.Product{}).
		Where("product_id = ? AND product_name = ?", input.ProductId, input.ProductName).
		First(&p.product)

	if result.Error != nil {
		log.Print(result.Error.Error())
		return nil, result.Error
	}

	return p.product, nil
}

func (p *productImpl) InsertProduct(input *pb.DataInsertProduct) (*entities.ResponseInsertProduct, error) {
	var db gorm.DB = *postgres.DB

	// begin
	tx := db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		log.Print(tx.Error.Error())
		return nil, tx.Error
	}

	// logic
	p.dataInsertProduct = &entities.Product{
		UserId:       input.UserId,
		ProductId:    input.ProductId,
		ProductName:  input.ProductName,
		ProductInfo:  input.ProductInfo,
		ProductStock: int(input.ProductStock),
		ProductPrice: int(input.ProductPrice),
		ProductSell:  int(input.ProductSell),
		LastUpdate:   input.LastUpdate,
	}

	updateProduct := tx.Model(p.dataInsertProduct).Updates(p.dataInsertProduct).
		Where("user_id = ? AND product_id=?", input.UserId, input.ProductId)

	if updateProduct.Error != nil {
		log.Print(updateProduct.Error)
		tx.Rollback()
		return nil, updateProduct.Error
	}

	err := tx.Model(p.dataInsertProduct).Where("product_id = ?", input.ProductId).First(p.dataInsertProduct).Error

	if err != nil {
		log.Print(err)
		tx.Rollback()
		return nil, err
	}

	// comit
	comit := tx.Commit()

	if comit.Error != nil {
		log.Print(comit.Error.Error())
		return nil, comit.Error
	}

	p.responseInsertProduct = &entities.ResponseInsertProduct{
		UrlImage: p.dataInsertProduct.ProductImage,
		Status:   "ok",
	}

	return p.responseInsertProduct, nil
}

func (p *productImpl) DeleteProduct(i *pb.DataDeleteProduct) (*entities.ResponseDeleteProduct, error) {
	var db gorm.DB = *postgres.DB

	// begin
	tx := db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		log.Print(tx.Error.Error())
		return nil, tx.Error
	}

	// logic
	deleteProduct := tx.Where("user_id = ? AND product_id=?", i.UserId, i.ProductId).Delete(&p.product)

	if deleteProduct.Error != nil {
		log.Print(deleteProduct.Error.Error())
		tx.Rollback()
		return nil, deleteProduct.Error
	}

	// delete image product from directory server image
	deleteDirImage, err := p.h.DeleteImage(i.UserId, i.ProductId)

	if err != nil {
		log.Print(err.Error())
		tx.Rollback()
		return nil, err
	}

	if deleteDirImage.Message != "ok" {
		log.Print(deleteDirImage.Message)
		tx.Rollback()
		return nil, errors.New(deleteDirImage.Message)
	}

	// comit
	comit := tx.Commit()

	if comit.Error != nil {
		log.Print(comit.Error.Error())
		return nil, comit.Error
	}

	p.responseDeleteProduct = &entities.ResponseDeleteProduct{
		Status:    "ok",
		ProductId: i.ProductId,
	}

	return p.responseDeleteProduct, nil
}

func (p *productImpl) UpdateProduct(i *pb.DataUpdateProduct) (*pb.ResponseUpdateProduct, error) {
	var db gorm.DB = *postgres.DB
	var time string = time.Now().Format("20060102150405")
	// begin
	tx := db.Begin()
	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		log.Print(tx.Error.Error())
		return nil, tx.Error
	}

	// logic
	product := &entities.Product{
		UserId:       i.UserId,
		ProductId:    i.ProductId,
		ProductName:  i.ProductName,
		ProductImage: i.ProductImage,
		ProductInfo:  i.ProductInfo,
		ProductStock: int(i.ProductStock),
		ProductPrice: int(i.ProductPrice),
		ProductSell:  int(i.ProductSell),
		LastUpdate:   time,
	}

	updateProduct := tx.Model(product).Updates(product).Where("user_id = ? AND product_id=?", i.UserId, i.ProductId)

	if updateProduct.Error != nil {
		log.Print(updateProduct.Error)
		tx.Rollback()
		return nil, updateProduct.Error
	}

	// comit
	comit := tx.Commit()

	if comit.Error != nil {
		log.Print(comit.Error)
		return nil, comit.Error
	}

	i.LastUpdate = product.LastUpdate

	return &pb.ResponseUpdateProduct{
		Status:  200,
		Message: "ok",
		Product: i,
	}, nil
}
