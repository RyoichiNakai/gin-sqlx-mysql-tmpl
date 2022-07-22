package handler

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine) {
	r.Use(newCORS())

	v1 := r.Group("/v1")
	{
		v1.GET("/helloworld", Helloworld)
	}
}

// @title テストドキュメント
// @version バージョン(1.0)
// @description 仕様書に関する内容説明
// @termsOfService 仕様書使用する際の注意事項
// @host localhost:8080
// @BasePath /v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func newCORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	})
}
