package products

import (
	"context"
	pb "go/service1/src/protos"
	"go/service1/src/repository"
)

type ProductService interface {
	SingelProduct(ctx context.Context, input *pb.RequestProduct) (*pb.Product, error)
	ProductsRPC(ctx context.Context, input *pb.User) (*pb.Products, error)
}

type productsImpl struct {
	Repository repository.ProductRepository
	products   *pb.Products
	product    *pb.Product
}

func NewProductService() ProductService {
	return &productsImpl{
		Repository: repository.NewProductRepo(),
	}
}

func (s *productsImpl) SingelProduct(ctx context.Context, input *pb.RequestProduct) (*pb.Product, error) {
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

func (s *productsImpl) ProductsRPC(ctx context.Context, input *pb.User) (*pb.Products, error) {
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
