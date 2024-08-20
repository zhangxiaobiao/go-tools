package main

import (
	"bb/cmd/check"
	"bb/cmd/fileinfo"
	"bb/cmd/hosts"
	"bb/cmd/system"
	"bb/cmd/utils"
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var host = hosts.GetHosts()
var ck = check.GetCheck()
var sys = system.GetSystem()
var file = fileinfo.GetFileInfo()

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

/* 读取hosts */
func (a *App) ReadHostsFile() string {
	return host.ReadHostsFile()
}

/* 写入hosts */
func (a *App) EditHostsFile(content string) string {
	return host.EditHostsFile(content)
}

/* 写入打卡时间 */
func (a *App) SaveTime(start string, end string) int {
	return ck.SaveTime(start, end)
}

/* 获取所有打卡时间 */
func (a *App) AllChecks(date string) []check.Check {
	return ck.AllChecks(date)
}

/* 删除打卡时间 */
func (a *App) DelChecks(date string) int {
	return ck.DelChecks(date)
}

/* 备份数据 */
func (a *App) BackUpFile(path string) {
	sys.BackUpFile(path)
}

/* 备份数据 */
func (a *App) DesktopPath() string {
	return sys.DesktopPath()
}

/* 导入数据 */
func (a *App) ImportBackup(file string) utils.Resp[any] {
	return sys.ImportBackup(file)
}

/* 系统信息 */
func (a *App) DeviceInfo() utils.Resp[any] {
	return sys.DeviceInfo()
}

/* 系统信息 */
func (a *App) SelectFile() utils.Resp[any] {
	selection, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "选择文件",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "所有文件",
				Pattern:     "*.*",
			},
		},
	})
	if err != nil {
		// 处理错误情况
		return utils.Error[any](err.Error())
	}
	// 返回选中的文件路径
	return file.SelectFile(selection)
}

func (a *App) ChangeFileInfo(fileInfo fileinfo.FileInfo) utils.Resp[any] {
	return file.ChangeFileInfo(fileInfo)
}
