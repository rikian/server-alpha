package products

import (
	"context"
	"go/service1/src/entities"
	pb "go/service1/src/protos"
	"go/service1/src/repository"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var productRepository = repository.ProductRepositoryMock{
	Mock: mock.Mock{},
}

var productService = ProductsImpl{
	Repository: &productRepository,
}

func TestGetProduct(t *testing.T) {
	var ctx context.Context
	pMock := &entities.Product{
		ProductId:    "00bf59df-a928-4ad9-a35f-7e19fad101dd",
		UserId:       "00bf59df-a928-4ad9-a35f-7868768jjgj6",
		ProductName:  "Kue Bantent",
		ProductImage: "product/image",
		ProductInfo:  "product info",
		ProductStock: 200,
		ProductPrice: 200,
		ProductSell:  200,
		CreatedDate:  "12345",
		LastUpdate:   "12345",
	}

	productRepository.Mock.On("Product", &pb.RequestProduct{
		ProductName: "Kue Bantet",
		ProductId:   "00bf59df-a928-4ad9-a35f-7e19fad101dd",
	}).Return(pMock, nil)

	p, err := productService.GetProductById(ctx, &pb.RequestProduct{
		ProductName: "Kue Bantet",
		ProductId:   "00bf59df-a928-4ad9-a35f-7e19fad101dd",
	})

	require.Nil(t, err)
	require.Equal(t, pMock.ProductId, p.ProductId)
	require.Equal(t, pMock.UserId, p.UserId)
	require.Equal(t, pMock.ProductName, p.ProductName)
	require.Equal(t, pMock.ProductImage, p.ProductImage)
	require.Equal(t, pMock.ProductInfo, p.ProductInfo)
	require.Equal(t, uint32(pMock.ProductStock), p.ProductStock)
	require.Equal(t, uint32(pMock.ProductPrice), p.ProductPrice)
	require.Equal(t, uint32(pMock.ProductSell), p.ProductSell)
	require.Equal(t, pMock.CreatedDate, p.CreatedDate)
	require.Equal(t, pMock.LastUpdate, p.LastUpdate)
}

func TestGetProducts(t *testing.T) {
	var ctx context.Context
	var psMock []entities.Product
	for i := 0; i < 2; i++ {
		psMock = append(psMock, entities.Product{
			ProductId:    "00bf59df-a928-4ad9-a35f-7e19fad101dd",
			UserId:       "00bf59df-a928-4ad9-a35f-7868768jjgj6",
			ProductName:  "Kue Bantent",
			ProductImage: "product/image",
			ProductInfo:  "product info",
			ProductStock: 200,
			ProductPrice: 200,
			ProductSell:  200,
			CreatedDate:  "12345",
			LastUpdate:   "12345",
		})
	}

	productRepository.Mock.On("Products", &pb.User{
		Id: "1",
	}).Return(psMock, nil)

	ps, err := productService.GetAllProduct(ctx, &pb.User{
		Id: "1",
	})

	require.Nil(t, err)
	require.NotNil(t, ps)
	for i, p := range ps.Products {
		require.Equal(t, psMock[i].ProductId, p.ProductId)
		require.Equal(t, psMock[i].UserId, p.UserId)
		require.Equal(t, psMock[i].ProductName, p.ProductName)
		require.Equal(t, psMock[i].ProductImage, p.ProductImage)
		require.Equal(t, psMock[i].ProductInfo, p.ProductInfo)
		require.Equal(t, uint32(psMock[i].ProductStock), p.ProductStock)
		require.Equal(t, uint32(psMock[i].ProductPrice), p.ProductPrice)
		require.Equal(t, uint32(psMock[i].ProductSell), p.ProductSell)
		require.Equal(t, psMock[i].CreatedDate, p.CreatedDate)
		require.Equal(t, psMock[i].LastUpdate, p.LastUpdate)
	}
}
