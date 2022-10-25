package http

import (
	"log"
	"testing"
)

func TestXxx(t *testing.T) {
	resultDeleteImg, err := DeleteImage("12345", "taik.jpg")

	if err != nil {
		t.Fatalf(err.Error())
	}

	log.Print(resultDeleteImg)
}
