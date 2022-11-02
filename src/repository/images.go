package repository

import (
	"go/service1/src/entities"
	"go/service1/src/listener/postgres"

	pb "go/service1/src/protos"
	"log"

	"gorm.io/gorm"
)

type ImageRepository interface {
	InsertProductId(input *pb.DataProduct) (*pb.ResponseInsertProductID, error)
}

type imageImpl struct {
	dataInsertProduct       *entities.Product
	responseInsertProductId *pb.ResponseInsertProductID
}

func NewImageRepository() ImageRepository {
	return &imageImpl{}
}

func (u *imageImpl) InsertProductId(input *pb.DataProduct) (*pb.ResponseInsertProductID, error) {
	var db gorm.DB = *postgres.DB

	u.dataInsertProduct = &entities.Product{
		UserId:       input.UserId,
		ProductId:    input.ProductId,
		ProductName:  "",
		ProductImage: input.ProductImage,
		ProductInfo:  "",
		ProductStock: 0,
		ProductPrice: 0,
		ProductSell:  0,
		CreatedDate:  input.CreatedDate,
		LastUpdate:   "",
	}

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
	insertProduct := tx.Create(&u.dataInsertProduct)

	if insertProduct.Error != nil {
		log.Print(insertProduct.Error.Error())
		tx.Rollback()
		return nil, insertProduct.Error
	}

	// comit
	comit := tx.Commit()

	if comit.Error != nil {
		log.Print(comit.Error.Error())
		return nil, comit.Error
	}

	u.responseInsertProductId = &pb.ResponseInsertProductID{
		Status: "ok",
		Error:  "null",
	}

	return u.responseInsertProductId, nil
}
