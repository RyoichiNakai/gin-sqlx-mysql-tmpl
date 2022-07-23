package main

import (
	"github.com/gin-gonic/gin"

	"gin-sql-mysql-tmpl/app/handler"
)

// @title sample
// @version 1.0
// @description 仕様書に関する内容説明
// @termsOfService 仕様書使用する際の注意事項
// @host localhost:8080
// @BasePath /v1

func main() {
	r := gin.Default()

	handler.NewRouter(r)

	r.Run(":8080")
}
