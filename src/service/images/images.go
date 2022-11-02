package images

import (
	"context"
	pb "go/service1/src/protos"
	"go/service1/src/repository"
)

type ImageService interface {
	InsertProductId(ctx context.Context, input *pb.DataProduct) (*pb.ResponseInsertProductID, error)
}

type ImageImpl struct {
	pb.UnimplementedImageRPCServer
	Repository repository.ImageRepository
}

func (i *ImageImpl) InsertProductId(ctx context.Context, input *pb.DataProduct) (*pb.ResponseInsertProductID, error) {
	res, err := i.Repository.InsertProductId(input)

	if err != nil {
		return nil, err
	}

	return res, nil
}
