package hosts

import (
	"bb/cmd/utils"
	"os"
	"runtime"

	"github.com/siddontang/go/log"
)

type Hosts struct {
}

var hosts *Hosts

var deskPath string

// var hostsPath = "D:\\code\\TEST\\go\\etc\\hosts"

func GetHosts() *Hosts {
	return hosts
}

/* 读hosts */
func (h *Hosts) ReadHostsFile() utils.Resp[any] {
	switch runtime.GOOS {
	case "windows":
		deskPath = "C:\\Windows\\System32\\drivers\\etc\\hosts"
	case "darwin":
		deskPath = "/private/etc/hosts"
	case "linux":
		deskPath = "/etc/hosts"
	default:
		deskPath = ""
		log.Error("未知操作系统:", runtime.GOOS)
	}
	info, err := os.Stat(deskPath)
	if os.IsNotExist(err) || info.IsDir() {
		return utils.Error[any]("Hosts file not found")
	}
	content, err := os.ReadFile(deskPath)
	if err != nil {
		return utils.Error[any]("Unable to read hosts file")
	}
	return utils.Success[any](string(content), "读取成功")
}

/* 写hosts */
func (h *Hosts) EditHostsFile(content string) utils.Resp[any] {
	info, err := os.OpenFile(deskPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return utils.Error[any](err.Error())
	}
	defer info.Close()
	_, err = info.WriteString(content)
	if err != nil {
		return utils.Error[any](err.Error())
	}
	return utils.Success[any](nil, "写入成功")
}
