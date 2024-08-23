package editor

import (
	"bb/cmd/utils"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-module/carbon"
)

type Editor struct {
}

var editor *Editor

func GetEditor() *Editor {
	return editor
}

func (e *Editor) SaveFile(str string, filePath string) utils.Resp[any] {
	filename := "Docs/" + carbon.Now().Format("YmdHis") + ".md"
	if filePath != "" {
		filename = filePath
	}
	fmt.Println(filename)
	dir := filepath.Dir(filename)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return utils.Error[any](err.Error())
	}
	data := []byte(str)
	err = os.WriteFile(filename, data, os.ModePerm)
	if err != nil {
		return utils.Error[any](err.Error())
	}
	return utils.Success[any](nil, "保存成功")
}

func (e *Editor) OpenFile(filePath string) utils.Resp[any] {
	_, err := os.Stat(filePath)
	if err != nil {
		return utils.Error[any](err.Error())
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return utils.Error[any](err.Error())
	}
	result := map[string]string{
		"content": string(data),
		"path":    filePath,
		"name":    filepath.Base(filePath),
	}
	return utils.Success[any](result, "保存成功")
}
