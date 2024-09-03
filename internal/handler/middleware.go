package handler

import (
	"github.com/Sabyradinov/golang-hex-layout/common"
	"github.com/Sabyradinov/golang-hex-layout/internal/domain/port/logger"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	log logger.AppLogger
}

func NewMiddleware(log logger.AppLogger) *Middleware {
	return &Middleware{log: log}
}

// Locale middleware for setting locale
func (m *Middleware) Locale(ctx *gin.Context) {
	locale := ctx.GetHeader("Accept-Language")
	if len(locale) == 0 {
		locale = common.DefaultLang
	}
	if locale == "kk" {
		locale = "KAZ"
	} else if locale == "en" {
		locale = "ENG"
	}

	ctx.Set("Locale", locale)
	ctx.Next()
}

// UserTokenVerify middleware for user token verification
func (m *Middleware) UserTokenVerify(ctx *gin.Context) {
	ti, exists := ctx.Get("TokenInfo")
	if !exists {
		m.log.Error("UserTokenVerify", 401, "Token not found in context", nil)
		ctx.JSON(401, common.BaseResponse{
			Code:    401,
			Message: "Token not found in context",
		})
		ctx.Abort()
		return
	}
	ctx.Set("TokenInfo", ti)
	ctx.Next()
}
