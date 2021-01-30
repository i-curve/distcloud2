package cloud

import (
	"clouddist/pkg/file"
	"clouddist/pkg/setting"
	"io/ioutil"
	"log"
	"os"
)

var (
	// F 文件
	F        *os.File
	dataPath = "data/"
	path     = ""
)

// Item 文件样式结构
type Item struct {
	Name  string
	IsDir bool
	Size  int
}

// Setup 初始化
func Setup() {
	if setting.AppSetting.DataPath != "" {
		dataPath = setting.AppSetting.DataPath
	}
	dir, _ := os.Getwd()
	path = dir + "/" + dataPath
	err := file.IsNotExistMkDir(path)
	if err != nil {
		log.Fatalf("data.Setup: %v", err)
	}
}

// CreateUSER 创建新的用户文件
func CreateUSER(username string) bool {
	userpath := path + "/" + username
	if err := file.MkDir(userpath); err != nil {
		return false
	}
	if err := file.MkDir(userpath + "/Documents"); err != nil {
		return false
	}
	if err := file.MkDir(userpath + "/Videos"); err != nil {
		return false
	}
	if err := file.MkDir(userpath + "/Music"); err != nil {
		return false
	}
	if err := file.MkDir(userpath + "/Pictures"); err != nil {
		return false
	}
	return true
}

// ShowDir 显示所有目录和文件
func ShowDir(username, ipath string) []Item {
	dir := path + username + ipath
	files, _ := ioutil.ReadDir(dir)

	var data []Item
	for _, f := range files {
		data = append(data, Item{
			Name:  f.Name(),
			IsDir: f.IsDir(),
			Size:  int(f.Size()),
		})
	}
	return data
}
