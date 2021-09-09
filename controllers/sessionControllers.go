package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GetSession 获取session
func GetSession(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	loginuser := session.Get("userLogin")

	if loginuser != nil {
		return true
	} else {
		return false
	}
}

func GetUserName(ctx *gin.Context) interface{} {
	session := sessions.Default(ctx)
	loginuser := session.Get("userLogin")

	return loginuser
}
