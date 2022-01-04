package gin

import (
	"github.com/gin-gonic/gin"

	"github.com/koderover/zadig/pkg/microservice/aslan/core/collaboration/service"
	internalhandler "github.com/koderover/zadig/pkg/shared/handler"
)

func GetCollaborationNew() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := internalhandler.NewContext(c)
		defer func() { internalhandler.JSONResponse(c, ctx) }()

		projectName := c.Query("projectName")
		if projectName == "" {
			c.Next()
		}
		newResp, err := service.GetCollaborationNew(projectName, ctx.UserID, ctx.UserName, ctx.Logger)
		if err != nil {
			ctx.Err = err
			c.Abort()
		}
		if len(newResp.Product) == 0 && len(newResp.Workflow) == 0 {
			c.Next()
		} else {
			ctx.Resp = newResp
			c.Abort()
		}
	}
}
