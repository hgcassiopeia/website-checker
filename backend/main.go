package main

import (
	"fmt"

	"github.com/hgcassiopeia/website-checker/backend/checker"
	"github.com/hgcassiopeia/website-checker/backend/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		fmt.Printf("please consider enviroment variables: %s", err)
	}

	r := router.NewMyRouter()
	handler := checker.NewCheckerHandler()
	r.POST("/upload", router.NewGinHandler(handler.Upload))
	r.Run()
}
