package system

import (
	"bb/cmd/db/entity"
	"bb/cmd/global"
	"encoding/base64"
	"encoding/json"
	"os/user"
	"path/filepath"
	"runtime"
	"time"

	"bb/cmd/utils"

	"github.com/labstack/gommon/log"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
)

type System struct {
}

var system *System

func GetSystem() *System {
	return system
}

func (s *System) BackUpFile(path string) {
	var checks []entity.Check
	result := global.Db.Model(&entity.Check{}).Find(&checks)
	if result.Error != nil {
		log.Fatal("查询checks失败", result.Error)
		return
	}
	data := make(map[string]interface{})
	data["checks"] = checks
	dt, _ := json.Marshal(data)
	dtStr := string(dt)
	path = path + "/bb.json"
	utils.WriteFile(dtStr, path)
}

func (s *System) DesktopPath() string {
	u, err := user.Current()
	if err != nil {
		return ""
	}
	var deskPath string
	switch runtime.GOOS {
	case "windows":
		deskPath = filepath.Join(u.HomeDir, "Desktop")
	case "darwin":
		deskPath = filepath.Join(u.HomeDir, "Desktop")
	case "linux":
		deskPath = filepath.Join(u.HomeDir, "Desktop")
	default:
		deskPath = ""
		log.Error("未知操作系统:", runtime.GOOS)
	}
	return deskPath
}

func (s *System) ImportBackup(base64File string) utils.Resp[any] {
	// 解码Base64字符串
	fileBytes, err := base64.StdEncoding.DecodeString(base64File)
	var jsonData map[string]interface{}
	err = json.Unmarshal(fileBytes, &jsonData)
	if err != nil {
		return utils.Error[any]("文件解析失败")
	}
	var checks []entity.Check
	checksMarshal, _ := json.Marshal(jsonData["checks"])
	_ = json.Unmarshal(checksMarshal, &checks)
	for _, item := range checks {
		global.Db.Save(&item)
	}
	return utils.Success[any](nil, "导入成功")
}

func (s *System) DeviceInfo() utils.Resp[any] {
	info, _ := cpu.Info()
	host, _ := host.Info()
	percent, _ := cpu.Percent(time.Second, false)
	load, _ := load.Avg()
	memory, _ := mem.VirtualMemory()
	parts, _ := disk.Partitions(true)
	var partsArr = make([]interface{}, 0, 10)
	for _, part := range parts {
		diskInfo, _ := disk.Usage(part.Mountpoint)
		partsArr = append(partsArr, diskInfo)
	}
	var result = make(map[string]interface{})
	result["info"] = info
	result["percent"] = percent
	result["load"] = load
	result["host"] = host
	result["disk"] = partsArr
	result["memory"] = memory
	return utils.Success[any](result, "")
}
