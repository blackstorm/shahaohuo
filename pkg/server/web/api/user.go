package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"shahaohuo.com/shahaohuo/pkg/server/dto"
	"shahaohuo.com/shahaohuo/pkg/server/orm"
)

func PatchUser(c *gin.Context) {
	userId, _ := mustGetUserIdByContent(c)
	u, e := orm.FindUserById(userId)
	if e != nil {
		internalServerError(c, nil)
		return
	}
	if u == nil {
		notfound(c)
		return
	}

	var pathReq dto.PatchUserRequest
	if e := c.BindJSON(&pathReq); e != nil {
		bad(c, gin.H{"message": "bad json"})
		return
	}

	needPatch := false
	if len(pathReq.Avatar) > 0 && u.Avatar != pathReq.Avatar {
		needPatch = true
		u.Avatar = pathReq.Avatar
	}
	if len(pathReq.Name) > 0 && u.Name != pathReq.Name {
		needPatch = true
		u.Name = pathReq.Name
	}
	if len(pathReq.Bio) > 0 && u.Bio != pathReq.Bio {
		needPatch = true
		u.Bio = pathReq.Bio
	}
	if needPatch {
		if e := u.Update(); e != nil {
			logrus.Error(e)
			c.JSON(500, nil)
		}
	}

	var ret dto.UserResponse
	_ = copier.Copy(&ret, &u)

	c.JSON(200, &ret)
}
