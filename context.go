package httpserver

import (
	"github.com/gin-gonic/gin"

	"github.com/yeencloud/lib-shared"
)

func SetRequestContextValue(ctx *gin.Context, key, value interface{}) {
	// Get the existing context from the request
	reqCtx, _ := ctx.Get("shared")
	cc, ok := reqCtx.(*shared.Context)
	if ok {
		cc.WithValue(key, value)
	}
}

func CreateSharedRequest(ctx *gin.Context) {
	// Create a new context
	newCtx := shared.NewCustomContext()
	// Replace the request's context with the new context
	ctx.Set("shared", newCtx)
}
