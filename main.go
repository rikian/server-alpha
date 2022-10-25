package main

import (
	"go/service1/src"
	"go/service1/src/config"
)

func main() {
	src.ListenAndServe(config.Address)
}
