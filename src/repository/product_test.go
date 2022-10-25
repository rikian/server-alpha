package repository

import (
	"go/service1/src/config"

	pb "go/service1/src/protos"
	"testing"

	"github.com/stretchr/testify/require"
)

var initProduct = NewProductRepo()

func TestMain(m *testing.M) {
	config.ConnectDB()
	m.Run()
}

func TestGetProducts(t *testing.T) {
	products, err := initProduct.Products(&pb.User{})
	require.Nil(t, err)
	require.NotNil(t, products)
}

func TestGetProductFailed(t *testing.T) {
	product, err := initProduct.Product(&pb.RequestProduct{})
	require.NotNil(t, err)
	require.Nil(t, product)
}

func TestGetProductSuccess(t *testing.T) {
	product, err := initProduct.Product(&pb.RequestProduct{
		ProductName: "Kue Bantet",
		ProductId:   "00bf59df-a928-4ad9-a35f-7e19fad101dd",
	})
	require.Nil(t, err)
	require.NotNil(t, product)
}
