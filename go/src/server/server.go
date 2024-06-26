package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"nothing-behind.com/sample_gin/features/vocabulary/handler"
	"nothing-behind.com/sample_gin/features/vocabulary/infra/postgres"
	"time"
)

func Init() {
	r := router()
	err := r.Run()
	if err != nil {
		log.Fatalf("%+v", err)
		return
	}
	postgres.Close()
}

func router() *gin.Engine {
	r := gin.Default()

	// ここからCorsの設定
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"*",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Origin",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: false,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	v := r.Group("/vocabularies")
	{
		db, err := postgres.New()
		if err != nil {
			log.Fatal("fail init database")
		}
		vocabularyListUsecase := InitVocabularyList(db)
		listVocabulariesHandler := handler.ListVocabulariesHandler(vocabularyListUsecase)

		v.GET("/", listVocabulariesHandler)
	}

	return r
}
