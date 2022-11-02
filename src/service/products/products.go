package products

import (
	"context"
	pb "go/service1/src/protos"
	"go/service1/src/repository"
	"time"
)

type ProductService interface {
	GetProductById(ctx context.Context, input *pb.RequestProduct) (*pb.Product, error)
	GetAllProduct(ctx context.Context, input *pb.User) (*pb.Products, error)
	InsertProduct(ctx context.Context, input *pb.DataInsertProduct) (*pb.ResponseInsertProduct, error)
	DeleteProduct(ctx context.Context, input *pb.DataDeleteProduct) (*pb.ResponseDeleteProduct, error)
	UpdateProduct(ctx context.Context, input *pb.DataUpdateProduct) (*pb.ResponseUpdateProduct, error)
}

type ProductsImpl struct {
	pb.UnimplementedProductRPCServer
	Repository repository.ProductRepository
	products   *pb.Products
	product    *pb.Product
}

func (s *ProductsImpl) GetProductById(ctx context.Context, input *pb.RequestProduct) (*pb.Product, error) {
	product, err := s.Repository.Product(input)

	if err != nil {
		return nil, err
	}

	s.product = &pb.Product{
		UserId:       product.UserId,
		ProductId:    product.ProductId,
		ProductName:  product.ProductName,
		ProductStock: uint32(product.ProductStock),
		ProductPrice: uint32(product.ProductPrice),
		CreatedDate:  product.CreatedDate,
		LastUpdate:   product.LastUpdate,
		ProductImage: product.ProductImage,
		ProductSell:  uint32(product.ProductSell),
		ProductInfo:  product.ProductInfo,
	}

	return s.product, nil
}

func (s *ProductsImpl) GetAllProduct(ctx context.Context, input *pb.User) (*pb.Products, error) {
	products, err := s.Repository.Products(input)

	if err != nil {
		return nil, err
	}

	s.products = &pb.Products{}

	for i := 0; i < len(products); i++ {
		s.products.Products = append(s.products.Products, &pb.Product{
			UserId:       products[i].UserId,
			ProductId:    products[i].ProductId,
			ProductName:  products[i].ProductName,
			ProductStock: uint32(products[i].ProductStock),
			ProductPrice: uint32(products[i].ProductPrice),
			ProductSell:  uint32(products[i].ProductSell),
			ProductInfo:  products[i].ProductInfo,
			ProductImage: products[i].ProductImage,
			CreatedDate:  products[i].CreatedDate,
			LastUpdate:   products[i].LastUpdate,
		})
	}

	return s.products, nil
}

func (s *ProductsImpl) InsertProduct(ctx context.Context, input *pb.DataInsertProduct) (*pb.ResponseInsertProduct, error) {
	var time string = time.Now().Format("20060102150405")
	input.LastUpdate = time
	insertProduct, err := s.Repository.InsertProduct(input)

	if err != nil {
		return nil, err
	}

	input.ProductImage = insertProduct.UrlImage

	return &pb.ResponseInsertProduct{
		Status:  200,
		Message: "ok",
		Product: input,
	}, nil
}

func (s *ProductsImpl) DeleteProduct(ctx context.Context, input *pb.DataDeleteProduct) (*pb.ResponseDeleteProduct, error) {
	deleteProduct, err := s.Repository.DeleteProduct(input)

	if err != nil {
		return nil, err
	}

	return &pb.ResponseDeleteProduct{
		Status:    200,
		Message:   "ok",
		ProductId: deleteProduct.ProductId,
	}, nil
}

func (s *ProductsImpl) UpdateProduct(ctx context.Context, input *pb.DataUpdateProduct) (*pb.ResponseUpdateProduct, error) {
	updateProduct, err := s.Repository.UpdateProduct(input)

	if err != nil {
		return nil, err
	}

	return updateProduct, nil
}
