package hosts

import (
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
func (h *Hosts) ReadHostsFile() string {
	info, err := os.Stat(hostsPath)
	if os.IsNotExist(err) || info.IsDir() {
		return "Hosts file not found"
	}
	content, err := os.ReadFile(hostsPath)
	if err != nil {
		return "Unable to read hosts file"
	}
	return string(content)
}

/* 写hosts */
func (h *Hosts) EditHostsFile(content string) string {
	info, err := os.OpenFile(hostsPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err.Error()
	}
	defer info.Close()
	_, err = info.WriteString(content)
	if err != nil {
		return err.Error()
	}
	return "写入成功"
}
