package util

import (
	"github.com/vegchic/fullbottle/bottle/dao"
)

func GetSubEntryRecursive(folderId int64) (folders []*dao.FolderInfo, files []*dao.FileInfo, err error) {
	folders = make([]*dao.FolderInfo, 0)
	files = make([]*dao.FileInfo, 0)

	var subfolders []*dao.FolderInfo
	var subfiles []*dao.FileInfo
	parentIds := []int64{folderId}
	for len(parentIds) > 0 {
		subfolders, err = dao.GetFoldersByParentIds(parentIds)
		if err != nil {
			return
		}
		folders = append(folders, subfolders...)

		subfiles, err = dao.GetFilesByFolderIds(parentIds)
		if err != nil {
			return
		}
		files = append(files, subfiles...)

		parentIds = make([]int64, len(subfolders))
		for i, f := range subfolders {
			parentIds[i] = f.ID
		}
	}

	return
}

func GetSubEntry(folderId int64) (subfolder []*dao.FolderInfo, subfile []*dao.FileInfo, err error) {
	subfolder, err = dao.GetFoldersByParentId(folderId)
	if err != nil {
		return
	}

	subfile, err = dao.GetFilesByFolderId(folderId)
	if err != nil {
		return
	}

	return
}
