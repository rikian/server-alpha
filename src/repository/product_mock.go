package repository

import (
	"go/service1/src/entities"
	pb "go/service1/src/protos"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (m *ProductRepositoryMock) Product(input *pb.RequestProduct) (*entities.Product, error) {
	args := m.Mock.Called(input)

	if args.Get(0) == nil {
		return nil, nil
	} else {
		product := args.Get(0).(*entities.Product)
		return product, nil
	}
}

func (m *ProductRepositoryMock) Products(input *pb.User) ([]entities.Product, error) {
	args := m.Mock.Called(input)

	if args.Get(0) == nil {
		return nil, nil
	} else {
		products := args.Get(0).([]entities.Product)
		return products, nil
	}
}
