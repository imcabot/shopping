package api_helper

import (
	"github.com/gin-gonic/gin"
	"shopping/utils/pagination"
)

var userIdText = "userId"

func GetUserId(g *gin.Context) uint {
	return uint(pagination.ParseInt(g.GetString(userIdText), -1))
}
