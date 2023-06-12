package utils

import "github.com/gin-gonic/gin"

func JoinHandlerAtStart(handler gin.HandlerFunc, handlers ...gin.HandlerFunc) []gin.HandlerFunc {
	return append([]gin.HandlerFunc{handler}, handlers...)
}
