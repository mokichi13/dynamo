package main

import (
	"fmt"

	"github.com/mokichi13/dynamo/internal/service"
)

func main() {
	sv, err := service.NewService()
	if err != nil {
		fmt.Printf("service init failed. err:%+v\n", err)
		return
	}

	sv.Run()
}
