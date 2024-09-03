package test

import (
	"github.com/Sabyradinov/golang-hex-layout/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoggerMock struct {
	TestPanic bool
}

func (l LoggerMock) HttpPanicHandler(c *gin.Context, recovered interface{}) {
	if l.TestPanic {
		panic(common.PanicErrorMsg)
	}
}

func (l LoggerMock) ErrorWithCode(ctx *gin.Context, action string, code int, message string, opt *common.LogOptions) {
	if l.TestPanic {
		panic(common.PanicErrorMsg)
	}
	ctx.JSON(http.StatusUnprocessableEntity, common.PanicErrorMsg)
}

func (l LoggerMock) Error(action string, code int, message string, opt *common.LogOptions) {
	if l.TestPanic {
		panic(common.PanicErrorMsg)
	}
}

func (l LoggerMock) Warn(action, message string, opt *common.LogOptions) {
	if l.TestPanic {
		panic(common.PanicErrorMsg)
	}
}

func (l LoggerMock) Debug(action, message string, opt *common.LogOptions) {
	if l.TestPanic {
		panic(common.PanicErrorMsg)
	}
}

func (l LoggerMock) Info(action, message string, opt *common.LogOptions) {
	if l.TestPanic {
		panic(common.PanicErrorMsg)
	}
}
