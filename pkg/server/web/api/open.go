package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"shahaohuo.com/shahaohuo/pkg/server/service"
	"shahaohuo.com/shahaohuo/pkg/server/web/token"
	"time"
)

type Payload struct {
	Id       string `json:"id" binding:"required,max=12,min=1"`
	Password string `json:"password" binding:"required,max=18,min=8"`
}

type LoginResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expiredAt"`
}

func Login(ctx *gin.Context) {
	if payload, validated := payload(ctx); validated {
		if u, e := service.Login(payload.Id, payload.Password); e != nil {
			if service.IsServiceError(e) {
				ctx.JSON(http.StatusInternalServerError, newBadResp(-1, e.Error()))
				return
			}
			// password error
			if service.IsUsernameOrPasswordError(e) {
				ctx.JSON(http.StatusUnauthorized, newBadResp(1, e.Error()))
			}
		} else {
			days, _ := time.ParseDuration("360h")
			expiredAt := time.Now().Add(days).Unix()
			if t, e := token.Sign(u.Id, expiredAt); e == nil {
				ctx.SetCookie("token", t, 30*24*60*60, "/", "", false, true)
				ctx.JSON(200, &LoginResponse{
					Id:        u.Id,
					Name:      u.Name,
					Token:     t,
					ExpiredAt: expiredAt,
				})
			} else {
				logrus.Error(e)
				ctx.JSON(http.StatusInternalServerError, newBadResp(-1, "server error"))
			}
		}
	}
}

func Register(ctx *gin.Context) {
	if payload, validated := payload(ctx); validated {
		if _, e := service.Register(payload.Id, payload.Password); e == nil {
			ok(ctx, "success")
		} else {
			if service.IsResourcesExist(e) {
				bad(ctx, newBadResp(1, "用户已注册"))
				return
			}
			if service.IsServiceError(e) {
				ctx.JSON(http.StatusInternalServerError, newBadResp(-1, e.Error()))
				return
			}
		}
	}
}

func payload(ctx *gin.Context) (*Payload, bool) {
	var payload Payload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		bad(ctx, gin.H{"error": err.Error()})
		return nil, false
	}
	return &payload, true
}
