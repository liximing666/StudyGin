package midware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type GlobleMidware struct {

}

func (globle GlobleMidware) Show(context *gin.Context)  {
	fmt.Println("globle midware")
}
