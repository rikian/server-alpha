package repository

import (
	"context"
	"errors"
	"go/service1/src/config"
	"go/service1/src/entities"
	"go/service1/src/http"
	"log"
	"os"
	"time"

	pb "go/service1/src/protos"

	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func InitUserRepo() *userRepo {
	return &userRepo{}
}

type userRepo struct {
	User  *entities.User
	Users []entities.User

	product *entities.Product

	GrpcUser  *pb.ResponseSelectUser
	GrpcUsers *pb.ResponseSelectUsers

	DataInsertProduct     *entities.Product
	ResponseInsertProduct *pb.ResponseInsertProduct

	ResponseInsertProductId *pb.ResponseInsertProductID
}

func (u *userRepo) InsertProductId(ctx context.Context, input *pb.DataProduct) (*pb.ResponseInsertProductID, error) {
	var db gorm.DB = *config.DB

	u.DataInsertProduct = &entities.Product{
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
	insertProduct := tx.Create(&u.DataInsertProduct)

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

	u.ResponseInsertProductId = &pb.ResponseInsertProductID{
		Status: "ok",
		Error:  "null",
	}

	return u.ResponseInsertProductId, nil
}

func (u *userRepo) SelectUser(ctx context.Context, input *pb.DataSelectUser) (*pb.ResponseSelectUser, error) {
	var db gorm.DB = *config.DB
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

func (u *userRepo) InsertProduct(ctx context.Context, i *pb.DataInsertProduct) (*pb.ResponseInsertProduct, error) {
	log.Print(i)
	var db gorm.DB = *config.DB
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

	err := tx.Model(u.product).Where("product_id = ?", i.ProductId).First(product).Error

	if err != nil {
		log.Print(err)
		tx.Rollback()
		return nil, err
	}

	// comit
	comit := tx.Commit()

	if comit.Error != nil {
		log.Print(comit.Error)
		return nil, comit.Error
	}

	i.LastUpdate = product.LastUpdate
	i.ProductImage = product.ProductImage

	return &pb.ResponseInsertProduct{
		Status:  200,
		Message: "ok",
		Product: i,
	}, nil
}

func (u *userRepo) DeleteProduct(ctx context.Context, i *pb.DataDeleteProduct) (*pb.ResponseDeleteProduct, error) {
	var db gorm.DB = *config.DB

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
	deleteProduct := tx.Where("user_id = ? AND product_id=?", i.UserId, i.ProductId).Delete(&u.product)

	if deleteProduct.Error != nil {
		log.Print(deleteProduct.Error.Error())
		tx.Rollback()
		return nil, deleteProduct.Error
	}

	// delete image product from directory server image
	deleteDirImage, err := http.DeleteImage(i.UserId, i.ProductId)

	if err != nil {
		log.Print(err.Error())
		tx.Rollback()
		return nil, err
	}

	if deleteDirImage.Message != "ok" {
		log.Print(deleteDirImage.Message)
		tx.Rollback()
		return nil, err
	}

	// comit
	comit := tx.Commit()

	if comit.Error != nil {
		log.Print(comit.Error.Error())
		return nil, comit.Error
	}

	return &pb.ResponseDeleteProduct{
		Status:    200,
		Message:   "ok",
		ProductId: i.ProductId,
	}, nil
}

func (u *userRepo) UpdateProduct(ctx context.Context, i *pb.DataUpdateProduct) (*pb.ResponseUpdateProduct, error) {
	var db gorm.DB = *config.DB
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

/*-----------------------AUTH SECTION-----------------------------*/
// register
func (u *userRepo) RegisterUser(ctx context.Context, input *pb.DataRegister) (*pb.ResponseRegister, error) {
	var db gorm.DB = *config.DB
	var time string = time.Now().Format("20060102150405")

	u.User = &entities.User{
		UserId:       uuid.New().String(),
		UserEmail:    input.UserEmail,
		UserName:     input.UserName,
		UserPassword: input.UserPassword,
		UserImage:    "default.svg",
		UserSession:  "12345",
		UserStatusId: 1,
		CreatedDate:  time,
		LastUpdate:   time,
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
	createUser := tx.Create(&u.User)

	if createUser.Error != nil {
		log.Print(createUser.Error.Error())
		tx.Rollback()
		return nil, createUser.Error
	}

	// create directory image
	createDirMedia, err := http.CreateDirectoryImage(u.User.UserId)

	if err != nil {
		log.Print(err.Error())
		tx.Rollback()
		return nil, err
	}

	if createDirMedia.Message != "ok" {
		log.Print(createDirMedia.Message)
		tx.Rollback()
		return nil, errors.New(createDirMedia.Message)
	}

	// comit
	comit := tx.Commit()

	if comit.Error != nil {
		log.Print(comit.Error.Error())
		return nil, comit.Error
	}

	return &pb.ResponseRegister{
		Status:  200,
		Message: "ok",
	}, nil
}

// login
func (u *userRepo) LoginUser(ctx context.Context, input *pb.DataLogin) (*pb.ResponseLogin, error) {
	var db gorm.DB = *config.DB
	u.User = &entities.User{}

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
	selectUser := tx.Select("user_id").
		Where("user_email = ? AND user_password = ?", input.Email, input.Password).
		First(u.User)

	if selectUser.Error != nil {
		log.Print(selectUser.Error.Error())
		return nil, selectUser.Error
	}

	var timeDuration int

	if input.RememberMe {
		timeDuration = 31536000
	} else {
		timeDuration = 1800
	}

	sessionUser, err := EncryptSession(u.User.UserId, timeDuration)

	if err != nil {
		return nil, err
	}

	u.User.UserSession = sessionUser
	u.User.RememberMe = input.RememberMe
	u.User.LastUpdate = time.Now().Format("20060102150405")

	insertSession := db.Model(u.User).
		Select("UserSession", "LastUpdate", "RememberMe").Updates(u.User).
		Where("user_email = ? AND user_password = ?", input.Email, input.Password)

	if insertSession.Error != nil {
		return nil, err
	}

	// comit
	comit := tx.Commit()

	if comit.Error != nil {
		log.Print(comit.Error.Error())
		return nil, comit.Error
	}

	// store session in to redis
	// rDB := RedisClient()
	// setSessionInToRedis := rDB.Set(u.User.UserId, u.User.UserSession, time.Duration(18000)*time.Second)

	// if err := setSessionInToRedis.Err(); err != nil {
	// 	log.Print(err.Error())
	// }

	return &pb.ResponseLogin{
		Status:  200,
		Message: "ok",
		UserId:  u.User.UserId,
		Session: u.User.UserSession,
	}, nil
}

// select session user
func (u *userRepo) SelectSessionUserById(ctx context.Context, input *pb.DataSession) (*pb.ResponseSession, error) {
	var db gorm.DB = *config.DB
	u.User = &entities.User{}

	sessionUser := db.Select("user_session", "remember_me").
		Where("user_id = ?", input.Id).
		First(u.User)

	if sessionUser.Error != nil {
		log.Print(sessionUser.Error.Error())
		return nil, sessionUser.Error
	}

	return &pb.ResponseSession{
		UserSession: u.User.UserSession,
		RememberMe:  u.User.RememberMe,
	}, nil
}

var SecretJwt = []byte(os.Getenv("JWT_SECRET"))

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

func UpdateSession(userId, session string) {}

// func (u *userRepo) SelectUsers(ctx context.Context, input *pb.Limit) (*pb.DataUsers, error) {
// 	var db gorm.DB = *config.DB

// 	u.User = &entities.User{}
// 	u.Users = []entities.User{}

// 	err := db.Model(u.User).Preload("UserStatus").Preload("Products").Find(&u.Users)

// 	if err.Error != nil {
// 		log.Print(err.Error.Error())
// 		return nil, errors.New("something wrong")
// 	}

// 	u.GrpcUsers = &pb.DataUsers{}

// 	for i := 0; i < len(u.Users); i++ {
// 		u.GrpcUser = &pb.DataUser{
// 			UserId:      u.Users[i].UserId,
// 			UserEmail:   u.Users[i].UserEmail,
// 			UserName:    u.Users[i].UserName,
// 			UserImage:   u.Users[i].UserImage,
// 			UserStatus:  u.Users[i].UserStatus.Status,
// 			CreatedDate: u.Users[i].CreatedDate,
// 			LastUpdate:  u.Users[i].LastUpdate,
// 		}

// 		for j := 0; j < len(u.Users[i].Products); j++ {
// 			u.GrpcUser.Products = append(u.GrpcUser.Products, &pb.Product{
// 				UserId:       u.Users[i].Products[j].UserId,
// 				ProductId:    u.Users[i].Products[j].ProductId,
// 				ProductName:  u.Users[i].Products[j].ProductName,
// 				ProductPrice: uint32(u.Users[i].Products[j].ProductPrice),
// 				ProductStock: uint32(u.Users[i].Products[j].ProductStock),
// 				CreatedDate:  u.Users[i].Products[j].CreatedDate,
// 				LastUpdate:   u.Users[i].Products[j].LastUpdate,
// 			})
// 		}

// 		u.GrpcUsers.Data = append(u.GrpcUsers.Data, u.GrpcUser)
// 	}

// 	return u.GrpcUsers, nil
// }

// var db gorm.DB = *config.DB
// 	var time string = time.Now().Format("20060102150405")

// 	u.DataInsertProduct = &entities.Product{
// 		UserId:       i.UserId,
// 		ProductId:    uuid.New().String(),
// 		ProductName:  i.ProductName,
// 		ProductImage: i.ProductImage,
// 		ProductInfo:  i.ProductInfo,
// 		ProductStock: int(i.ProductStock),
// 		ProductPrice: int(i.ProductPrice),
// 		ProductSell:  int(i.ProductSell),
// 		CreatedDate:  time,
// 		LastUpdate:   time,
// 	}

// 	// begin
// 	tx := db.Begin()

// 	defer func() {
// 		r := recover()
// 		if r != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	if tx.Error != nil {
// 		log.Print(tx.Error.Error())
// 		return nil, tx.Error
// 	}

// logic
// insertProduct := tx.Create(&u.DataInsertProduct)

// if insertProduct.Error != nil {
// 	log.Print(insertProduct.Error.Error())
// 	tx.Rollback()
// 	return nil, insertProduct.Error
// }

// save image

// comit
// comit := tx.Commit()

// if comit.Error != nil {
// 	log.Print(comit.Error.Error())
// 	return nil, comit.Error
// }

// i.ProductId = u.DataInsertProduct.ProductId
// i.CreatedDate = u.DataInsertProduct.CreatedDate
// i.LastUpdate = u.DataInsertProduct.LastUpdate

// return &pb.ResponseInsertProduct{
// 	Status:  200,
// 	Message: "ok",
// 	Product: i,
// }, nil
