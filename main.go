package main

import (
	"github.com/gin-gonic/gin"

	"gin-sql-mysql-tmpl/app/handler"
)

func main() {
	r := gin.Default()

	handler.NewRouter(r)

	r.Run(":8080")
}
