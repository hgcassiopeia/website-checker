package main

import (
	"github.com/hgcassiopeia/website-checker/backend/checker"
	"github.com/hgcassiopeia/website-checker/backend/router"
)

func main() {
	r := router.NewMyRouter()
	handler := checker.NewCheckerHandler()
	r.POST("/upload", router.NewGinHandler(handler.Upload))
	r.Run(":3000")
}
