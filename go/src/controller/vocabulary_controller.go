package vocabulary

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nothing-behind.com/sample_gin/usecase"
)

type Controller struct{}

func (pc Controller) Index(c *gin.Context) {
	level := c.Query("level")
	var u vocabulary.Usecase
	input := vocabulary.ListInput{
		Level: level,
	}
	vocList, err := u.GetAll(&input)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, vocList)
	}
}
