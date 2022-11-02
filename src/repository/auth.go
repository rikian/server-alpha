package repository

import (
	"errors"
	"go/service1/src/entities"
	"go/service1/src/http"
	"go/service1/src/listener/postgres"
	pb "go/service1/src/protos"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository interface {
	LoginUser(input *pb.DataLogin) (*entities.ResponLogin, error)
	RegisterUser(input *pb.DataRegister) (*entities.ResponRegister, error)
	SelectSessionUserById(input *pb.DataSession) (*entities.ResponGetSession, error)
}

type authImpl struct {
	user *entities.User
	h    http.HttpRequest
}

func NewAuthRepository() AuthRepository {
	return &authImpl{
		h: http.NewHttpRequest(),
	}
}

func (u *authImpl) LoginUser(input *pb.DataLogin) (*entities.ResponLogin, error) {
	var db gorm.DB = *postgres.DB
	u.user = &entities.User{}

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
		First(u.user)

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

	sessionUser, err := EncryptSession(u.user.UserId, timeDuration)

	if err != nil {
		return nil, err
	}

	u.user.UserSession = sessionUser
	u.user.RememberMe = input.RememberMe
	u.user.LastUpdate = time.Now().Format("20060102150405")

	insertSession := db.Model(u.user).
		Select("UserSession", "LastUpdate", "RememberMe").Updates(u.user).
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

	return &entities.ResponLogin{
		UserId:  u.user.UserId,
		Session: u.user.UserSession,
	}, nil
}

func (u *authImpl) RegisterUser(input *pb.DataRegister) (*entities.ResponRegister, error) {
	var db gorm.DB = *postgres.DB
	var time string = time.Now().Format("20060102150405")

	u.user = &entities.User{
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
	createUser := tx.Create(&u.user)

	if createUser.Error != nil {
		log.Print(createUser.Error.Error())
		tx.Rollback()
		return nil, createUser.Error
	}

	// create directory image
	createDirMedia, err := u.h.CreateDirectoryImage(u.user.UserId)

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

	return &entities.ResponRegister{
		Status: "ok",
	}, nil
}

func (u *authImpl) SelectSessionUserById(input *pb.DataSession) (*entities.ResponGetSession, error) {
	var db gorm.DB = *postgres.DB
	u.user = &entities.User{}

	sessionUser := db.Select("user_session", "remember_me").
		Where("user_id = ?", input.Id).
		First(u.user)

	if sessionUser.Error != nil {
		log.Print(sessionUser.Error.Error())
		return nil, sessionUser.Error
	}

	return &entities.ResponGetSession{
		UserSession: u.user.UserSession,
		RememberMe:  u.user.RememberMe,
	}, nil
}
