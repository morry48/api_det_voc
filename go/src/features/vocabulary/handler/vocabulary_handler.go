package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"nothing-behind.com/sample_gin/features/vocabulary/usecase"
)

func ListVocabulariesHandler(uc usecase.ListVocabularies) gin.HandlerFunc {
	type Request struct {
		level string `form:"level"`
	}
	return func(c *gin.Context) {
		var err error
		var req Request
		// リクエストのバリデーションが必要になった場合はここでリクエストの内容をチェックする
		input := usecase.ListInput{
			Level: &req.level,
		}
		output, err := uc.Exec(&input)

		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, output)
	}

}
