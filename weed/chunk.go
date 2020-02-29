package weed

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/vegchic/fullbottle/common"
	"github.com/vegchic/fullbottle/config"
	"github.com/vegchic/fullbottle/util"
)

type ChunkInfo struct {
	Fid    string `json:"fid"`
	Offset int64  `json:"offset"`
	Size   int64  `json:"size"`
}

type ChunkList []*ChunkInfo

type ChunkManifest struct {
	Name   string    `json:"name,omitempty"`
	Mime   string    `json:"mime,omitempty"`
	Size   int64     `json:"size,omitempty"`
	Chunks ChunkList `json:"chunks,omitempty"`
}

// generate manifest reader
func (m *ChunkManifest) Json() ([]byte, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, common.NewWeedError(err)
	}
	return b, nil
}

const (
	Inited = iota
	Uploading
	Manifest
	WeedDone
	DBDone
)

type FileUploadMeta struct {
	Token    string
	OwnerId  int64
	FolderId int64

	Status int // shouldn't modify directly

	Hash string
	Mime string
	Fid  string

	// chunk info
	ChunkManifest
	ChunkSize int64
}

func (f *FileUploadMeta) init() {
	raw := fmt.Sprintf("%d:%s:%d:%s", f.OwnerId, f.Name, f.FolderId, f.Hash)
	f.Token = util.TokenMd5(raw)
}

func (f *FileUploadMeta) SetStatus(s int) {
	// cannot rollback
	if s < f.Status {
		return
	}
	f.Status = s
}

func (f *FileUploadMeta) Marshal() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(*f)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (f *FileUploadMeta) Unmarshal(b []byte) error {
	buf := bytes.NewReader(b)
	dec := gob.NewDecoder(buf)

	err := dec.Decode(f)
	if err != nil {
		return err
	}

	return nil
}

// always upload in chunk, just make it simple
func (f *FileUploadMeta) Upload(raw []byte, offset int64) error {
	// if done, return
	if f.Status == WeedDone || f.Status == DBDone {
		return nil
	}
	// if in upload chunk step
	if f.Status == Inited || f.Status == Uploading {
		// reset the status
		f.Status = Uploading

		// upload
		reader := bytes.NewReader(raw)
		key, err := AssignFileKey()
		if err != nil {
			return err
		}

		_, err = UploadSingleFile(reader, f.Name, key.Fid, key.Url, false)
		if err != nil {
			return err
		}

		// add chunk info
		f.Chunks = append(f.Chunks, &ChunkInfo{
			Fid:    key.Fid,
			Offset: offset,
			Size:   reader.Size(),
		})

		// calculate if upload step is finish
		uploaded := int64(0)
		for _, c := range f.Chunks {
			uploaded += c.Size
		}
		if uploaded == f.Size {
			f.SetStatus(Manifest)
		}
	}

	// if in manifest step
	if f.Status == Manifest {
		// upload
		key, err := AssignFileKey()
		if err != nil {
			return err
		}
		b, err := f.ChunkManifest.Json()
		if err != nil {
			return err
		}
		_, err = UploadSingleFile(bytes.NewBuffer(b), f.Name, key.Fid, key.Url, true)
		if err != nil {
			return err
		}

		// update info
		f.SetStatus(WeedDone)
		f.Fid = key.Fid
	}
	return nil
}

func (f *FileUploadMeta) UploadedChunk() []int32 {
	var ranges []int32
	for _, c := range f.Chunks {
		ranges = append(ranges, int32(c.Offset/f.ChunkSize))
	}
	return ranges
}

func NewUploadMeta(ownerId int64, folderId int64, filename string, hash string, size int64, mime string) *FileUploadMeta {
	meta := &FileUploadMeta{
		OwnerId:       ownerId,
		FolderId:      folderId,
		Hash:          hash,
		Mime:          mime,
		Status:        Inited,
		ChunkManifest: ChunkManifest{Size: size, Name: filename},
		ChunkSize:     config.DefaultChunkSize,
	}
	meta.init()
	return meta
}
