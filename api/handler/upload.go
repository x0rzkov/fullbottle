package handler

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/errors"
	"github.com/vegchic/fullbottle/api/util"
	pbbottle "github.com/vegchic/fullbottle/bottle/proto/bottle"
	"github.com/vegchic/fullbottle/common"
	"io"
	"mime/multipart"
	"net/http"
)

func GetUploadToken(c *gin.Context) {
	u, _ := c.Get("cur_user_id")
	uid := u.(int64)

	body := struct {
		FolderId int64  `json:"folder_id" bindings:"required"`
		Filename string `json:"filename" bindings:"required,min=1,max=100"`
		Mime     string `json:"mime" bindings:"required"`
		Hash     string `json:"hash" bindings:"required"`
		Size     int64  `json:"size" bindings:"required"`
	}{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	bottleClient := common.BottleSrvClient()

	bottleMeta, err := bottleClient.GetBottleMetadata(util.RpcContext(c), &pbbottle.GetBottleMetadataRequest{Uid: uid})
	if err != nil {
		e := errors.Parse(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": e.Detail,
		})
		return
	}
	if bottleMeta.Remain < body.Size {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "Remain space not enough",
		})
		return
	}

	req := &pbbottle.GenerateUploadTokenRequest{
		OwnerId:  uid,
		Filename: body.Filename,
		FolderId: body.FolderId,
		Hash:     body.Hash,
		Size:     body.Size,
		Mime:     body.Mime,
	}
	resp, err := bottleClient.GenerateUploadToken(util.RpcContext(c), req)
	if err != nil {
		e := errors.Parse(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": e.Detail,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Success",
		"result": gin.H{
			"token":       resp.Token,
			"need_upload": resp.NeedUpload,
			"uploaded":    resp.Uploaded,
		},
	})
}

func UploadFile(c *gin.Context) {
	body := struct {
		FilePart *multipart.FileHeader `form:"file_part"`
		Token    string                `form:"token" bindings:"required"`
		Offset   int64                 `form:"offset" bindings:"required"`
	}{}
	if err := c.ShouldBind(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid token",
		})
		return
	}
	var b []byte
	if body.FilePart != nil {
		f, err := body.FilePart.Open()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": "Invalid file",
			})
			return
		}
		defer f.Close()

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, f); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": "Invalid file",
			})
			return
		}
		b = buf.Bytes()
	}

	bottleClient := common.BottleSrvClient()
	req := &pbbottle.UploadFileRequest{
		Token:  body.Token,
		Offset: body.Offset,
		Raw:    b,
	}
	resp, err := bottleClient.UploadFile(util.RpcContext(c), req)
	if err != nil {
		e := errors.Parse(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": e.Detail,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Success",
		"result": gin.H{
			"status":   resp.Status,
			"uploaded": resp.Uploaded,
		},
	})
}
