package api

import (
	"bufio"
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	goImg "image"
	"io/ioutil"
	"net/http"
	"shahaohuo.com/shahaohuo/pkg/server/image"
	"shahaohuo.com/shahaohuo/pkg/server/orm"
	"shahaohuo.com/shahaohuo/pkg/server/storage"
	"shahaohuo.com/shahaohuo/pkg/util"
)

const (
	uploadImageTypeHaohuo = "haohuo"
	uploadImageTypeAvatar = "avatar"
)

func UploadImage(c *gin.Context) {
	userId, _ := mustGetUserIdByContent(c)

	// gen image path
	var imagePath string
	uploadType := c.Query("t")
	switch uploadType {
	case uploadImageTypeHaohuo:
		imagePath = image.HaohuoImagePathRandom()
		break
	case uploadImageTypeAvatar:
		imagePath = image.UserAvatarPathGen(userId)
		break
	default:
		bad(c, "must set upload type")
		return
	}

	// image process
	fileHeader, err := c.FormFile("image")
	if err != nil {
		logrus.Error(err)
		bad(c, "get file error")
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		logrus.Error(err)
		internalServerError(c, "server error")
		return
	}
	// TODO zero copy
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		logrus.Error(err)
		internalServerError(c, "server error")
		return
	}

	// check file content type
	contentType := http.DetectContentType(fileBytes)
	if !util.CheckIsContain(contentType, util.PNG, util.JPEG) {
		bad(c, newBadResp(1, "file must is an image"))
		return
	}

	// resize avatar
	if isUploadAvatar(uploadType) {
		if ret, e := avatarResize(fileBytes, 128, 128); e != nil {
			internalServerError(c, "server error")
		} else {
			fileBytes = ret
		}
	}

	if err := storage.GetBucket().Upload(imagePath, contentType, fileBytes); err != nil {
		logrus.Error(err)
		internalServerError(c, "server error")
		return
	}

	e := orm.SaveImage(userId, xid.New().String(), imagePath)
	if e != nil {
		logrus.Error(e)
	}

	ok(c, gin.H{
		"path": imagePath,
	})
}

func isUploadAvatar(t string) bool {
	return t == uploadImageTypeAvatar
}

func avatarResize(data []byte, width, height int) ([]byte, error) {
	img, _, e := goImg.Decode(bytes.NewReader(data))
	if e != nil {
		logrus.Error(e)
		return nil, e
	}
	src := imaging.Thumbnail(img, width, height, imaging.Lanczos)
	var b bytes.Buffer
	if e := imaging.Encode(bufio.NewWriter(&b), src, imaging.PNG); e != nil {
		return nil, e
	}
	return b.Bytes(), nil
}
