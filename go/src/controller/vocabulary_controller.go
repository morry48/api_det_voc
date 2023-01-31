package vocabulary

import (
	"fmt"
	"nothing-behind.com/sample_gin/usecase"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (pc Controller) Index(c *gin.Context) {
	var u vocabulary.Usecase
	p, err := u.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}
