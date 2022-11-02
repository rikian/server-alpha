package repository

import (
	"errors"
	"go/service1/src/entities"
	"go/service1/src/listener/postgres"
	"log"
	"os"
	"time"

	pb "go/service1/src/protos"

	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type UserRepository interface {
	SelectUser(input *pb.DataSelectUser) (*pb.ResponseSelectUser, error)
}

type userImpl struct {
	User  *entities.User
	Users []entities.User

	// product *entities.Product

	GrpcUser  *pb.ResponseSelectUser
	GrpcUsers *pb.ResponseSelectUsers

	DataInsertProduct     *entities.Product
	ResponseInsertProduct *pb.ResponseInsertProduct

	ResponseInsertProductId *pb.ResponseInsertProductID
}

func NewUserRepository() UserRepository {
	return &userImpl{}
}

func (u *userImpl) SelectUser(input *pb.DataSelectUser) (*pb.ResponseSelectUser, error) {
	var db gorm.DB = *postgres.DB
	u.User = &entities.User{}

	err := db.Model(u.User).
		Preload("UserStatus").
		Preload("Products").
		Where("user_id = ?", input.Id).
		Find(u.User).Error

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	if u.User.UserId == "" {
		return nil, errors.New("data not found")
	}

	u.GrpcUser = &pb.ResponseSelectUser{
		UserId:      u.User.UserId,
		UserEmail:   u.User.UserEmail,
		UserName:    u.User.UserName,
		UserImage:   u.User.UserImage,
		UserStatus:  u.User.UserStatus.Status,
		CreatedDate: u.User.CreatedDate,
		LastUpdate:  u.User.LastUpdate,
	}

	for j := 0; j < len(u.User.Products); j++ {
		u.GrpcUser.Products = append(u.GrpcUser.Products, &pb.Product{
			UserId:       u.User.Products[j].UserId,
			ProductId:    u.User.Products[j].ProductId,
			ProductName:  u.User.Products[j].ProductName,
			ProductImage: u.User.Products[j].ProductImage,
			ProductInfo:  u.User.Products[j].ProductInfo,
			ProductPrice: uint32(u.User.Products[j].ProductPrice),
			ProductStock: uint32(u.User.Products[j].ProductStock),
			ProductSell:  uint32(u.User.Products[j].ProductSell),
			CreatedDate:  u.User.Products[j].CreatedDate,
			LastUpdate:   u.User.Products[j].LastUpdate,
		})
	}

	return u.GrpcUser, nil
}

var SecretJwt = []byte("S4n94t_R4h4S14_BRO...")

func EncryptSession(id string, expired int) (string, error) {
	encrypt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(time.Second * time.Duration(expired)).Unix(),
	})

	token, err := encrypt.SignedString(SecretJwt)

	if err != nil {
		return "", err
	}

	return token, nil
}

func DecryptToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretJwt, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("not ok")
	}
}

func RedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func CheckDir(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
