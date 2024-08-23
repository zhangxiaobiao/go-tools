package fileinfo

import (
	"bb/cmd/utils"
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/golang-module/carbon"
)

type FileInfo struct {
	FilePath   string `json:"filePath"`
	CreateTime int64  `json:"createTime"`
	AccessTime int64  `json:"accessTime"`
	UpdateTime int64  `json:"updateTime"`
}

var fileInfo *FileInfo

func GetFileInfo() *FileInfo {
	return fileInfo
}

func (f *FileInfo) SelectFile(filePath string) utils.Resp[any] {
	// 获取文件状态信息
	finfo, err := os.Stat(filePath)
	if err != nil {
		return utils.Error[any](err.Error())
	}
	fio := FileInfo{}
	fio.FilePath = filePath
	// 获取底层系统信息
	sys := finfo.Sys()
	switch sys := sys.(type) {
	case *syscall.Stat_t:
		// Unix-like 系统
		fio.CreateTime = time.Unix(sys.Ctimespec.Sec, sys.Ctimespec.Nsec).Unix()
		fio.AccessTime = time.Unix(sys.Atimespec.Sec, sys.Atimespec.Nsec).Unix()
		fio.UpdateTime = time.Unix(sys.Mtimespec.Sec, sys.Ctimespec.Nsec).Unix()
		break
		// case *syscall.Win32FileAttributeData:
		// 	// Windows 系统
		// 	fio.CreateTime = time.Unix(0, sys.CreationTime.Nanoseconds()).Unix()
		// 	fio.AccessTime = time.Unix(0, sys.LastAccessTime.Nanoseconds()).Unix()
		// 	fio.UpdateTime = time.Unix(0, sys.LastWriteTime.Nanoseconds()).Unix()
		// 	break
	}
	return utils.Success[any](fio, "获取成功")
}

/* 修改文件信息 */
func (f *FileInfo) ChangeFileInfo(fileInfo FileInfo) utils.Resp[any] {
	fmt.Printf("%+v", fileInfo)
	accessTime := carbon.CreateFromTimestamp(fileInfo.AccessTime).ToStdTime()
	modifyTime := carbon.CreateFromTimestamp(fileInfo.UpdateTime).ToStdTime()
	err := os.Chtimes(fileInfo.FilePath, accessTime, modifyTime)
	if err != nil {
		return utils.Error[any](err.Error())
	}
	return utils.Success[any](nil, "修改成功")
}
