package http

import (
	"bytes"
	"encoding/json"
	"go/service1/src/config"
	"log"
	"net/http"
	"net/url"
)

var client = &http.Client{}

type resHttp struct {
	Message string
}

func CreateDirectoryImage(id string) (*resHttp, error) {
	statusUser := &resHttp{}
	formUserId := url.Values{}
	formUserId.Set("id", id)
	formUserId.Set("token", "12345")
	payload := bytes.NewBuffer([]byte(formUserId.Encode()))
	request, err := http.NewRequest("POST", config.HttpAddress+"/user/create-directory-image/", payload)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(statusUser)

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return statusUser, nil
}

func DeleteImage(userId, productId string) (*resHttp, error) {
	responseDeleteDirectoryImage := &resHttp{}
	formDeleteImageProduct := url.Values{}
	formDeleteImageProduct.Set("userId", userId)
	formDeleteImageProduct.Set("productId", productId)
	formDeleteImageProduct.Set("token", "12345")
	payload := bytes.NewBuffer([]byte(formDeleteImageProduct.Encode()))
	request, err := http.NewRequest("POST", config.HttpAddress+"/user/delete-directory-image/", payload)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(responseDeleteDirectoryImage)

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return responseDeleteDirectoryImage, nil
}
