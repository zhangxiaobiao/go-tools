package hosts

import (
	"bb/cmd/utils"
	"os"
)

type Hosts struct {
}

var hosts *Hosts

var hostsPath = "C:\\Windows\\System32\\drivers\\etc\\hosts"

// var hostsPath = "D:\\code\\TEST\\go\\etc\\hosts"

func GetHosts() *Hosts {
	return hosts
}

/* 读hosts */
func (h *Hosts) ReadHostsFile() utils.Resp[any] {
	info, err := os.Stat(hostsPath)
	if os.IsNotExist(err) || info.IsDir() {
		return utils.Error[any]("Hosts file not found")
	}
	content, err := os.ReadFile(hostsPath)
	if err != nil {
		return utils.Error[any]("Unable to read hosts file")
	}
	return utils.Success[any](string(content), "读取成功")
}

/* 写hosts */
func (h *Hosts) EditHostsFile(content string) utils.Resp[any] {
	info, err := os.OpenFile(hostsPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
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
