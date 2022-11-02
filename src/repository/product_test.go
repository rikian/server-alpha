package repository

import (
	"crypto"
	"encoding/hex"
	"go/service1/src/config"
	"log"
	"strings"

	pb "go/service1/src/protos"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

var initProduct = NewProductRepo()
var auth = NewAuthRepository()

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Print(err.Error())
		log.Fatalf("Error when loading .env file")
	}

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

func TestLogin(t *testing.T) {
	res, err := auth.LoginUser(&pb.DataLogin{
		Email:      "bangtole@gmail.com",
		Password:   sha256("12345"),
		RememberMe: true,
	})

	require.Nil(t, err)
	require.NotNil(t, res)
	require.Equal(t, 3, len(strings.Split(res.Session, ".")))
}

func TestGetSession(t *testing.T) {
	res, err := auth.SelectSessionUserById(&pb.DataSession{
		Id: "3efcafba-84a8-4bec-aab1-46aaa1736e92",
	})

	require.Nil(t, err)
	require.NotNil(t, res)
	require.Equal(t, 3, len(strings.Split(res.UserSession, ".")))
}

func TestRegister(t *testing.T) {
	res, err := auth.RegisterUser(&pb.DataRegister{
		UserEmail:    "",
		UserName:     "",
		UserPassword: "",
	})

	require.Nil(t, err)
	require.NotNil(t, res)
}

func sha256(text string) string {
	algorithm := crypto.SHA256.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}
