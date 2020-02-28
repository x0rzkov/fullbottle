package handler

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/errors"
	"github.com/vegchic/fullbottle/bottle/dao"
	pb "github.com/vegchic/fullbottle/bottle/proto/bottle"
	"github.com/vegchic/fullbottle/bottle/service"
	"github.com/vegchic/fullbottle/common"
	"github.com/vegchic/fullbottle/common/kv"
	"github.com/vegchic/fullbottle/config"
	"github.com/vegchic/fullbottle/weed"
	"time"
)

const UploadLockKey = "lock:token=%s"

type UploadHandler struct{}

func (*UploadHandler) GenerateUploadToken(ctx context.Context, req *pb.GenerateUploadTokenRequest, resp *pb.GenerateUploadTokenResponse) error {
	ownerId := req.GetOwnerId()
	filename := req.GetFilename()
	folderId := req.GetFolderId()
	hash := req.GetHash()
	size := req.GetSize()
	mime := req.GetMime()

	// create upload meta
	upload := weed.NewUploadMeta(ownerId, folderId, filename, hash, size, mime)
	var history weed.FileUploadMeta
	if err := kv.Get(upload.Token, &history); err == nil {
		upload = &history
	} else {
		// store for 15 days
		if err := kv.Set(upload.Token, upload, 15*24*time.Hour); err != nil {
			return err
		}
	}

	// check file is already uploaded, then client only need to call UploadFile without file data
	meta, err := dao.GetFileMetaByHash(hash)
	if err != nil {
		return err
	} else if meta != nil {
		resp.NeedUpload = false
	}

	resp.Uploaded = upload.UploadedChunk()
	resp.Token = upload.Token
	return nil
}

func (*UploadHandler) UploadFile(ctx context.Context, req *pb.UploadFileRequest, resp *pb.UploadFileResponse) error {
	token := req.GetToken()

	// lock for upload meta
	lock, err := kv.Obtain(fmt.Sprintf(UploadLockKey, token), 100*time.Millisecond)
	if err != nil {
		return err
	}
	defer lock.Release()

	// fetch upload meta
	upload := &weed.FileUploadMeta{}
	if err := kv.Get(token, upload); err != nil {
		return err
	}

	// check file is already uploaded
	meta, err := dao.GetFileMetaByHash(upload.Hash)
	if err != nil {
		return err
	} else if meta != nil {
		upload.SetStatus(weed.WeedDone)
	} else {
		// upload chunk
		offset := req.GetOffset()
		raw := req.GetRaw()
		if len(raw) == 0 {
			return errors.New(config.BottleSrvName, "Empty file", common.BadArgError)
		}
		err = upload.Upload(raw, offset)
		if err != nil {
			return err
		}
	}

	// upload weed done, create meta and file
	if upload.Status == weed.WeedDone {
		if meta == nil {
			meta = &dao.FileMeta{}
			meta.FromUploadMeta(upload)
			err := dao.CreateFileMeta(meta)
			if err != nil {
				return err
			}
		}

		info := &dao.FileInfo{}
		info.FromUploadMeta(upload)
		err = service.CreateFile(info, meta)
		if err != nil {
			return err
		}
		upload.SetStatus(weed.DBDone)
	}

	// refresh token
	if err := kv.RefreshValue(token, upload); err != nil {
		return err
	}

	resp.Status = pb.UploadStatus(upload.Status)
	resp.Uploaded = upload.UploadedChunk()
	return nil
}